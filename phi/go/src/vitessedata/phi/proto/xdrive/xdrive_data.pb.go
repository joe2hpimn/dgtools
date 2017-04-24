// Code generated by protoc-gen-go.
// source: xdrive_data.proto
// DO NOT EDIT!

/*
Package xdrive is a generated protocol buffer package.

It is generated from these files:
	xdrive_data.proto

It has these top-level messages:
	ColumnDesc
	Filter
	CSVSpec
	FileSpec
	StringList
	KeyValue
	KeyValueList
	RmgrInfo
	ReadRequest
	SampleRequest
	DataReply
	XCol
	XRowSet
	PluginDataReply
	SizeMetaRequest
	SizeMetaReply
	PluginSizeMetaReply
	PageData
	WriteRequest
	PluginWriteRequest
	WriteReply
	PluginWriteReply
	XMsg
*/
package xdrive

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SpqType int32

const (
	SpqType_UNKNOWN SpqType = 0
	SpqType_BOOL    SpqType = 1
	SpqType_INT16   SpqType = 2
	SpqType_INT32   SpqType = 3
	SpqType_INT64   SpqType = 4
	SpqType_INT128  SpqType = 5
	SpqType_FLOAT   SpqType = 6
	SpqType_DOUBLE  SpqType = 7
	// BYTEA    = 0x0008;   not supported.
	SpqType_CSTR             SpqType = 9
	SpqType_DEC64            SpqType = 10
	SpqType_DEC128           SpqType = 11
	SpqType_DATE             SpqType = 65539
	SpqType_TIME_MILLIS      SpqType = 131075
	SpqType_TIMESTAMP_MILLIS SpqType = 196612
	SpqType_TIME_MICROS      SpqType = 262148
	SpqType_TIMESTAMP_MICROS SpqType = 327684
	SpqType_JSON             SpqType = 393225
)

var SpqType_name = map[int32]string{
	0:      "UNKNOWN",
	1:      "BOOL",
	2:      "INT16",
	3:      "INT32",
	4:      "INT64",
	5:      "INT128",
	6:      "FLOAT",
	7:      "DOUBLE",
	9:      "CSTR",
	10:     "DEC64",
	11:     "DEC128",
	65539:  "DATE",
	131075: "TIME_MILLIS",
	196612: "TIMESTAMP_MILLIS",
	262148: "TIME_MICROS",
	327684: "TIMESTAMP_MICROS",
	393225: "JSON",
}
var SpqType_value = map[string]int32{
	"UNKNOWN":          0,
	"BOOL":             1,
	"INT16":            2,
	"INT32":            3,
	"INT64":            4,
	"INT128":           5,
	"FLOAT":            6,
	"DOUBLE":           7,
	"CSTR":             9,
	"DEC64":            10,
	"DEC128":           11,
	"DATE":             65539,
	"TIME_MILLIS":      131075,
	"TIMESTAMP_MILLIS": 196612,
	"TIME_MICROS":      262148,
	"TIMESTAMP_MICROS": 327684,
	"JSON":             393225,
}

func (x SpqType) String() string {
	return proto.EnumName(SpqType_name, int32(x))
}
func (SpqType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ColumnDesc struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type int32  `protobuf:"varint,2,opt,name=type" json:"type,omitempty"`
}

func (m *ColumnDesc) Reset()                    { *m = ColumnDesc{} }
func (m *ColumnDesc) String() string            { return proto.CompactTextString(m) }
func (*ColumnDesc) ProtoMessage()               {}
func (*ColumnDesc) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ColumnDesc) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ColumnDesc) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type Filter struct {
	Op     string   `protobuf:"bytes,1,opt,name=op" json:"op,omitempty"`
	Column string   `protobuf:"bytes,2,opt,name=column" json:"column,omitempty"`
	Args   []string `protobuf:"bytes,3,rep,name=args" json:"args,omitempty"`
}

