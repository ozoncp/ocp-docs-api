// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: api/ocp-docs-api/ocp-docs-api.proto

package ocp_docs_api

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ListDocsV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  uint64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset uint64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListDocsV1Request) Reset() {
	*x = ListDocsV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_docs_api_ocp_docs_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDocsV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDocsV1Request) ProtoMessage() {}

func (x *ListDocsV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_docs_api_ocp_docs_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDocsV1Request.ProtoReflect.Descriptor instead.
func (*ListDocsV1Request) Descriptor() ([]byte, []int) {
	return file_api_ocp_docs_api_ocp_docs_api_proto_rawDescGZIP(), []int{0}
}

func (x *ListDocsV1Request) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListDocsV1Request) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListDocsV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Docs []*Doc `protobuf:"bytes,1,rep,name=docs,proto3" json:"docs,omitempty"`
}

func (x *ListDocsV1Response) Reset() {
	*x = ListDocsV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_docs_api_ocp_docs_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDocsV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDocsV1Response) ProtoMessage() {}

func (x *ListDocsV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_docs_api_ocp_docs_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDocsV1Response.ProtoReflect.Descriptor instead.
func (*ListDocsV1Response) Descriptor() ([]byte, []int) {
	return file_api_ocp_docs_api_ocp_docs_api_proto_rawDescGZIP(), []int{1}
}

func (x *ListDocsV1Response) GetDocs() []*Doc {
	if x != nil {
		return x.Docs
	}
	return nil
}

type DescribeDocV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DescribeDocV1Request) Reset() {
	*x = DescribeDocV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_docs_api_ocp_docs_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeDocV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeDocV1Request) ProtoMessage() {}

func (x *DescribeDocV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_docs_api_ocp_docs_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeDocV1Request.ProtoReflect.Descriptor instead.
func (*DescribeDocV1Request) Descriptor() ([]byte, []int) {
	return file_api_ocp_docs_api_ocp_docs_api_proto_rawDescGZIP(), []int{2}
}

func (x *DescribeDocV1Request) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DescribeDocV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Doc *Doc `protobuf:"bytes,1,opt,name=doc,proto3" json:"doc,omitempty"`
}

func (x *DescribeDocV1Response) Reset() {
	*x = DescribeDocV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_docs_api_ocp_docs_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeDocV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeDocV1Response) ProtoMessage() {}

func (x *DescribeDocV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_docs_api_ocp_docs_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeDocV1Response.ProtoReflect.Descriptor instead.
func (*DescribeDocV1Response) Descriptor() ([]byte, []int) {
	return file_api_ocp_docs_api_ocp_docs_api_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeDocV1Response) GetDoc() *Doc {
	if x != nil {
		return x.Doc
	}
	return nil
}

type CreateDocV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Link       string `protobuf:"bytes,2,opt,name=link,proto3" json:"link,omitempty"`
	SourceLink string `protobuf:"bytes,3,opt,name=sourceLink,proto3" json:"sourceLink,omitempty"`
}

func (x *CreateDocV1Request) Reset() {
	*x = CreateDocV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_docs_api_ocp_docs_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDocV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDocV1Request) ProtoMessage() {}

func (x *CreateDocV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_docs_api_ocp_docs_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDocV1Request.ProtoReflect.Descriptor instead.
func (*CreateDocV1Request) Descriptor() ([]byte, []int) {
	return file_api_ocp_docs_api_ocp_docs_api_proto_rawDescGZIP(), []int{4}
}

func (x *CreateDocV1Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateDocV1Request) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *CreateDocV1Request) GetSourceLink() string {
	if x != nil {
		return x.SourceLink
	}
	return ""
}

type CreateDocV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateDocV1Response) Reset() {
	*x = CreateDocV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_docs_api_ocp_docs_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDocV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDocV1Response) ProtoMessage() {}

func (x *CreateDocV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_docs_api_ocp_docs_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDocV1Response.ProtoReflect.Descriptor instead.
func (*CreateDocV1Response) Descriptor() ([]byte, []int) {
	return file_api_ocp_docs_api_ocp_docs_api_proto_rawDescGZIP(), []int{5}
}

func (x *CreateDocV1Response) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RemoveDocV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RemoveDocV1Request) Reset() {
	*x = RemoveDocV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_docs_api_ocp_docs_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveDocV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveDocV1Request) ProtoMessage() {}

func (x *RemoveDocV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_docs_api_ocp_docs_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveDocV1Request.ProtoReflect.Descriptor instead.
func (*RemoveDocV1Request) Descriptor() ([]byte, []int) {
	return file_api_ocp_docs_api_ocp_docs_api_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveDocV1Request) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RemoveDocV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Found bool `protobuf:"varint,1,opt,name=found,proto3" json:"found,omitempty"`
}

