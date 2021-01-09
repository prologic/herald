// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.3
// source: herald.proto

package services

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

//
//Status is used to determine if runs/samples have tagged processes, or if they have been announced via the message server
type Status int32

const (
	Status_UN_INITIALIZED Status = 0
	Status_untagged       Status = 1
	Status_tagged         Status = 2
	Status_announced      Status = 3
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "UN_INITIALIZED",
		1: "untagged",
		2: "tagged",
		3: "announced",
	}
	Status_value = map[string]int32{
		"UN_INITIALIZED": 0,
		"untagged":       1,
		"tagged":         2,
		"announced":      3,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_herald_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_herald_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_herald_proto_rawDescGZIP(), []int{0}
}

//
//Comments are used to record a message history
type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Text      string               `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_herald_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_herald_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_herald_proto_rawDescGZIP(), []int{0}
}

func (x *Comment) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Comment) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

//
//User is used to identify the owner of HeraldData and is required for announcements
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Created *timestamp.Timestamp `protobuf:"bytes,1,opt,name=created,proto3" json:"created,omitempty"`
	Name    string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email   string               `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_herald_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_herald_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_herald_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetCreated() *timestamp.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

//
//Config is used to describe a Herald instance
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Created    *timestamp.Timestamp `protobuf:"bytes,1,opt,name=created,proto3" json:"created,omitempty"`
	Filepath   string               `protobuf:"bytes,2,opt,name=filepath,proto3" json:"filepath,omitempty"`     // filepath to config
	Fileformat string               `protobuf:"bytes,3,opt,name=fileformat,proto3" json:"fileformat,omitempty"` // the fileformat of the config on disk
	Version    string               `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`       // version of Herald used
	User       *User                `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`             // user details
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_herald_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_herald_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_herald_proto_rawDescGZIP(), []int{2}
}

func (x *Config) GetCreated() *timestamp.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *Config) GetFilepath() string {
	if x != nil {
		return x.Filepath
	}
	return ""
}

func (x *Config) GetFileformat() string {
	if x != nil {
		return x.Fileformat
	}
	return ""
}

func (x *Config) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Config) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

//
//HeraldData is the base type, used by both Run and Sample
type HeraldData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Created      *timestamp.Timestamp `protobuf:"bytes,1,opt,name=created,proto3" json:"created,omitempty"`
	Label        string               `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`                                                                                        // the run or sample name
	History      []*Comment           `protobuf:"bytes,4,rep,name=history,proto3" json:"history,omitempty"`                                                                                    // describes the history of the run
	Status       Status               `protobuf:"varint,5,opt,name=status,proto3,enum=services.Status" json:"status,omitempty"`                                                                // describes if untagged/tagged/announced
	Tags         map[string]bool      `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` // tagged services and their complete status (true=complete, false=incomplete)
	RequestOrder []string             `protobuf:"bytes,7,rep,name=requestOrder,proto3" json:"requestOrder,omitempty"`                                                                          // the order to send requests to the tagged services
}

func (x *HeraldData) Reset() {
	*x = HeraldData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_herald_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeraldData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeraldData) ProtoMessage() {}

func (x *HeraldData) ProtoReflect() protoreflect.Message {
	mi := &file_herald_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeraldData.ProtoReflect.Descriptor instead.
func (*HeraldData) Descriptor() ([]byte, []int) {
	return file_herald_proto_rawDescGZIP(), []int{3}
}

func (x *HeraldData) GetCreated() *timestamp.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *HeraldData) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *HeraldData) GetHistory() []*Comment {
	if x != nil {
		return x.History
	}
	return nil
}

func (x *HeraldData) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_UN_INITIALIZED
}

func (x *HeraldData) GetTags() map[string]bool {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *HeraldData) GetRequestOrder() []string {
	if x != nil {
		return x.RequestOrder
	}
	return nil
}

//
//Run is used to describe a Nanopore sequencing run
type Run struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata             *HeraldData `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`                         // the base info
	OutputDirectory      string      `protobuf:"bytes,2,opt,name=outputDirectory,proto3" json:"outputDirectory,omitempty"`           // where the run is stored
	Fast5OutputDirectory string      `protobuf:"bytes,3,opt,name=fast5OutputDirectory,proto3" json:"fast5OutputDirectory,omitempty"` // where the run fast5 data is stored
	FastqOutputDirectory string      `protobuf:"bytes,4,opt,name=fastqOutputDirectory,proto3" json:"fastqOutputDirectory,omitempty"` // where the run fastq data is stored
}

func (x *Run) Reset() {
	*x = Run{}
	if protoimpl.UnsafeEnabled {
		mi := &file_herald_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Run) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Run) ProtoMessage() {}

func (x *Run) ProtoReflect() protoreflect.Message {
	mi := &file_herald_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Run.ProtoReflect.Descriptor instead.
func (*Run) Descriptor() ([]byte, []int) {
	return file_herald_proto_rawDescGZIP(), []int{4}
}

func (x *Run) GetMetadata() *HeraldData {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Run) GetOutputDirectory() string {
	if x != nil {
		return x.OutputDirectory
	}
	return ""
}

func (x *Run) GetFast5OutputDirectory() string {
	if x != nil {
		return x.Fast5OutputDirectory
	}
	return ""
}

func (x *Run) GetFastqOutputDirectory() string {
	if x != nil {
		return x.FastqOutputDirectory
	}
	return ""
}

//
//Sample is used to describe a biological sample which is being sequenced as part of a Run
type Sample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata  *HeraldData `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"` // the base info
	ParentRun string      `protobuf:"bytes,2,opt,name=parentRun,proto3" json:"parentRun,omitempty"`
	Barcode   int32       `protobuf:"varint,3,opt,name=barcode,proto3" json:"barcode,omitempty"`
}