func (m *Filter) Reset()                    { *m = Filter{} }
func (m *Filter) String() string            { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()               {}
func (*Filter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Filter) GetOp() string {
	if m != nil {
		return m.Op
	}
	return ""
}

func (m *Filter) GetColumn() string {
	if m != nil {
		return m.Column
	}
	return ""
}

func (m *Filter) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

type CSVSpec struct {
	Delimiter string `protobuf:"bytes,1,opt,name=delimiter" json:"delimiter,omitempty"`
	Nullstr   string `protobuf:"bytes,2,opt,name=nullstr" json:"nullstr,omitempty"`
	Header    bool   `protobuf:"varint,3,opt,name=header" json:"header,omitempty"`
	Quote     string `protobuf:"bytes,4,opt,name=quote" json:"quote,omitempty"`
	Escape    string `protobuf:"bytes,5,opt,name=escape" json:"escape,omitempty"`
}

func (m *CSVSpec) Reset()                    { *m = CSVSpec{} }
func (m *CSVSpec) String() string            { return proto.CompactTextString(m) }
func (*CSVSpec) ProtoMessage()               {}
func (*CSVSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CSVSpec) GetDelimiter() string {
	if m != nil {
		return m.Delimiter
	}
	return ""
}

func (m *CSVSpec) GetNullstr() string {
	if m != nil {
		return m.Nullstr
	}
	return ""
}

func (m *CSVSpec) GetHeader() bool {
	if m != nil {
		return m.Header
	}
	return false
}

func (m *CSVSpec) GetQuote() string {
	if m != nil {
		return m.Quote
	}
	return ""
}

func (m *CSVSpec) GetEscape() string {
	if m != nil {
		return m.Escape
	}
	return ""
}

type FileSpec struct {
	Path    string   `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Format  string   `protobuf:"bytes,2,opt,name=format" json:"format,omitempty"`
	Csvspec *CSVSpec `protobuf:"bytes,3,opt,name=csvspec" json:"csvspec,omitempty"`
}

func (m *FileSpec) Reset()                    { *m = FileSpec{} }
func (m *FileSpec) String() string            { return proto.CompactTextString(m) }
func (*FileSpec) ProtoMessage()               {}
func (*FileSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *FileSpec) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *FileSpec) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *FileSpec) GetCsvspec() *CSVSpec {
	if m != nil {
		return m.Csvspec
	}
	return nil
}

type StringList struct {
	Str []string `protobuf:"bytes,1,rep,name=str" json:"str,omitempty"`
}

func (m *StringList) Reset()                    { *m = StringList{} }
func (m *StringList) String() string            { return proto.CompactTextString(m) }
func (*StringList) ProtoMessage()               {}
func (*StringList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StringList) GetStr() []string {
	if m != nil {
		return m.Str
	}
	return nil
}

type KeyValue struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *KeyValue) Reset()                    { *m = KeyValue{} }
func (m *KeyValue) String() string            { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()               {}
func (*KeyValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *KeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type KeyValueList struct {
	Kv []*KeyValue `protobuf:"bytes,1,rep,name=kv" json:"kv,omitempty"`
}

func (m *KeyValueList) Reset()                    { *m = KeyValueList{} }
func (m *KeyValueList) String() string            { return proto.CompactTextString(m) }
func (*KeyValueList) ProtoMessage()               {}
func (*KeyValueList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *KeyValueList) GetKv() []*KeyValue {
	if m != nil {
		return m.Kv
	}
	return nil
}

type RmgrInfo struct {
	Scheme   string        `protobuf:"bytes,1,opt,name=scheme" json:"scheme,omitempty"`
	Format   string        `protobuf:"bytes,2,opt,name=format" json:"format,omitempty"`
	Rpath    string        `protobuf:"bytes,3,opt,name=rpath" json:"rpath,omitempty"`
	Conf     *KeyValueList `protobuf:"bytes,4,opt,name=conf" json:"conf,omitempty"`
	Pluginop string        `protobuf:"bytes,5,opt,name=pluginop" json:"pluginop,omitempty"`
}

func (m *RmgrInfo) Reset()                    { *m = RmgrInfo{} }
func (m *RmgrInfo) String() string            { return proto.CompactTextString(m) }
func (*RmgrInfo) ProtoMessage()               {}
func (*RmgrInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *RmgrInfo) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *RmgrInfo) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *RmgrInfo) GetRpath() string {
	if m != nil {
		return m.Rpath
	}
	return ""
}

func (m *RmgrInfo) GetConf() *KeyValueList {
	if m != nil {
		return m.Conf
	}
	return nil
}

func (m *RmgrInfo) GetPluginop() string {
	if m != nil {
		return m.Pluginop
	}
	return ""
}

type ReadRequest struct {
	// Which file(s)
	Filespec *FileSpec `protobuf:"bytes,1,opt,name=filespec" json:"filespec,omitempty"`
	// Table Schema
	Columndesc []*ColumnDesc `protobuf:"bytes,2,rep,name=columndesc" json:"columndesc,omitempty"`
	// Names of required columns
	Columnlist []string `protobuf:"bytes,3,rep,name=columnlist" json:"columnlist,omitempty"`
	// Filters
	Filter []*Filter `protobuf:"bytes,4,rep,name=filter" json:"filter,omitempty"`
	// Fragment
	FragId  int32 `protobuf:"varint,5,opt,name=frag_id,json=fragId" json:"frag_id,omitempty"`
	FragCnt int32 `protobuf:"varint,6,opt,name=frag_cnt,json=fragCnt" json:"frag_cnt,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ReadRequest) GetFilespec() *FileSpec {
	if m != nil {
		return m.Filespec
	}
	return nil
}

func (m *ReadRequest) GetColumndesc() []*ColumnDesc {
	if m != nil {
		return m.Columndesc
	}
	return nil
}

func (m *ReadRequest) GetColumnlist() []string {
	if m != nil {
		return m.Columnlist
	}
	return nil
}

func (m *ReadRequest) GetFilter() []*Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *ReadRequest) GetFragId() int32 {
	if m != nil {
		return m.FragId
	}
	return 0
}