func (x *RemoveDocV1Response) Reset() {
	*x = RemoveDocV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_docs_api_ocp_docs_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveDocV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveDocV1Response) ProtoMessage() {}

func (x *RemoveDocV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_docs_api_ocp_docs_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveDocV1Response.ProtoReflect.Descriptor instead.
func (*RemoveDocV1Response) Descriptor() ([]byte, []int) {
	return file_api_ocp_docs_api_ocp_docs_api_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveDocV1Response) GetFound() bool {
	if x != nil {
		return x.Found
	}
	return false
}

type Doc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Link       string `protobuf:"bytes,3,opt,name=link,proto3" json:"link,omitempty"`
	SourceLink string `protobuf:"bytes,4,opt,name=sourceLink,proto3" json:"sourceLink,omitempty"`
}

func (x *Doc) Reset() {
	*x = Doc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_docs_api_ocp_docs_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Doc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Doc) ProtoMessage() {}

func (x *Doc) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_docs_api_ocp_docs_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Doc.ProtoReflect.Descriptor instead.
func (*Doc) Descriptor() ([]byte, []int) {
	return file_api_ocp_docs_api_ocp_docs_api_proto_rawDescGZIP(), []int{8}
}

func (x *Doc) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Doc) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Doc) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *Doc) GetSourceLink() string {
	if x != nil {
		return x.SourceLink
	}
	return ""
}

var File_api_ocp_docs_api_ocp_docs_api_proto protoreflect.FileDescriptor

var file_api_ocp_docs_api_ocp_docs_api_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x64, 0x6f, 0x63, 0x73, 0x2d, 0x61,
	0x70, 0x69, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x64, 0x6f, 0x63, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6f, 0x63, 0x70, 0x2e, 0x64, 0x6f, 0x63, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x41, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x6f, 0x63, 0x73, 0x56, 0x31, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x22, 0x3b, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x6f, 0x63, 0x73,
	0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x6f,
	0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x64,
	0x6f, 0x63, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x6f, 0x63, 0x52, 0x04, 0x64, 0x6f, 0x63,
	0x73, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x44, 0x6f, 0x63,
	0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3c, 0x0a, 0x15, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x44, 0x6f, 0x63, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x23, 0x0a, 0x03, 0x64, 0x6f, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x64, 0x6f, 0x63, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44,
	0x6f, 0x63, 0x52, 0x03, 0x64, 0x6f, 0x63, 0x22, 0x5c, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x6f, 0x63, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c,
	0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x25, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x6f, 0x63, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x12,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x6f, 0x63, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x2b, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x6f, 0x63, 0x56,
	0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x75,
	0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x22,
	0x5d, 0x0a, 0x03, 0x44, 0x6f, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69,
	0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x1e,
	0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x32, 0xb1,
	0x03, 0x0a, 0x0a, 0x4f, 0x63, 0x70, 0x44, 0x6f, 0x63, 0x73, 0x41, 0x70, 0x69, 0x12, 0x61, 0x0a,
	0x0a, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x6f, 0x63, 0x73, 0x56, 0x31, 0x12, 0x1f, 0x2e, 0x6f, 0x63,
	0x70, 0x2e, 0x64, 0x6f, 0x63, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44,
	0x6f, 0x63, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6f,
	0x63, 0x70, 0x2e, 0x64, 0x6f, 0x63, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x44, 0x6f, 0x63, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x63, 0x73,
	0x12, 0x6f, 0x0a, 0x0d, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x44, 0x6f, 0x63, 0x56,
	0x31, 0x12, 0x22, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x64, 0x6f, 0x63, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x44, 0x6f, 0x63, 0x56, 0x31, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x64, 0x6f, 0x63, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x44, 0x6f, 0x63,
	0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x63, 0x73, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x12, 0x64, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x56, 0x31,
	0x12, 0x20, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x64, 0x6f, 0x63, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x64, 0x6f, 0x63, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x56, 0x31, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x22, 0x08, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x6f, 0x63, 0x73, 0x12, 0x69, 0x0a, 0x0b, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x44, 0x6f, 0x63, 0x56, 0x31, 0x12, 0x20, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x64, 0x6f, 0x63,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x6f, 0x63, 0x56,
	0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x64,
	0x6f, 0x63, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x6f,
	0x63, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0f, 0x2a, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x63, 0x73, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6f, 0x7a, 0x6f, 0x6e, 0x63, 0x70, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x64, 0x6f, 0x63, 0x73,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x64, 0x6f, 0x63,
	0x73, 0x2d, 0x61, 0x70, 0x69, 0x3b, 0x6f, 0x63, 0x70, 0x5f, 0x64, 0x6f, 0x63, 0x73, 0x5f, 0x61,
	0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_ocp_docs_api_ocp_docs_api_proto_rawDescOnce sync.Once
	file_api_ocp_docs_api_ocp_docs_api_proto_rawDescData = file_api_ocp_docs_api_ocp_docs_api_proto_rawDesc
)

