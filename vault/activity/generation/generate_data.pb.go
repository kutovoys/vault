// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: vault/activity/generation/generate_data.proto

package generation

import (
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

type WriteOptions int32

const (
	WriteOptions_WRITE_UNKNOWN             WriteOptions = 0
	WriteOptions_WRITE_PRECOMPUTED_QUERIES WriteOptions = 1
	WriteOptions_WRITE_DISTINCT_CLIENTS    WriteOptions = 2
	WriteOptions_WRITE_ENTITIES            WriteOptions = 3
	WriteOptions_WRITE_DIRECT_TOKENS       WriteOptions = 4
	WriteOptions_WRITE_INTENT_LOGS         WriteOptions = 5
)

// Enum value maps for WriteOptions.
var (
	WriteOptions_name = map[int32]string{
		0: "WRITE_UNKNOWN",
		1: "WRITE_PRECOMPUTED_QUERIES",
		2: "WRITE_DISTINCT_CLIENTS",
		3: "WRITE_ENTITIES",
		4: "WRITE_DIRECT_TOKENS",
		5: "WRITE_INTENT_LOGS",
	}
	WriteOptions_value = map[string]int32{
		"WRITE_UNKNOWN":             0,
		"WRITE_PRECOMPUTED_QUERIES": 1,
		"WRITE_DISTINCT_CLIENTS":    2,
		"WRITE_ENTITIES":            3,
		"WRITE_DIRECT_TOKENS":       4,
		"WRITE_INTENT_LOGS":         5,
	}
)

func (x WriteOptions) Enum() *WriteOptions {
	p := new(WriteOptions)
	*p = x
	return p
}

func (x WriteOptions) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WriteOptions) Descriptor() protoreflect.EnumDescriptor {
	return file_vault_activity_generation_generate_data_proto_enumTypes[0].Descriptor()
}

func (WriteOptions) Type() protoreflect.EnumType {
	return &file_vault_activity_generation_generate_data_proto_enumTypes[0]
}

func (x WriteOptions) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WriteOptions.Descriptor instead.
func (WriteOptions) EnumDescriptor() ([]byte, []int) {
	return file_vault_activity_generation_generate_data_proto_rawDescGZIP(), []int{0}
}

type ActivityLogMockInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Write []WriteOptions `protobuf:"varint,1,rep,packed,name=write,proto3,enum=generation.WriteOptions" json:"write,omitempty"`
	Data  []*Data        `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ActivityLogMockInput) Reset() {
	*x = ActivityLogMockInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_activity_generation_generate_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityLogMockInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityLogMockInput) ProtoMessage() {}

func (x *ActivityLogMockInput) ProtoReflect() protoreflect.Message {
	mi := &file_vault_activity_generation_generate_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityLogMockInput.ProtoReflect.Descriptor instead.
func (*ActivityLogMockInput) Descriptor() ([]byte, []int) {
	return file_vault_activity_generation_generate_data_proto_rawDescGZIP(), []int{0}
}

func (x *ActivityLogMockInput) GetWrite() []WriteOptions {
	if x != nil {
		return x.Write
	}
	return nil
}

func (x *ActivityLogMockInput) GetData() []*Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Month:
	//
	//	*Data_CurrentMonth
	//	*Data_MonthsAgo
	Month isData_Month `protobuf_oneof:"month"`
	// Types that are assignable to Clients:
	//
	//	*Data_All
	//	*Data_Segments
	Clients             isData_Clients `protobuf_oneof:"clients"`
	EmptySegmentIndexes []int32        `protobuf:"varint,5,rep,packed,name=empty_segment_indexes,json=emptySegmentIndexes,proto3" json:"empty_segment_indexes,omitempty"`
	SkipSegmentIndexes  []int32        `protobuf:"varint,6,rep,packed,name=skip_segment_indexes,json=skipSegmentIndexes,proto3" json:"skip_segment_indexes,omitempty"`
	NumSegments         int32          `protobuf:"varint,7,opt,name=num_segments,json=numSegments,proto3" json:"num_segments,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_activity_generation_generate_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_vault_activity_generation_generate_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_vault_activity_generation_generate_data_proto_rawDescGZIP(), []int{1}
}

func (m *Data) GetMonth() isData_Month {
	if m != nil {
		return m.Month
	}
	return nil
}