func (m *ReadRequest) GetFragCnt() int32 {
	if m != nil {
		return m.FragCnt
	}
	return 0
}

type SampleRequest struct {
	// Which file(s)
	Filespec *FileSpec `protobuf:"bytes,1,opt,name=filespec" json:"filespec,omitempty"`
	// Table Schema
	Columndesc []*ColumnDesc `protobuf:"bytes,2,rep,name=columndesc" json:"columndesc,omitempty"`
	// Fragment
	FragId  int32 `protobuf:"varint,3,opt,name=frag_id,json=fragId" json:"frag_id,omitempty"`
	FragCnt int32 `protobuf:"varint,4,opt,name=frag_cnt,json=fragCnt" json:"frag_cnt,omitempty"`
	// Sample size
	Nrow int32 `protobuf:"varint,5,opt,name=nrow" json:"nrow,omitempty"`
}

func (m *SampleRequest) Reset()                    { *m = SampleRequest{} }
func (m *SampleRequest) String() string            { return proto.CompactTextString(m) }
func (*SampleRequest) ProtoMessage()               {}
func (*SampleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *SampleRequest) GetFilespec() *FileSpec {
	if m != nil {
		return m.Filespec
	}
	return nil
}

func (m *SampleRequest) GetColumndesc() []*ColumnDesc {
	if m != nil {
		return m.Columndesc
	}
	return nil
}

func (m *SampleRequest) GetFragId() int32 {
	if m != nil {
		return m.FragId
	}
	return 0
}

func (m *SampleRequest) GetFragCnt() int32 {
	if m != nil {
		return m.FragCnt
	}
	return 0
}

func (m *SampleRequest) GetNrow() int32 {
	if m != nil {
		return m.Nrow
	}
	return 0
}

type DataReply struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *DataReply) Reset()                    { *m = DataReply{} }
func (m *DataReply) String() string            { return proto.CompactTextString(m) }
func (*DataReply) ProtoMessage()               {}
func (*DataReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *DataReply) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type XCol struct {
	Colname string    `protobuf:"bytes,1,opt,name=colname" json:"colname,omitempty"`
	Nrow    int32     `protobuf:"varint,2,opt,name=nrow" json:"nrow,omitempty"`
	Nullmap []bool    `protobuf:"varint,3,rep,packed,name=nullmap" json:"nullmap,omitempty"`
	Sdata   []string  `protobuf:"bytes,4,rep,name=sdata" json:"sdata,omitempty"`
	I32Data []int32   `protobuf:"varint,5,rep,packed,name=i32data" json:"i32data,omitempty"`
	I64Data []int64   `protobuf:"varint,6,rep,packed,name=i64data" json:"i64data,omitempty"`
	F32Data []float32 `protobuf:"fixed32,7,rep,packed,name=f32data" json:"f32data,omitempty"`
	F64Data []float64 `protobuf:"fixed64,8,rep,packed,name=f64data" json:"f64data,omitempty"`
}

func (m *XCol) Reset()                    { *m = XCol{} }
func (m *XCol) String() string            { return proto.CompactTextString(m) }
func (*XCol) ProtoMessage()               {}
func (*XCol) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *XCol) GetColname() string {
	if m != nil {
		return m.Colname
	}
	return ""
}

func (m *XCol) GetNrow() int32 {
	if m != nil {
		return m.Nrow
	}
	return 0
}

func (m *XCol) GetNullmap() []bool {
	if m != nil {
		return m.Nullmap
	}
	return nil
}

func (m *XCol) GetSdata() []string {
	if m != nil {
		return m.Sdata
	}
	return nil
}

func (m *XCol) GetI32Data() []int32 {
	if m != nil {
		return m.I32Data
	}
	return nil
}

func (m *XCol) GetI64Data() []int64 {
	if m != nil {
		return m.I64Data
	}
	return nil
}

func (m *XCol) GetF32Data() []float32 {
	if m != nil {
		return m.F32Data
	}
	return nil
}

func (m *XCol) GetF64Data() []float64 {
	if m != nil {
		return m.F64Data
	}
	return nil
}

type XRowSet struct {
	Tag     int32   `protobuf:"varint,1,opt,name=tag" json:"tag,omitempty"`
	Round   int32   `protobuf:"varint,2,opt,name=round" json:"round,omitempty"`
	Columns []*XCol `protobuf:"bytes,3,rep,name=columns" json:"columns,omitempty"`
}

func (m *XRowSet) Reset()                    { *m = XRowSet{} }
func (m *XRowSet) String() string            { return proto.CompactTextString(m) }
func (*XRowSet) ProtoMessage()               {}
func (*XRowSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *XRowSet) GetTag() int32 {
	if m != nil {
		return m.Tag
	}
	return 0
}

func (m *XRowSet) GetRound() int32 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *XRowSet) GetColumns() []*XCol {
	if m != nil {
		return m.Columns
	}
	return nil
}

