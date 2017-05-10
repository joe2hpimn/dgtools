package codegen

import (
	"fmt"
	"log"
	"vitessedata/phi/datatype"
)

var srccode string = `
//
// DO NOT EDIT
// Code is GENERATED BY VitesseData Phi Code Gen.  DO NOT EDIT
// 

package main

import (
	"vitessedata/phi/datatype"
	"vitessedata/phi/phirun"
	"vitessedata/phi/proto/xdrive"
)

func Log(msg string, args ...interface{}) {
	phirun.Log(msg, args...)
}

%s

type phi_runtime_t struct {
	inRecs []InRecord
	currInRec int32

	outRecs []*OutRecord
	currOutRec int32
}

%s

var phirt phi_runtime_t 

func init() {
	phirt.outRecs = make([]*OutRecord, 1024)
}

func NextInput() *InRecord {
	if phirt.inRecs == nil {
		inMsg, err := phirun.ReadXMsg()
		if err != nil || inMsg == nil || inMsg.Xflag == xdrive.XMsg_EOS { 
			// End of input stream.
			return nil
		} else if inMsg.Rowset == nil {
			phirt.inRecs = nil
			return NextInput()
		}

		// if inMsg.Rowset != nil, it must has col data.
		phirt.loadInRecs(inMsg.Rowset)
	}
	
	if phirt.currInRec < int32(len(phirt.inRecs)) {
		ret := &phirt.inRecs[phirt.currInRec]
		phirt.currInRec += 1
		return ret
	} else {
		// All InRec from inMsg exhausted.  We will flush outRec.
		Log("Flush Output Because We Need More Input.\n")
		FlushOutput(xdrive.XMsg_EOB)
		phirt.inRecs = nil
		return NextInput()
	}
}



func WriteOutput(r *OutRecord) {
	if (r == nil) {
		Log("Flush Output Because Done.\n")
		FlushOutput(xdrive.XMsg_EOB)
		FlushOutput(xdrive.XMsg_EOS)
	} else {
		if phirt.currOutRec < 1024 {
			phirt.outRecs[phirt.currOutRec] = r
			phirt.currOutRec += 1
			return
		} else {
			Log("Flush Output Because Output Buffer Full.\n")
			FlushOutput(xdrive.XMsg_CONTINUE)
			WriteOutput(r)
		}
	}
}

func FlushOutput(flag xdrive.XMsg_XMsgFlag) { 
	var msg xdrive.XMsg
	msg.Xflag = flag
	msg.Code = 0
	Log("Flush output, flag is %%v, currOutRec is %%v.\n", flag, phirt.currOutRec)

	if flag == xdrive.XMsg_EOB || flag == xdrive.XMsg_CONTINUE { 
		msg.Rowset = phirt.writeOutRecs()
		phirt.currOutRec = 0
	}
	phirun.WriteXMsg(&msg)
}
`

func genGoRecDecl(isch *ioSchema, osch *ioSchema) string {
	s1 := isch.genGoRec("InRecord")
	s2 := osch.genGoRec("OutRecord")
	return s1 + "\n" + s2
}

func genGoRtMethods(isch *ioSchema, osch *ioSchema) string {
	m1 := `
func (rt *phi_runtime_t) loadInRecs(rs *xdrive.XRowSet) {
	cols := rs.Columns
	nrow := cols[0].Nrow
	rt.currInRec = 0
	rt.inRecs = make([]InRecord, nrow, nrow)
`
	for i, colname := range isch.ColNames {
		gotyp := isch.ColTypes[i].Names()[datatype.Name]
		xcolname := isch.ColTypes[i].Names()[datatype.XDrColName]
		m1 += fmt.Sprintf("\tfor r := int32(0); r < nrow; r++ {\n")
		m1 += fmt.Sprintf("\t\tif cols[%d].Nullmap[r] {\n", i)
		m1 += fmt.Sprintf("\t\t\trt.inRecs[r].Set_%s_Null()\n", colname)
		m1 += fmt.Sprintf("\t\t} else {\n")
		if gotyp == "bool" {
			m1 += fmt.Sprintf("\t\t\trt.inRecs[r].Set_%s(cols[%d].%s[r] != 0)\n", colname, i, xcolname)
		} else {
			m1 += fmt.Sprintf("\t\t\trt.inRecs[r].Set_%s(cols[%d].%s[r])\n", colname, i, xcolname)
		}
		m1 += fmt.Sprintf("\t\t}\n")
		m1 += fmt.Sprintf("\t}\n")
	}

	m1 += "}\n\n"

	oncol := len(osch.ColNames)

	m2 := `
func (rt *phi_runtime_t) writeOutRecs() *xdrive.XRowSet {
	if rt.currOutRec == 0 {
		return nil
	}
	var rs xdrive.XRowSet
	nrow := rt.currOutRec
`
	m2 += fmt.Sprintf("\trs.Columns = make([]*xdrive.XCol, %d, %d)\n", oncol, oncol)
	for i, colname := range osch.ColNames {
		gotyp := osch.ColTypes[i].Names()[datatype.Name]
		xcolname := osch.ColTypes[i].Names()[datatype.XDrColName]
		m2 += fmt.Sprintf("\trs.Columns[%d] = new(xdrive.XCol)\n", i)
		m2 += fmt.Sprintf("\trs.Columns[%d].Colname = \"%s\"\n", i, colname)
		m2 += fmt.Sprintf("\trs.Columns[%d].Nrow = nrow\n", i)
		m2 += fmt.Sprintf("\trs.Columns[%d].Nullmap = make([]bool, nrow, nrow)\n", i)
		if gotyp == "bool" {
			m2 += fmt.Sprintf("\trs.Columns[%d].%s = make([]int32, nrow, nrow)\n", i, xcolname)
		} else {
			m2 += fmt.Sprintf("\trs.Columns[%d].%s = make([]%s, nrow, nrow)\n", i, xcolname, gotyp)
		}
		m2 += fmt.Sprintf("\tfor r := int32(0); r < nrow; r++ {\n")
		m2 += fmt.Sprintf("\t\tv, ok := rt.outRecs[r].Get_%s()\n", colname)
		m2 += fmt.Sprintf("\t\tif !ok {\n")
		m2 += fmt.Sprintf("\t\t\trs.Columns[%d].Nullmap[r] = true\n", i)
		m2 += fmt.Sprintf("\t\t} else {\n")
		m2 += fmt.Sprintf("\t\t\trs.Columns[%d].Nullmap[r] = false\n", i)
		m2 += fmt.Sprintf("\t\t\trs.Columns[%d].%s[r] = v\n", i, xcolname)
		m2 += fmt.Sprintf("\t\t}\n")
		m2 += fmt.Sprintf("\t}\n\n")
	}
	m2 += fmt.Sprintf("\treturn &rs\n")
	m2 += fmt.Sprintf("}\n")

	return m1 + m2
}

func genGo(fn string) {
	isch, osch, err := processSchema(fn)
	if err != nil {
		log.Fatal(err)
	}

	recStr := genGoRecDecl(isch, osch)
	rtStr := genGoRtMethods(isch, osch)

	result := fmt.Sprintf(srccode, recStr, rtStr)
	fmt.Println(result)
}