func (x *Data) GetCurrentMonth() bool {
	if x, ok := x.GetMonth().(*Data_CurrentMonth); ok {
		return x.CurrentMonth
	}
	return false
}

func (x *Data) GetMonthsAgo() int32 {
	if x, ok := x.GetMonth().(*Data_MonthsAgo); ok {
		return x.MonthsAgo
	}
	return 0
}

func (m *Data) GetClients() isData_Clients {
	if m != nil {
		return m.Clients
	}
	return nil
}

func (x *Data) GetAll() *Clients {
	if x, ok := x.GetClients().(*Data_All); ok {
		return x.All
	}
	return nil
}

func (x *Data) GetSegments() *Segments {
	if x, ok := x.GetClients().(*Data_Segments); ok {
		return x.Segments
	}
	return nil
}

func (x *Data) GetEmptySegmentIndexes() []int32 {
	if x != nil {
		return x.EmptySegmentIndexes
	}
	return nil
}

func (x *Data) GetSkipSegmentIndexes() []int32 {
	if x != nil {
		return x.SkipSegmentIndexes
	}
	return nil
}

func (x *Data) GetNumSegments() int32 {
	if x != nil {
		return x.NumSegments
	}
	return 0
}

type isData_Month interface {
	isData_Month()
}

type Data_CurrentMonth struct {
	CurrentMonth bool `protobuf:"varint,1,opt,name=current_month,json=currentMonth,proto3,oneof"`
}

type Data_MonthsAgo struct {
	MonthsAgo int32 `protobuf:"varint,2,opt,name=months_ago,json=monthsAgo,proto3,oneof"`
}

func (*Data_CurrentMonth) isData_Month() {}

func (*Data_MonthsAgo) isData_Month() {}

type isData_Clients interface {
	isData_Clients()
}

type Data_All struct {
	All *Clients `protobuf:"bytes,3,opt,name=all,proto3,oneof"` // you can’t have repeated fields in a oneof, which is why these are separate message types
}

type Data_Segments struct {
	Segments *Segments `protobuf:"bytes,4,opt,name=segments,proto3,oneof"`
}

func (*Data_All) isData_Clients() {}

func (*Data_Segments) isData_Clients() {}

type Segments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Segments []*Segment `protobuf:"bytes,1,rep,name=segments,proto3" json:"segments,omitempty"`
}

func (x *Segments) Reset() {
	*x = Segments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_activity_generation_generate_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Segments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Segments) ProtoMessage() {}

func (x *Segments) ProtoReflect() protoreflect.Message {
	mi := &file_vault_activity_generation_generate_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Segments.ProtoReflect.Descriptor instead.
func (*Segments) Descriptor() ([]byte, []int) {
	return file_vault_activity_generation_generate_data_proto_rawDescGZIP(), []int{2}
}

func (x *Segments) GetSegments() []*Segment {
	if x != nil {
		return x.Segments
	}
	return nil
}

type Segment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SegmentIndex *int32   `protobuf:"varint,1,opt,name=segment_index,json=segmentIndex,proto3,oneof" json:"segment_index,omitempty"`
	Clients      *Clients `protobuf:"bytes,2,opt,name=clients,proto3" json:"clients,omitempty"`
}

func (x *Segment) Reset() {
	*x = Segment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_activity_generation_generate_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Segment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Segment) ProtoMessage() {}

func (x *Segment) ProtoReflect() protoreflect.Message {
	mi := &file_vault_activity_generation_generate_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Segment.ProtoReflect.Descriptor instead.
func (*Segment) Descriptor() ([]byte, []int) {
	return file_vault_activity_generation_generate_data_proto_rawDescGZIP(), []int{3}
}

func (x *Segment) GetSegmentIndex() int32 {
	if x != nil && x.SegmentIndex != nil {
		return *x.SegmentIndex
	}
	return 0
}

func (x *Segment) GetClients() *Clients {
	if x != nil {
		return x.Clients
	}
	return nil
}

type Clients struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clients []*Client `protobuf:"bytes,1,rep,name=clients,proto3" json:"clients,omitempty"`
}

func (x *Clients) Reset() {
	*x = Clients{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_activity_generation_generate_data_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Clients) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Clients) ProtoMessage() {}