type PluginDataReply struct {
	Errcode int32    `protobuf:"varint,1,opt,name=errcode" json:"errcode,omitempty"`
	Errmsg  string   `protobuf:"bytes,2,opt,name=errmsg" json:"errmsg,omitempty"`
	Rowset  *XRowSet `protobuf:"bytes,3,opt,name=rowset" json:"rowset,omitempty"`
}

func (m *PluginDataReply) Reset()                    { *m = PluginDataReply{} }
func (m *PluginDataReply) String() string            { return proto.CompactTextString(m) }
func (*PluginDataReply) ProtoMessage()               {}
func (*PluginDataReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *PluginDataReply) GetErrcode() int32 {
	if m != nil {
		return m.Errcode
	}
	return 0
}

func (m *PluginDataReply) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *PluginDataReply) GetRowset() *XRowSet {
	if m != nil {
		return m.Rowset
	}
	return nil
}

type SizeMetaRequest struct {
	// Which file(s)
	Filespec *FileSpec `protobuf:"bytes,1,opt,name=filespec" json:"filespec,omitempty"`
	// Table Schema
	Columndesc []*ColumnDesc `protobuf:"bytes,2,rep,name=columndesc" json:"columndesc,omitempty"`
	// Fragment
	FragId  int32 `protobuf:"varint,3,opt,name=frag_id,json=fragId" json:"frag_id,omitempty"`
	FragCnt int32 `protobuf:"varint,4,opt,name=frag_cnt,json=fragCnt" json:"frag_cnt,omitempty"`
}

func (m *SizeMetaRequest) Reset()                    { *m = SizeMetaRequest{} }
func (m *SizeMetaRequest) String() string            { return proto.CompactTextString(m) }
func (*SizeMetaRequest) ProtoMessage()               {}
func (*SizeMetaRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *SizeMetaRequest) GetFilespec() *FileSpec {
	if m != nil {
		return m.Filespec
	}
	return nil
}

func (m *SizeMetaRequest) GetColumndesc() []*ColumnDesc {
	if m != nil {
		return m.Columndesc
	}
	return nil
}

func (m *SizeMetaRequest) GetFragId() int32 {
	if m != nil {
		return m.FragId
	}
	return 0
}

func (m *SizeMetaRequest) GetFragCnt() int32 {
	if m != nil {
		return m.FragCnt
	}
	return 0
}

type SizeMetaReply struct {
	Nrow  int64 `protobuf:"varint,1,opt,name=nrow" json:"nrow,omitempty"`
	Nbyte int64 `protobuf:"varint,2,opt,name=nbyte" json:"nbyte,omitempty"`
}

func (m *SizeMetaReply) Reset()                    { *m = SizeMetaReply{} }
func (m *SizeMetaReply) String() string            { return proto.CompactTextString(m) }
func (*SizeMetaReply) ProtoMessage()               {}
func (*SizeMetaReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *SizeMetaReply) GetNrow() int64 {
	if m != nil {
		return m.Nrow
	}
	return 0
}

func (m *SizeMetaReply) GetNbyte() int64 {
	if m != nil {
		return m.Nbyte
	}
	return 0
}

type PluginSizeMetaReply struct {
	Errcode  int32          `protobuf:"varint,1,opt,name=errcode" json:"errcode,omitempty"`
	Errmsg   string         `protobuf:"bytes,2,opt,name=errmsg" json:"errmsg,omitempty"`
	Sizemeta *SizeMetaReply `protobuf:"bytes,3,opt,name=sizemeta" json:"sizemeta,omitempty"`
}

func (m *PluginSizeMetaReply) Reset()                    { *m = PluginSizeMetaReply{} }
func (m *PluginSizeMetaReply) String() string            { return proto.CompactTextString(m) }
func (*PluginSizeMetaReply) ProtoMessage()               {}
func (*PluginSizeMetaReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *PluginSizeMetaReply) GetErrcode() int32 {
	if m != nil {
		return m.Errcode
	}
	return 0
}

func (m *PluginSizeMetaReply) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *PluginSizeMetaReply) GetSizemeta() *SizeMetaReply {
	if m != nil {
		return m.Sizemeta
	}
	return nil
}