func file_api_ocp_docs_api_ocp_docs_api_proto_rawDescGZIP() []byte {
	file_api_ocp_docs_api_ocp_docs_api_proto_rawDescOnce.Do(func() {
		file_api_ocp_docs_api_ocp_docs_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_ocp_docs_api_ocp_docs_api_proto_rawDescData)
	})
	return file_api_ocp_docs_api_ocp_docs_api_proto_rawDescData
}

var file_api_ocp_docs_api_ocp_docs_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_ocp_docs_api_ocp_docs_api_proto_goTypes = []interface{}{
	(*ListDocsV1Request)(nil),     // 0: ocp.docs.api.ListDocsV1Request
	(*ListDocsV1Response)(nil),    // 1: ocp.docs.api.ListDocsV1Response
	(*DescribeDocV1Request)(nil),  // 2: ocp.docs.api.DescribeDocV1Request
	(*DescribeDocV1Response)(nil), // 3: ocp.docs.api.DescribeDocV1Response
	(*CreateDocV1Request)(nil),    // 4: ocp.docs.api.CreateDocV1Request
	(*CreateDocV1Response)(nil),   // 5: ocp.docs.api.CreateDocV1Response
	(*RemoveDocV1Request)(nil),    // 6: ocp.docs.api.RemoveDocV1Request
	(*RemoveDocV1Response)(nil),   // 7: ocp.docs.api.RemoveDocV1Response
	(*Doc)(nil),                   // 8: ocp.docs.api.Doc
}
var file_api_ocp_docs_api_ocp_docs_api_proto_depIdxs = []int32{
	8, // 0: ocp.docs.api.ListDocsV1Response.docs:type_name -> ocp.docs.api.Doc
	8, // 1: ocp.docs.api.DescribeDocV1Response.doc:type_name -> ocp.docs.api.Doc
	0, // 2: ocp.docs.api.OcpDocsApi.ListDocsV1:input_type -> ocp.docs.api.ListDocsV1Request
	2, // 3: ocp.docs.api.OcpDocsApi.DescribeDocV1:input_type -> ocp.docs.api.DescribeDocV1Request
	4, // 4: ocp.docs.api.OcpDocsApi.CreateDocV1:input_type -> ocp.docs.api.CreateDocV1Request
	6, // 5: ocp.docs.api.OcpDocsApi.RemoveDocV1:input_type -> ocp.docs.api.RemoveDocV1Request
	1, // 6: ocp.docs.api.OcpDocsApi.ListDocsV1:output_type -> ocp.docs.api.ListDocsV1Response
	3, // 7: ocp.docs.api.OcpDocsApi.DescribeDocV1:output_type -> ocp.docs.api.DescribeDocV1Response
	5, // 8: ocp.docs.api.OcpDocsApi.CreateDocV1:output_type -> ocp.docs.api.CreateDocV1Response
	7, // 9: ocp.docs.api.OcpDocsApi.RemoveDocV1:output_type -> ocp.docs.api.RemoveDocV1Response
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_ocp_docs_api_ocp_docs_api_proto_init() }
func file_api_ocp_docs_api_ocp_docs_api_proto_init() {
	if File_api_ocp_docs_api_ocp_docs_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_ocp_docs_api_ocp_docs_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDocsV1Request); i {
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
		file_api_ocp_docs_api_ocp_docs_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDocsV1Response); i {
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
		file_api_ocp_docs_api_ocp_docs_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeDocV1Request); i {
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
		file_api_ocp_docs_api_ocp_docs_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeDocV1Response); i {
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
		file_api_ocp_docs_api_ocp_docs_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDocV1Request); i {
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
		file_api_ocp_docs_api_ocp_docs_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDocV1Response); i {
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
		file_api_ocp_docs_api_ocp_docs_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveDocV1Request); i {
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
		file_api_ocp_docs_api_ocp_docs_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveDocV1Response); i {
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
		file_api_ocp_docs_api_ocp_docs_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Doc); i {
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
			RawDescriptor: file_api_ocp_docs_api_ocp_docs_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_ocp_docs_api_ocp_docs_api_proto_goTypes,
		DependencyIndexes: file_api_ocp_docs_api_ocp_docs_api_proto_depIdxs,
		MessageInfos:      file_api_ocp_docs_api_ocp_docs_api_proto_msgTypes,
	}.Build()
	File_api_ocp_docs_api_ocp_docs_api_proto = out.File
	file_api_ocp_docs_api_ocp_docs_api_proto_rawDesc = nil
	file_api_ocp_docs_api_ocp_docs_api_proto_goTypes = nil
	file_api_ocp_docs_api_ocp_docs_api_proto_depIdxs = nil
}