// Code generated by protoc-gen-go. DO NOT EDIT.
// source: herald.proto

package data

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//
//Status is used to determine if the sample is untagged, tagged or in the announcment queue
type Sample_Status int32

const (
	Sample_UN_INITIALIZED Sample_Status = 0
	Sample_untagged       Sample_Status = 1
	Sample_tagged         Sample_Status = 2
	Sample_announced      Sample_Status = 3
)

var Sample_Status_name = map[int32]string{
	0: "UN_INITIALIZED",
	1: "untagged",
	2: "tagged",
	3: "announced",
}

var Sample_Status_value = map[string]int32{
	"UN_INITIALIZED": 0,
	"untagged":       1,
	"tagged":         2,
	"announced":      3,
}

func (x Sample_Status) String() string {
	return proto.EnumName(Sample_Status_name, int32(x))
}

func (Sample_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0fa959798cfd2bf0, []int{2, 0}
}

//
//Comments are used by Herald to record a message history
type Comment struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Text                 string               `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Comment) Reset()         { *m = Comment{} }
func (m *Comment) String() string { return proto.CompactTextString(m) }
func (*Comment) ProtoMessage()    {}
func (*Comment) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fa959798cfd2bf0, []int{0}
}

func (m *Comment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Comment.Unmarshal(m, b)
}
func (m *Comment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Comment.Marshal(b, m, deterministic)
}
func (m *Comment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Comment.Merge(m, src)
}
func (m *Comment) XXX_Size() int {
	return xxx_messageInfo_Comment.Size(m)
}
func (m *Comment) XXX_DiscardUnknown() {
	xxx_messageInfo_Comment.DiscardUnknown(m)
}

var xxx_messageInfo_Comment proto.InternalMessageInfo

func (m *Comment) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Comment) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

//
//Experiment is used to describe a Nanopore sequencing run
type Experiment struct {
	Created              *timestamp.Timestamp `protobuf:"bytes,1,opt,name=created,proto3" json:"created,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	History              []*Comment           `protobuf:"bytes,3,rep,name=history,proto3" json:"history,omitempty"`
	OutputDirectory      string               `protobuf:"bytes,4,opt,name=output_directory,json=outputDirectory,proto3" json:"output_directory,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Experiment) Reset()         { *m = Experiment{} }
func (m *Experiment) String() string { return proto.CompactTextString(m) }
func (*Experiment) ProtoMessage()    {}
func (*Experiment) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fa959798cfd2bf0, []int{1}
}

func (m *Experiment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Experiment.Unmarshal(m, b)
}
func (m *Experiment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Experiment.Marshal(b, m, deterministic)
}
func (m *Experiment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Experiment.Merge(m, src)
}
func (m *Experiment) XXX_Size() int {
	return xxx_messageInfo_Experiment.Size(m)
}
func (m *Experiment) XXX_DiscardUnknown() {
	xxx_messageInfo_Experiment.DiscardUnknown(m)
}

var xxx_messageInfo_Experiment proto.InternalMessageInfo

func (m *Experiment) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *Experiment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Experiment) GetHistory() []*Comment {
	if m != nil {
		return m.History
	}
	return nil
}

func (m *Experiment) GetOutputDirectory() string {
	if m != nil {
		return m.OutputDirectory
	}
	return ""
}

//
//Sample is used to describe a biological sample which is being sequenced
type Sample struct {
	Created              *timestamp.Timestamp `protobuf:"bytes,1,opt,name=created,proto3" json:"created,omitempty"`
	Label                string               `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	History              []*Comment           `protobuf:"bytes,3,rep,name=history,proto3" json:"history,omitempty"`
	Experiment           *Experiment          `protobuf:"bytes,4,opt,name=experiment,proto3" json:"experiment,omitempty"`
	Status               Sample_Status        `protobuf:"varint,5,opt,name=status,proto3,enum=data.Sample_Status" json:"status,omitempty"`
	Barcode              int32                `protobuf:"varint,6,opt,name=barcode,proto3" json:"barcode,omitempty"`
	Tags                 *Sample_Tags         `protobuf:"bytes,7,opt,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Sample) Reset()         { *m = Sample{} }
func (m *Sample) String() string { return proto.CompactTextString(m) }
func (*Sample) ProtoMessage()    {}
func (*Sample) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fa959798cfd2bf0, []int{2}
}

func (m *Sample) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sample.Unmarshal(m, b)
}
func (m *Sample) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sample.Marshal(b, m, deterministic)
}
func (m *Sample) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sample.Merge(m, src)
}
func (m *Sample) XXX_Size() int {
	return xxx_messageInfo_Sample.Size(m)
}
func (m *Sample) XXX_DiscardUnknown() {
	xxx_messageInfo_Sample.DiscardUnknown(m)
}

var xxx_messageInfo_Sample proto.InternalMessageInfo

func (m *Sample) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *Sample) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Sample) GetHistory() []*Comment {
	if m != nil {
		return m.History
	}
	return nil
}

func (m *Sample) GetExperiment() *Experiment {
	if m != nil {
		return m.Experiment
	}
	return nil
}

func (m *Sample) GetStatus() Sample_Status {
	if m != nil {
		return m.Status
	}
	return Sample_UN_INITIALIZED
}

func (m *Sample) GetBarcode() int32 {
	if m != nil {
		return m.Barcode
	}
	return 0
}

func (m *Sample) GetTags() *Sample_Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

//
//Tags are used to indicate what processes to run on a sample
type Sample_Tags struct {
	Sequence             bool     `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Rampart              bool     `protobuf:"varint,2,opt,name=rampart,proto3" json:"rampart,omitempty"`
	ArticPipeline        bool     `protobuf:"varint,3,opt,name=artic_pipeline,json=articPipeline,proto3" json:"artic_pipeline,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sample_Tags) Reset()         { *m = Sample_Tags{} }
func (m *Sample_Tags) String() string { return proto.CompactTextString(m) }
func (*Sample_Tags) ProtoMessage()    {}
func (*Sample_Tags) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fa959798cfd2bf0, []int{2, 0}
}

func (m *Sample_Tags) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sample_Tags.Unmarshal(m, b)
}
func (m *Sample_Tags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sample_Tags.Marshal(b, m, deterministic)
}
func (m *Sample_Tags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sample_Tags.Merge(m, src)
}
func (m *Sample_Tags) XXX_Size() int {
	return xxx_messageInfo_Sample_Tags.Size(m)
}
func (m *Sample_Tags) XXX_DiscardUnknown() {
	xxx_messageInfo_Sample_Tags.DiscardUnknown(m)
}

var xxx_messageInfo_Sample_Tags proto.InternalMessageInfo

func (m *Sample_Tags) GetSequence() bool {
	if m != nil {
		return m.Sequence
	}
	return false
}

func (m *Sample_Tags) GetRampart() bool {
	if m != nil {
		return m.Rampart
	}
	return false
}

func (m *Sample_Tags) GetArticPipeline() bool {
	if m != nil {
		return m.ArticPipeline
	}
	return false
}

func init() {
	proto.RegisterEnum("data.Sample_Status", Sample_Status_name, Sample_Status_value)
	proto.RegisterType((*Comment)(nil), "data.Comment")
	proto.RegisterType((*Experiment)(nil), "data.Experiment")
	proto.RegisterType((*Sample)(nil), "data.Sample")
	proto.RegisterType((*Sample_Tags)(nil), "data.Sample.Tags")
}

func init() { proto.RegisterFile("herald.proto", fileDescriptor_0fa959798cfd2bf0) }

var fileDescriptor_0fa959798cfd2bf0 = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0xb5, 0x63, 0x27, 0x93, 0x26, 0x98, 0x85, 0x83, 0x95, 0x0b, 0x56, 0xa4, 0x0a, 0x23,
	0x24, 0x17, 0x05, 0x0e, 0x5c, 0x11, 0xcd, 0x21, 0x12, 0xaa, 0xd0, 0x36, 0x08, 0x89, 0x4b, 0x34,
	0xb1, 0x07, 0xd7, 0x92, 0xed, 0x35, 0xeb, 0xb5, 0x54, 0x7e, 0x89, 0x3f, 0xe1, 0xaf, 0x90, 0x77,
	0xbd, 0x29, 0xdc, 0xaa, 0xde, 0xe6, 0xbd, 0x79, 0xf3, 0xf6, 0xcd, 0xd8, 0x70, 0x7e, 0x4b, 0x12,
	0xab, 0x3c, 0x6d, 0xa5, 0x50, 0x82, 0x79, 0x39, 0x2a, 0x5c, 0xbd, 0x2c, 0x84, 0x28, 0x2a, 0xba,
	0xd4, 0xdc, 0xb1, 0xff, 0x71, 0xa9, 0xca, 0x9a, 0x3a, 0x85, 0x75, 0x6b, 0x64, 0xeb, 0x6f, 0x10,
	0x7c, 0x12, 0x75, 0x4d, 0x8d, 0x62, 0x1f, 0x60, 0x76, 0xea, 0x46, 0x4e, 0xec, 0x24, 0xf3, 0xcd,
	0x2a, 0x35, 0xf3, 0xa9, 0x9d, 0x4f, 0xf7, 0x56, 0xc1, 0xef, 0xc5, 0x8c, 0x81, 0xa7, 0xe8, 0x4e,
	0x45, 0x67, 0xb1, 0x93, 0xcc, 0xb8, 0xae, 0xd7, 0xbf, 0x1d, 0x80, 0xed, 0x5d, 0x4b, 0xb2, 0xd4,
	0xe6, 0xef, 0x21, 0xc8, 0x24, 0xa1, 0xa2, 0xfc, 0x01, 0xd6, 0x56, 0x3a, 0x18, 0x37, 0x58, 0x93,
	0x35, 0x1e, 0x6a, 0xf6, 0x0a, 0x82, 0xdb, 0xb2, 0x53, 0x42, 0xfe, 0x8a, 0xdc, 0xd8, 0x4d, 0xe6,
	0x9b, 0x45, 0x3a, 0xac, 0x9a, 0x8e, 0x6b, 0x70, 0xdb, 0x65, 0xaf, 0x21, 0x14, 0xbd, 0x6a, 0x7b,
	0x75, 0xc8, 0x4b, 0x49, 0x99, 0x9e, 0xf0, 0xb4, 0xd1, 0x53, 0xc3, 0x5f, 0x59, 0x7a, 0xfd, 0xc7,
	0x05, 0xff, 0x06, 0xeb, 0xb6, 0xa2, 0x47, 0x06, 0x7d, 0x01, 0x93, 0x0a, 0x8f, 0x54, 0x8d, 0x49,
	0x0d, 0x78, 0x78, 0xd4, 0xb7, 0x00, 0x74, 0xba, 0x95, 0x0e, 0x39, 0xdf, 0x84, 0x46, 0x7b, 0x7f,
	0x43, 0xfe, 0x8f, 0x86, 0xbd, 0x01, 0xbf, 0x53, 0xa8, 0xfa, 0x2e, 0x9a, 0xc4, 0x4e, 0xb2, 0xdc,
	0x3c, 0x37, 0x6a, 0xb3, 0x44, 0x7a, 0xa3, 0x5b, 0x7c, 0x94, 0xb0, 0x08, 0x82, 0x23, 0xca, 0x4c,
	0xe4, 0x14, 0xf9, 0xb1, 0x93, 0x4c, 0xb8, 0x85, 0xec, 0x02, 0x3c, 0x85, 0x45, 0x17, 0x05, 0xfa,
	0xc9, 0x67, 0xff, 0x99, 0xec, 0xb1, 0xe8, 0xb8, 0x6e, 0xaf, 0x32, 0xf0, 0x06, 0xc4, 0x56, 0x30,
	0xed, 0xe8, 0x67, 0x4f, 0x4d, 0x46, 0xfa, 0x3a, 0x53, 0x7e, 0xc2, 0xc3, 0x23, 0x12, 0xeb, 0x16,
	0xa5, 0xf9, 0x0f, 0xa6, 0xdc, 0x42, 0x76, 0x01, 0x4b, 0x94, 0xaa, 0xcc, 0x0e, 0x6d, 0xd9, 0x52,
	0x55, 0x36, 0x14, 0xb9, 0x5a, 0xb0, 0xd0, 0xec, 0x97, 0x91, 0x5c, 0x6f, 0xc1, 0x37, 0xb9, 0x19,
	0x83, 0xe5, 0xd7, 0xeb, 0xc3, 0xee, 0x7a, 0xb7, 0xdf, 0x7d, 0xfc, 0xbc, 0xfb, 0xbe, 0xbd, 0x0a,
	0x9f, 0xb0, 0x73, 0x98, 0xf6, 0x8d, 0xc2, 0xa2, 0xa0, 0x3c, 0x74, 0x18, 0x80, 0x3f, 0xd6, 0x67,
	0x6c, 0x01, 0x33, 0x6c, 0x1a, 0xd1, 0x37, 0x19, 0xe5, 0xa1, 0x7b, 0xf4, 0xf5, 0x77, 0x7a, 0xf7,
	0x37, 0x00, 0x00, 0xff, 0xff, 0x40, 0x30, 0xf9, 0xd0, 0x0f, 0x03, 0x00, 0x00,
}