//
// BEGIN INPUT TYPES
// a int32
// b float32
// c string
// END INPUT TYPES
//
// BEGIN OUTPUT TYPES
// x int32
// y float32
// z string
// END OUTPUT TYPES
//

package main

import (
	"os"
)

func do_i() {
	var outrec OutRecord

	for xx := 0; xx < 1000; xx++ {
		Log("I write 1 rec\n")
		outrec.Set_x(1)
		outrec.Set_y(2.0)
		outrec.Set_z("foo")
		WriteOutput(&outrec)

		Log("I write 2 rec\n")
		outrec.Set_x(1)
		outrec.Set_y(2.0)
		outrec.Set_z("foo")
		WriteOutput(&outrec)

		Log("I write 3 rec\n")
		outrec.Set_x(1)
		outrec.Set_y(2.0)
		outrec.Set_z("foo")
		WriteOutput(&outrec)
	}

	Log("I am done.\n")
	WriteOutput(nil)
}

func do_x() {
	for rec := NextInput(); rec != nil; rec = NextInput() {
		Log("X get one rec\n")
		a, aok := rec.Get_a()
		b, bok := rec.Get_b()
		c, cok := rec.Get_c()

		var outrec OutRecord
		if aok {
			outrec.Set_x(a * 2)
		}
		if bok {
			outrec.Set_y(b * 2.0)
		}
		if cok {
			outrec.Set_z(c + c)
		}

		WriteOutput(&outrec)
	}
	Log("X is done\n")
	WriteOutput(nil)
}

func do_o() {
	for rec := NextInput(); rec != nil; rec = NextInput() {
		Log("O get one rec.\n")
		a, aok := rec.Get_a()
		b, bok := rec.Get_b()
		c, cok := rec.Get_c()
		Log("Rec: %d %v, %f %v, %s %v\n", a, aok, b, bok, c, cok)
	}
	// WriteOutput(nil)
	Log("O is done\n")
}

func main() {
	switch os.Args[1] {
	case "i":
		do_i()
	case "x":
		do_x()
	case "o":
		do_o()
	}
}