type PageData struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *PageData) Reset()                    { *m = PageData{} }
func (m *PageData) String() string            { return proto.CompactTextString(m) }
func (*PageData) ProtoMessage()               {}
func (*PageData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *PageData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type WriteRequest struct {
	Filespec   *FileSpec     `protobuf:"bytes,1,opt,name=filespec" json:"filespec,omitempty"`
	Columndesc []*ColumnDesc `protobuf:"bytes,2,rep,name=columndesc" json:"columndesc,omitempty"`
	Page       []*PageData   `protobuf:"bytes,3,rep,name=page" json:"page,omitempty"`
}

func (m *WriteRequest) Reset()                    { *m = WriteRequest{} }
func (m *WriteRequest) String() string            { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()               {}
func (*WriteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *WriteRequest) GetFilespec() *FileSpec {
	if m != nil {
		return m.Filespec
	}
	return nil
}

func (m *WriteRequest) GetColumndesc() []*ColumnDesc {
	if m != nil {
		return m.Columndesc
	}
	return nil
}

func (m *WriteRequest) GetPage() []*PageData {
	if m != nil {
		return m.Page
	}
	return nil
}

type PluginWriteRequest struct {
	Filespec   *FileSpec     `protobuf:"bytes,1,opt,name=filespec" json:"filespec,omitempty"`
	Columndesc []*ColumnDesc `protobuf:"bytes,2,rep,name=columndesc" json:"columndesc,omitempty"`
	Rowset     *XRowSet      `protobuf:"bytes,3,opt,name=rowset" json:"rowset,omitempty"`
}

func (m *PluginWriteRequest) Reset()                    { *m = PluginWriteRequest{} }
func (m *PluginWriteRequest) String() string            { return proto.CompactTextString(m) }
func (*PluginWriteRequest) ProtoMessage()               {}
func (*PluginWriteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *PluginWriteRequest) GetFilespec() *FileSpec {
	if m != nil {
		return m.Filespec
	}
	return nil
}

func (m *PluginWriteRequest) GetColumndesc() []*ColumnDesc {
	if m != nil {
		return m.Columndesc
	}
	return nil
}

func (m *PluginWriteRequest) GetRowset() *XRowSet {
	if m != nil {
		return m.Rowset
	}
	return nil
}

type WriteReply struct {
}

func (m *WriteReply) Reset()                    { *m = WriteReply{} }
func (m *WriteReply) String() string            { return proto.CompactTextString(m) }
func (*WriteReply) ProtoMessage()               {}
func (*WriteReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

type PluginWriteReply struct {
	Errcode int32  `protobuf:"varint,1,opt,name=errcode" json:"errcode,omitempty"`
	Errmsg  string `protobuf:"bytes,2,opt,name=errmsg" json:"errmsg,omitempty"`
}

func (m *PluginWriteReply) Reset()                    { *m = PluginWriteReply{} }
func (m *PluginWriteReply) String() string            { return proto.CompactTextString(m) }
func (*PluginWriteReply) ProtoMessage()               {}
func (*PluginWriteReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func (m *PluginWriteReply) GetErrcode() int32 {
	if m != nil {
		return m.Errcode
	}
	return 0
}

func (m *PluginWriteReply) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

type XMsg struct {
	Name   string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Flag   int64    `protobuf:"varint,2,opt,name=flag" json:"flag,omitempty"`
	Info   string   `protobuf:"bytes,3,opt,name=info" json:"info,omitempty"`
	Rowset *XRowSet `protobuf:"bytes,4,opt,name=rowset" json:"rowset,omitempty"`
}

func (m *XMsg) Reset()                    { *m = XMsg{} }
func (m *XMsg) String() string            { return proto.CompactTextString(m) }
func (*XMsg) ProtoMessage()               {}
func (*XMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{22} }

func (m *XMsg) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *XMsg) GetFlag() int64 {
	if m != nil {
		return m.Flag
	}
	return 0
}

func (m *XMsg) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *XMsg) GetRowset() *XRowSet {
	if m != nil {
		return m.Rowset
	}
	return nil
}

func init() {
	proto.RegisterType((*ColumnDesc)(nil), "xdrive.ColumnDesc")
	proto.RegisterType((*Filter)(nil), "xdrive.Filter")
	proto.RegisterType((*CSVSpec)(nil), "xdrive.CSVSpec")
	proto.RegisterType((*FileSpec)(nil), "xdrive.FileSpec")
	proto.RegisterType((*StringList)(nil), "xdrive.StringList")
	proto.RegisterType((*KeyValue)(nil), "xdrive.KeyValue")
	proto.RegisterType((*KeyValueList)(nil), "xdrive.KeyValueList")
	proto.RegisterType((*RmgrInfo)(nil), "xdrive.RmgrInfo")
	proto.RegisterType((*ReadRequest)(nil), "xdrive.ReadRequest")
	proto.RegisterType((*SampleRequest)(nil), "xdrive.SampleRequest")
	proto.RegisterType((*DataReply)(nil), "xdrive.DataReply")
	proto.RegisterType((*XCol)(nil), "xdrive.XCol")
	proto.RegisterType((*XRowSet)(nil), "xdrive.XRowSet")
	proto.RegisterType((*PluginDataReply)(nil), "xdrive.PluginDataReply")
	proto.RegisterType((*SizeMetaRequest)(nil), "xdrive.SizeMetaRequest")
	proto.RegisterType((*SizeMetaReply)(nil), "xdrive.SizeMetaReply")
	proto.RegisterType((*PluginSizeMetaReply)(nil), "xdrive.PluginSizeMetaReply")
	proto.RegisterType((*PageData)(nil), "xdrive.PageData")
	proto.RegisterType((*WriteRequest)(nil), "xdrive.WriteRequest")
	proto.RegisterType((*PluginWriteRequest)(nil), "xdrive.PluginWriteRequest")
	proto.RegisterType((*WriteReply)(nil), "xdrive.WriteReply")
	proto.RegisterType((*PluginWriteReply)(nil), "xdrive.PluginWriteReply")
	proto.RegisterType((*XMsg)(nil), "xdrive.XMsg")
	proto.RegisterEnum("xdrive.SpqType", SpqType_name, SpqType_value)
}

func init() { proto.RegisterFile("xdrive_data.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1075 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x56, 0x51, 0x6f, 0xe3, 0x44,
	0x10, 0xc6, 0xb1, 0xe3, 0x38, 0x93, 0xdc, 0xd5, 0xb7, 0x94, 0xc3, 0x20, 0x54, 0x22, 0x0b, 0x1d,
	0x01, 0xa1, 0x8a, 0x4b, 0xab, 0x0a, 0x1e, 0x7b, 0x49, 0x2b, 0x85, 0x4b, 0x9b, 0x6a, 0x9d, 0xbb,
	0x2b, 0x4f, 0x95, 0xcf, 0xd9, 0xb8, 0x56, 0x1d, 0xdb, 0xb5, 0x37, 0x29, 0xe9, 0x53, 0xa5, 0x43,
	0x42, 0xfc, 0x81, 0x7b, 0xe5, 0x91, 0xdf, 0xc0, 0x7f, 0xe0, 0x9f, 0xf0, 0xc8, 0x1f, 0x40, 0x3b,
	0xde, 0x4d, 0xd3, 0x8a, 0x22, 0xe0, 0xa1, 0xba, 0xb7, 0xf9, 0x66, 0xbe, 0x9d, 0xfd, 0x66, 0x76,
	0x76, 0x6d, 0x78, 0xf4, 0xc3, 0x38, 0x8f, 0xe6, 0xec, 0x64, 0xec, 0x73, 0x7f, 0x33, 0xcb, 0x53,
	0x9e, 0x12, 0xb3, 0x74, 0xb9, 0xdb, 0x00, 0xdd, 0x34, 0x9e, 0x4d, 0x93, 0x1e, 0x2b, 0x02, 0x42,
	0xc0, 0x48, 0xfc, 0x29, 0x73, 0xb4, 0x96, 0xd6, 0xae, 0x53, 0xb4, 0x85, 0x8f, 0x2f, 0x32, 0xe6,
	0x54, 0x5a, 0x5a, 0xbb, 0x4a, 0xd1, 0x76, 0x7b, 0x60, 0xee, 0x47, 0x31, 0x67, 0x39, 0x79, 0x08,
	0x95, 0x34, 0x93, 0xfc, 0x4a, 0x9a, 0x91, 0xc7, 0x60, 0x06, 0x98, 0x0f, 0xf9, 0x75, 0x2a, 0x91,
	0xc8, 0xe2, 0xe7, 0x61, 0xe1, 0xe8, 0x2d, 0x5d, 0x64, 0x16, 0xb6, 0xfb, 0x93, 0x06, 0xb5, 0xae,
	0xf7, 0xd2, 0xcb, 0x58, 0x40, 0x3e, 0x81, 0xfa, 0x98, 0xc5, 0xd1, 0x34, 0xe2, 0x2c, 0x97, 0xe9,
	0xae, 0x1d, 0xc4, 0x81, 0x5a, 0x32, 0x8b, 0xe3, 0x82, 0xe7, 0x32, 0xad, 0x82, 0x62, 0xbf, 0x53,
	0xe6, 0x8f, 0x59, 0xee, 0xe8, 0x2d, 0xad, 0x6d, 0x51, 0x89, 0xc8, 0x3a, 0x54, 0xcf, 0x67, 0x29,
	0x67, 0x8e, 0x81, 0xfc, 0x12, 0x08, 0x36, 0x2b, 0x02, 0x3f, 0x63, 0x4e, 0xb5, 0x54, 0x57, 0x22,
	0xd7, 0x07, 0x6b, 0x3f, 0x8a, 0x19, 0x2a, 0x21, 0x60, 0x64, 0x3e, 0x3f, 0x55, 0x3d, 0x10, 0xb6,
	0x58, 0x37, 0x49, 0xf3, 0xa9, 0xcf, 0x55, 0x55, 0x25, 0x22, 0x5f, 0x40, 0x2d, 0x28, 0xe6, 0x45,
	0xc6, 0x02, 0xdc, 0xbe, 0xd1, 0x59, 0xdb, 0x2c, 0xfb, 0xba, 0x29, 0xeb, 0xa2, 0x2a, 0xee, 0x6e,
	0x00, 0x78, 0x3c, 0x8f, 0x92, 0x70, 0x10, 0x15, 0x9c, 0xd8, 0xa0, 0x8b, 0x62, 0x34, 0xec, 0x86,
	0x30, 0xdd, 0x0e, 0x58, 0xcf, 0xd9, 0xe2, 0xa5, 0x1f, 0xcf, 0x98, 0x88, 0x9e, 0xb1, 0x85, 0x54,
	0x20, 0x4c, 0x51, 0xce, 0x5c, 0x84, 0xe4, 0xfe, 0x25, 0x70, 0xbf, 0x86, 0xa6, 0x5a, 0x83, 0x59,
	0x5b, 0x50, 0x39, 0x9b, 0x63, 0xd2, 0x46, 0xc7, 0x56, 0x4a, 0x14, 0x83, 0x56, 0xce, 0xe6, 0xee,
	0x5b, 0x0d, 0x2c, 0x3a, 0x0d, 0xf3, 0x7e, 0x32, 0x49, 0x45, 0x55, 0x45, 0x70, 0xca, 0x96, 0xe7,
	0x2d, 0xd1, 0x9d, 0xd5, 0xae, 0x43, 0x35, 0xc7, 0xd6, 0xe8, 0xa5, 0x08, 0x04, 0xa4, 0x0d, 0x46,
	0x90, 0x26, 0x13, 0x6c, 0x74, 0xa3, 0xb3, 0x7e, 0x7b, 0x5b, 0x21, 0x8c, 0x22, 0x83, 0x7c, 0x0c,
	0x56, 0x16, 0xcf, 0xc2, 0x28, 0x49, 0x33, 0xd9, 0xff, 0x25, 0x76, 0xff, 0xd0, 0xa0, 0x41, 0x99,
	0x3f, 0xa6, 0xec, 0x7c, 0xc6, 0x0a, 0x4e, 0xbe, 0x02, 0x6b, 0x12, 0xc5, 0x0c, 0x5b, 0xab, 0x61,
	0xe6, 0x65, 0x41, 0xea, 0xa4, 0xe8, 0x92, 0x41, 0x3a, 0x00, 0xe5, 0x9c, 0x8d, 0x59, 0x11, 0x38,
	0x15, 0x6c, 0x00, 0x59, 0x1e, 0xc5, 0x72, 0xbe, 0xe9, 0x0a, 0x8b, 0x6c, 0xa8, 0x35, 0x71, 0x54,
	0x70, 0x39, 0x97, 0x2b, 0x1e, 0xf2, 0x04, 0xcc, 0x09, 0xce, 0xb8, 0x63, 0x60, 0xbe, 0x87, 0x2b,
	0xfb, 0x73, 0x96, 0x53, 0x19, 0x25, 0x1f, 0x42, 0x6d, 0x92, 0xfb, 0xe1, 0x49, 0x34, 0xc6, 0xa2,
	0xaa, 0xd4, 0x14, 0xb0, 0x3f, 0x26, 0x1f, 0x81, 0x85, 0x81, 0x20, 0xe1, 0x8e, 0x89, 0x11, 0x24,
	0x76, 0x13, 0xee, 0xfe, 0xa6, 0xc1, 0x03, 0xcf, 0x9f, 0x66, 0x31, 0xbb, 0xbf, 0x7a, 0x57, 0x74,
	0xea, 0x77, 0xea, 0x34, 0x6e, 0xe8, 0xc4, 0xf7, 0x20, 0x4f, 0x2f, 0x64, 0x61, 0x68, 0xbb, 0x9f,
	0x42, 0xbd, 0xe7, 0x73, 0x9f, 0xb2, 0x2c, 0x5e, 0x08, 0x82, 0x78, 0x54, 0x50, 0x72, 0x93, 0xa2,
	0xed, 0xfe, 0xae, 0x81, 0x71, 0xdc, 0x4d, 0x63, 0x71, 0x6b, 0x83, 0x34, 0x5e, 0x79, 0x50, 0x14,
	0x5c, 0xe6, 0xad, 0x5c, 0xe7, 0x55, 0x77, 0x7c, 0xea, 0x67, 0x78, 0x18, 0x16, 0x55, 0x50, 0xcc,
	0x5d, 0x81, 0xbb, 0x18, 0x78, 0x48, 0x25, 0x10, 0xfc, 0x68, 0xab, 0x83, 0xfe, 0x6a, 0x4b, 0x17,
	0xaa, 0x25, 0xc4, 0xc8, 0xce, 0x36, 0x46, 0xcc, 0x96, 0xde, 0xd6, 0xa9, 0x82, 0x22, 0x32, 0x91,
	0x6b, 0x6a, 0x2d, 0xbd, 0x5d, 0xa1, 0x0a, 0x62, 0x44, 0xae, 0xb1, 0x5a, 0x7a, 0x5b, 0xa3, 0x0a,
	0xba, 0xdf, 0x43, 0xed, 0x98, 0xa6, 0x17, 0x1e, 0xc3, 0x5b, 0xcb, 0xfd, 0x10, 0x8b, 0xa9, 0x52,
	0x61, 0xe2, 0x95, 0x48, 0x67, 0xc9, 0x58, 0x56, 0x52, 0x02, 0xf2, 0x04, 0x0b, 0x9f, 0x4d, 0x93,
	0xf2, 0xbd, 0x6b, 0x74, 0x9a, 0xea, 0x6c, 0x44, 0x5f, 0xa8, 0x0a, 0xba, 0x31, 0xac, 0x1d, 0xe1,
	0x05, 0xb8, 0x6e, 0xa8, 0x03, 0x35, 0x96, 0xe7, 0x41, 0x3a, 0x66, 0x72, 0x1b, 0x05, 0xf1, 0xed,
	0xca, 0xf3, 0x69, 0x11, 0xaa, 0x5b, 0x59, 0x22, 0xf2, 0x39, 0x98, 0x79, 0x7a, 0x51, 0x30, 0x7e,
	0xfb, 0x09, 0x92, 0xaa, 0xa9, 0x0c, 0xbb, 0xbf, 0x6a, 0xb0, 0xe6, 0x45, 0x97, 0xec, 0x80, 0x89,
	0xcd, 0xde, 0xe1, 0xb1, 0x73, 0xbf, 0x85, 0x07, 0xd7, 0x42, 0xe5, 0x98, 0xe1, 0xbc, 0x08, 0x89,
	0xba, 0x9c, 0x97, 0x75, 0xa8, 0x26, 0xaf, 0x17, 0xbc, 0x7c, 0x12, 0x75, 0x5a, 0x02, 0xf7, 0x12,
	0xde, 0x2f, 0x5b, 0x7a, 0x33, 0xc1, 0x7f, 0x6f, 0xeb, 0x53, 0xb0, 0x8a, 0xe8, 0x92, 0x4d, 0x19,
	0xf7, 0x65, 0x63, 0x3f, 0x50, 0x95, 0xde, 0x48, 0x4d, 0x97, 0x34, 0x77, 0x03, 0xac, 0x23, 0x3f,
	0x64, 0xe2, 0x30, 0xff, 0xf6, 0x62, 0xbc, 0xd5, 0xa0, 0xf9, 0x2a, 0x8f, 0xf8, 0x3d, 0x5e, 0xfa,
	0xcf, 0xc4, 0xc7, 0x2c, 0x64, 0x72, 0x0c, 0x97, 0xd9, 0x95, 0x4c, 0x8a, 0x51, 0xf7, 0x17, 0x0d,
	0x48, 0xd9, 0xb5, 0x7b, 0x96, 0xf7, 0xaf, 0x67, 0xb7, 0x09, 0x20, 0xa5, 0x65, 0xf1, 0xc2, 0xed,
	0x81, 0x7d, 0x43, 0xee, 0xff, 0x3a, 0x61, 0xf7, 0x0c, 0x8c, 0xe3, 0x83, 0x22, 0xbc, 0xeb, 0xa7,
	0x67, 0x12, 0xfb, 0xa1, 0x9c, 0x2d, 0xb4, 0x85, 0x2f, 0x4a, 0x26, 0xa9, 0xfc, 0xfa, 0xa1, 0xbd,
	0x52, 0x80, 0xf1, 0x8f, 0x05, 0x7c, 0xf9, 0xa7, 0x06, 0x35, 0x2f, 0x3b, 0x1f, 0x2d, 0x32, 0x46,
	0x1a, 0x50, 0x7b, 0x71, 0xf8, 0xfc, 0x70, 0xf8, 0xea, 0xd0, 0x7e, 0x8f, 0x58, 0x60, 0x3c, 0x1b,
	0x0e, 0x07, 0xb6, 0x46, 0xea, 0x50, 0xed, 0x1f, 0x8e, 0x9e, 0xee, 0xd8, 0x15, 0x69, 0x6e, 0x75,
	0x6c, 0x5d, 0x9a, 0x3b, 0xdb, 0xb6, 0x41, 0x00, 0x4c, 0x41, 0xe8, 0x7c, 0x63, 0x57, 0x85, 0x7b,
	0x7f, 0x30, 0xdc, 0x1d, 0xd9, 0xa6, 0x70, 0xf7, 0x86, 0x2f, 0x9e, 0x0d, 0xf6, 0xec, 0x9a, 0xc8,
	0xd6, 0xf5, 0x46, 0xd4, 0xae, 0x0b, 0x42, 0x6f, 0xaf, 0xbb, 0xb3, 0x6d, 0x03, 0x12, 0xf6, 0xba,
	0x62, 0x5d, 0x83, 0x00, 0x18, 0xbd, 0xdd, 0xd1, 0x9e, 0xfd, 0xe6, 0xca, 0x20, 0x8f, 0xa0, 0x31,
	0xea, 0x1f, 0xec, 0x9d, 0x1c, 0xf4, 0x07, 0x83, 0xbe, 0x67, 0xbf, 0xb9, 0xb2, 0xc8, 0x63, 0xb0,
	0x85, 0xcb, 0x1b, 0xed, 0x1e, 0x1c, 0x29, 0xff, 0x8f, 0x57, 0xcd, 0x15, 0x6a, 0x97, 0x0e, 0x85,
	0xcb, 0xbe, 0x4d, 0x95, 0xfe, 0x75, 0xb1, 0xc3, 0x77, 0xde, 0xf0, 0xd0, 0xfe, 0xf9, 0xca, 0x79,
	0x6d, 0xe2, 0xcf, 0xe6, 0xd6, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x46, 0x8b, 0xa8, 0x81,
	0x0a, 0x00, 0x00,
}