func (x *Clients) ProtoReflect() protoreflect.Message {
	mi := &file_vault_activity_generation_generate_data_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Clients.ProtoReflect.Descriptor instead.
func (*Clients) Descriptor() ([]byte, []int) {
	return file_vault_activity_generation_generate_data_proto_rawDescGZIP(), []int{4}
}

func (x *Clients) GetClients() []*Client {
	if x != nil {
		return x.Clients
	}
	return nil
}

type Client struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Count             int32  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	TimesSeen         int32  `protobuf:"varint,3,opt,name=times_seen,json=timesSeen,proto3" json:"times_seen,omitempty"`
	Repeated          bool   `protobuf:"varint,4,opt,name=repeated,proto3" json:"repeated,omitempty"`
	RepeatedFromMonth int32  `protobuf:"varint,5,opt,name=repeated_from_month,json=repeatedFromMonth,proto3" json:"repeated_from_month,omitempty"`
	Namespace         string `protobuf:"bytes,6,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Mount             string `protobuf:"bytes,7,opt,name=mount,proto3" json:"mount,omitempty"`
	NonEntity         bool   `protobuf:"varint,8,opt,name=non_entity,json=nonEntity,proto3" json:"non_entity,omitempty"`
}

func (x *Client) Reset() {
	*x = Client{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_activity_generation_generate_data_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Client) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client) ProtoMessage() {}

func (x *Client) ProtoReflect() protoreflect.Message {
	mi := &file_vault_activity_generation_generate_data_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Client.ProtoReflect.Descriptor instead.
func (*Client) Descriptor() ([]byte, []int) {
	return file_vault_activity_generation_generate_data_proto_rawDescGZIP(), []int{5}
}

func (x *Client) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Client) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Client) GetTimesSeen() int32 {
	if x != nil {
		return x.TimesSeen
	}
	return 0
}

func (x *Client) GetRepeated() bool {
	if x != nil {
		return x.Repeated
	}
	return false
}

func (x *Client) GetRepeatedFromMonth() int32 {
	if x != nil {
		return x.RepeatedFromMonth
	}
	return 0
}

func (x *Client) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Client) GetMount() string {
	if x != nil {
		return x.Mount
	}
	return ""
}

func (x *Client) GetNonEntity() bool {
	if x != nil {
		return x.NonEntity
	}
	return false
}

var File_vault_activity_generation_generate_data_proto protoreflect.FileDescriptor

var file_vault_activity_generation_generate_data_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6c, 0x0a, 0x14, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x67, 0x4d, 0x6f, 0x63, 0x6b, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x2e, 0x0a, 0x05, 0x77, 0x72, 0x69, 0x74, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0e, 0x32, 0x18, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x05, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xc8, 0x02, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x25, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x6f,
	0x6e, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0c, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x1f, 0x0a, 0x0a, 0x6d, 0x6f, 0x6e,
	0x74, 0x68, 0x73, 0x5f, 0x61, 0x67, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52,
	0x09, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x73, 0x41, 0x67, 0x6f, 0x12, 0x27, 0x0a, 0x03, 0x61, 0x6c,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x48, 0x01, 0x52, 0x03,
	0x61, 0x6c, 0x6c, 0x12, 0x32, 0x0a, 0x08, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x48, 0x01, 0x52, 0x08, 0x73,
	0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x05, 0x52, 0x13, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x65, 0x67,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x73,
	0x6b, 0x69, 0x70, 0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x05, 0x52, 0x12, 0x73, 0x6b, 0x69, 0x70, 0x53,
	0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x12, 0x21, 0x0a,
	0x0c, 0x6e, 0x75, 0x6d, 0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x6e, 0x75, 0x6d, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x42, 0x07, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x22, 0x3b, 0x0a, 0x08, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x2f, 0x0a, 0x08, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x22, 0x74, 0x0a, 0x07, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x0d,
	0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x07, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x37, 0x0a, 0x07, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x2c, 0x0a, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x22, 0xec, 0x01, 0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x65, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x53, 0x65, 0x65, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x2e, 0x0a, 0x13,
	0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x6d, 0x6f,
	0x6e, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x72, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6e, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2a,
	0xa0, 0x01, 0x0a, 0x0c, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x11, 0x0a, 0x0d, 0x57, 0x52, 0x49, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x57, 0x52, 0x49, 0x54, 0x45, 0x5f, 0x50, 0x52, 0x45,
	0x43, 0x4f, 0x4d, 0x50, 0x55, 0x54, 0x45, 0x44, 0x5f, 0x51, 0x55, 0x45, 0x52, 0x49, 0x45, 0x53,
	0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x57, 0x52, 0x49, 0x54, 0x45, 0x5f, 0x44, 0x49, 0x53, 0x54,
	0x49, 0x4e, 0x43, 0x54, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x53, 0x10, 0x02, 0x12, 0x12,
	0x0a, 0x0e, 0x57, 0x52, 0x49, 0x54, 0x45, 0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x49, 0x45, 0x53,
	0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x57, 0x52, 0x49, 0x54, 0x45, 0x5f, 0x44, 0x49, 0x52, 0x45,
	0x43, 0x54, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x53, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x57,
	0x52, 0x49, 0x54, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x4c, 0x4f, 0x47, 0x53,
	0x10, 0x05, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x76, 0x61, 0x75, 0x6c, 0x74,
	0x2f, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_vault_activity_generation_generate_data_proto_rawDescOnce sync.Once
	file_vault_activity_generation_generate_data_proto_rawDescData = file_vault_activity_generation_generate_data_proto_rawDesc
)

func file_vault_activity_generation_generate_data_proto_rawDescGZIP() []byte {
	file_vault_activity_generation_generate_data_proto_rawDescOnce.Do(func() {
		file_vault_activity_generation_generate_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_vault_activity_generation_generate_data_proto_rawDescData)
	})
	return file_vault_activity_generation_generate_data_proto_rawDescData
}

var file_vault_activity_generation_generate_data_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_vault_activity_generation_generate_data_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_vault_activity_generation_generate_data_proto_goTypes = []interface{}{
	(WriteOptions)(0),            // 0: generation.WriteOptions
	(*ActivityLogMockInput)(nil), // 1: generation.ActivityLogMockInput
	(*Data)(nil),                 // 2: generation.Data
	(*Segments)(nil),             // 3: generation.Segments
	(*Segment)(nil),              // 4: generation.Segment
	(*Clients)(nil),              // 5: generation.Clients
	(*Client)(nil),               // 6: generation.Client
}
var file_vault_activity_generation_generate_data_proto_depIdxs = []int32{
	0, // 0: generation.ActivityLogMockInput.write:type_name -> generation.WriteOptions
	2, // 1: generation.ActivityLogMockInput.data:type_name -> generation.Data
	5, // 2: generation.Data.all:type_name -> generation.Clients
	3, // 3: generation.Data.segments:type_name -> generation.Segments
	4, // 4: generation.Segments.segments:type_name -> generation.Segment
	5, // 5: generation.Segment.clients:type_name -> generation.Clients
	6, // 6: generation.Clients.clients:type_name -> generation.Client
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_vault_activity_generation_generate_data_proto_init() }
func file_vault_activity_generation_generate_data_proto_init() {
	if File_vault_activity_generation_generate_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vault_activity_generation_generate_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityLogMockInput); i {
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
		file_vault_activity_generation_generate_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_vault_activity_generation_generate_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Segments); i {
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
		file_vault_activity_generation_generate_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Segment); i {
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
		file_vault_activity_generation_generate_data_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Clients); i {
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
		file_vault_activity_generation_generate_data_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Client); i {
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
	file_vault_activity_generation_generate_data_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Data_CurrentMonth)(nil),
		(*Data_MonthsAgo)(nil),
		(*Data_All)(nil),
		(*Data_Segments)(nil),
	}
	file_vault_activity_generation_generate_data_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vault_activity_generation_generate_data_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vault_activity_generation_generate_data_proto_goTypes,
		DependencyIndexes: file_vault_activity_generation_generate_data_proto_depIdxs,
		EnumInfos:         file_vault_activity_generation_generate_data_proto_enumTypes,
		MessageInfos:      file_vault_activity_generation_generate_data_proto_msgTypes,
	}.Build()
	File_vault_activity_generation_generate_data_proto = out.File
	file_vault_activity_generation_generate_data_proto_rawDesc = nil
	file_vault_activity_generation_generate_data_proto_goTypes = nil
	file_vault_activity_generation_generate_data_proto_depIdxs = nil
}