func (x *Sample) Reset() {
	*x = Sample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_herald_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sample) ProtoMessage() {}

func (x *Sample) ProtoReflect() protoreflect.Message {
	mi := &file_herald_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sample.ProtoReflect.Descriptor instead.
func (*Sample) Descriptor() ([]byte, []int) {
	return file_herald_proto_rawDescGZIP(), []int{5}
}

func (x *Sample) GetMetadata() *HeraldData {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Sample) GetParentRun() string {
	if x != nil {
		return x.ParentRun
	}
	return ""
}

func (x *Sample) GetBarcode() int32 {
	if x != nil {
		return x.Barcode
	}
	return 0
}

var File_herald_proto protoreflect.FileDescriptor

var file_herald_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x68, 0x65, 0x72, 0x61, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x07, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x22, 0x66, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0xb8, 0x01, 0x0a, 0x06, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x69, 0x6c,
	0x65, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x22, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0xc0, 0x02, 0x0a, 0x0a, 0x48, 0x65, 0x72, 0x61, 0x6c, 0x64,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x12, 0x2b, 0x0a, 0x07, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x28, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x48, 0x65, 0x72, 0x61, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x61, 0x67, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a,
	0x37, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xc9, 0x01, 0x0a, 0x03, 0x52, 0x75, 0x6e,
	0x12, 0x30, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x48, 0x65,
	0x72, 0x61, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x28, 0x0a, 0x0f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x32, 0x0a, 0x14,
	0x66, 0x61, 0x73, 0x74, 0x35, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x66, 0x61, 0x73, 0x74,
	0x35, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x32, 0x0a, 0x14, 0x66, 0x61, 0x73, 0x74, 0x71, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14,
	0x66, 0x61, 0x73, 0x74, 0x71, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x22, 0x72, 0x0a, 0x06, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x30,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x48, 0x65, 0x72, 0x61,
	0x6c, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x75, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x75, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2a, 0x45, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x5f, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c,
	0x49, 0x5a, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x75, 0x6e, 0x74, 0x61, 0x67, 0x67,
	0x65, 0x64, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x74, 0x61, 0x67, 0x67, 0x65, 0x64, 0x10, 0x02,
	0x12, 0x0d, 0x0a, 0x09, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x64, 0x10, 0x03, 0x42,
	0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_herald_proto_rawDescOnce sync.Once
	file_herald_proto_rawDescData = file_herald_proto_rawDesc
)

func file_herald_proto_rawDescGZIP() []byte {
	file_herald_proto_rawDescOnce.Do(func() {
		file_herald_proto_rawDescData = protoimpl.X.CompressGZIP(file_herald_proto_rawDescData)
	})
	return file_herald_proto_rawDescData
}

var file_herald_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_herald_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_herald_proto_goTypes = []interface{}{
	(Status)(0),                 // 0: services.Status
	(*Comment)(nil),             // 1: services.Comment
	(*User)(nil),                // 2: services.User
	(*Config)(nil),              // 3: services.Config
	(*HeraldData)(nil),          // 4: services.HeraldData
	(*Run)(nil),                 // 5: services.Run
	(*Sample)(nil),              // 6: services.Sample
	nil,                         // 7: services.HeraldData.TagsEntry
	(*timestamp.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_herald_proto_depIdxs = []int32{
	8,  // 0: services.Comment.timestamp:type_name -> google.protobuf.Timestamp
	8,  // 1: services.User.created:type_name -> google.protobuf.Timestamp
	8,  // 2: services.Config.created:type_name -> google.protobuf.Timestamp
	2,  // 3: services.Config.user:type_name -> services.User
	8,  // 4: services.HeraldData.created:type_name -> google.protobuf.Timestamp
	1,  // 5: services.HeraldData.history:type_name -> services.Comment
	0,  // 6: services.HeraldData.status:type_name -> services.Status
	7,  // 7: services.HeraldData.tags:type_name -> services.HeraldData.TagsEntry
	4,  // 8: services.Run.metadata:type_name -> services.HeraldData
	4,  // 9: services.Sample.metadata:type_name -> services.HeraldData
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_herald_proto_init() }
func file_herald_proto_init() {
	if File_herald_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_herald_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_herald_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_herald_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_herald_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeraldData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_herald_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Run); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_herald_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sample); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_herald_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_herald_proto_goTypes,
		DependencyIndexes: file_herald_proto_depIdxs,
		EnumInfos:         file_herald_proto_enumTypes,
		MessageInfos:      file_herald_proto_msgTypes,
	}.Build()
	File_herald_proto = out.File
	file_herald_proto_rawDesc = nil
	file_herald_proto_goTypes = nil
	file_herald_proto_depIdxs = nil
}
