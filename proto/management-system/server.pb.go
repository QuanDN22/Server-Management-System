// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: management-system/server.proto

package management_system

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ping response
type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_management_system_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_management_system_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_management_system_server_proto_rawDescGZIP(), []int{0}
}

func (x *PingResponse) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

// delete server
type DeleteServerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server_ID int64 `protobuf:"varint,1,opt,name=Server_ID,json=server_id,proto3" json:"Server_ID,omitempty"`
}

func (x *DeleteServerRequest) Reset() {
	*x = DeleteServerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_management_system_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteServerRequest) ProtoMessage() {}

func (x *DeleteServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_management_system_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteServerRequest.ProtoReflect.Descriptor instead.
func (*DeleteServerRequest) Descriptor() ([]byte, []int) {
	return file_management_system_server_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteServerRequest) GetServer_ID() int64 {
	if x != nil {
		return x.Server_ID
	}
	return 0
}

// update server request
type UpdateServerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server_ID     int64  `protobuf:"varint,1,opt,name=Server_ID,json=server_id,proto3" json:"Server_ID,omitempty"`
	Server_Name   string `protobuf:"bytes,2,opt,name=Server_Name,json=server_name,proto3" json:"Server_Name,omitempty"`
	Server_IPv4   string `protobuf:"bytes,3,opt,name=Server_IPv4,json=server_ipv4,proto3" json:"Server_IPv4,omitempty"`
	Server_Status bool   `protobuf:"varint,4,opt,name=Server_Status,json=server_status,proto3" json:"Server_Status,omitempty"`
}

func (x *UpdateServerRequest) Reset() {
	*x = UpdateServerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_management_system_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateServerRequest) ProtoMessage() {}

func (x *UpdateServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_management_system_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateServerRequest.ProtoReflect.Descriptor instead.
func (*UpdateServerRequest) Descriptor() ([]byte, []int) {
	return file_management_system_server_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateServerRequest) GetServer_ID() int64 {
	if x != nil {
		return x.Server_ID
	}
	return 0
}

func (x *UpdateServerRequest) GetServer_Name() string {
	if x != nil {
		return x.Server_Name
	}
	return ""
}

func (x *UpdateServerRequest) GetServer_IPv4() string {
	if x != nil {
		return x.Server_IPv4
	}
	return ""
}

func (x *UpdateServerRequest) GetServer_Status() bool {
	if x != nil {
		return x.Server_Status
	}
	return false
}

// create server request
type CreateServerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server_Name   string `protobuf:"bytes,1,opt,name=Server_Name,json=server_name,proto3" json:"Server_Name,omitempty"`
	Server_IPv4   string `protobuf:"bytes,2,opt,name=Server_IPv4,json=server_ipv4,proto3" json:"Server_IPv4,omitempty"`
	Server_Status bool   `protobuf:"varint,3,opt,name=Server_Status,json=server_status,proto3" json:"Server_Status,omitempty"`
}

func (x *CreateServerRequest) Reset() {
	*x = CreateServerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_management_system_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateServerRequest) ProtoMessage() {}

func (x *CreateServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_management_system_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateServerRequest.ProtoReflect.Descriptor instead.
func (*CreateServerRequest) Descriptor() ([]byte, []int) {
	return file_management_system_server_proto_rawDescGZIP(), []int{3}
}

func (x *CreateServerRequest) GetServer_Name() string {
	if x != nil {
		return x.Server_Name
	}
	return ""
}

func (x *CreateServerRequest) GetServer_IPv4() string {
	if x != nil {
		return x.Server_IPv4
	}
	return ""
}

func (x *CreateServerRequest) GetServer_Status() bool {
	if x != nil {
		return x.Server_Status
	}
	return false
}

// server
type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server_ID     int64                  `protobuf:"varint,1,opt,name=Server_ID,json=server_id,proto3" json:"Server_ID,omitempty"`
	Server_Name   string                 `protobuf:"bytes,2,opt,name=Server_Name,json=server_name,proto3" json:"Server_Name,omitempty"`
	Server_IPv4   string                 `protobuf:"bytes,3,opt,name=Server_IPv4,json=server_ipv4,proto3" json:"Server_IPv4,omitempty"`
	Server_Status bool                   `protobuf:"varint,4,opt,name=Server_Status,json=server_status,proto3" json:"Server_Status,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=CreatedAt,json=created_at,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=UpdatedAt,json=updated_at,proto3" json:"UpdatedAt,omitempty"`
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_management_system_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_management_system_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_management_system_server_proto_rawDescGZIP(), []int{4}
}

func (x *Server) GetServer_ID() int64 {
	if x != nil {
		return x.Server_ID
	}
	return 0
}

func (x *Server) GetServer_Name() string {
	if x != nil {
		return x.Server_Name
	}
	return ""
}

func (x *Server) GetServer_IPv4() string {
	if x != nil {
		return x.Server_IPv4
	}
	return ""
}

func (x *Server) GetServer_Status() bool {
	if x != nil {
		return x.Server_Status
	}
	return false
}

func (x *Server) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Server) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_management_system_server_proto protoreflect.FileDescriptor

var file_management_system_server_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x19, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x22, 0x33, 0x0a, 0x13, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x22, 0x9d, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x49, 0x50, 0x76, 0x34, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x7f, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x49, 0x50, 0x76, 0x34, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x12, 0x24, 0x0a, 0x0d, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x86, 0x02, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x49, 0x50, 0x76, 0x34, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x12, 0x24,
	0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x12,
	0x39, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x32, 0xd4, 0x05, 0x0a, 0x10, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12,
	0x90, 0x01, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x27, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x47, 0x92, 0x41, 0x28, 0x0a, 0x06,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x20, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x1a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x69,
	0x6e, 0x67, 0x12, 0xb4, 0x01, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x12, 0x2e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x51, 0x92, 0x41, 0x34, 0x0a, 0x06, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x12, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x1a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0xc0, 0x01, 0x0a, 0x0c, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x2e, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x5d, 0x92,
	0x41, 0x34, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x0d, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x1a,
	0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73,
	0x2f, 0x7b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x49, 0x44, 0x7d, 0x12, 0xb2, 0x01, 0x0a,
	0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x2e, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x5a, 0x92, 0x41, 0x34, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x12, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x1a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1d, 0x2a, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x49, 0x44,
	0x7d, 0x42, 0x8a, 0x03, 0x92, 0x41, 0xc0, 0x02, 0x12, 0x83, 0x01, 0x0a, 0x23, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x20, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x20, 0x41,
	0x50, 0x49, 0x20, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x57, 0x0a, 0x09, 0x44, 0x69, 0x6e, 0x68, 0x20, 0x51, 0x75, 0x61, 0x6e, 0x12, 0x2c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x51, 0x75, 0x61, 0x6e, 0x44, 0x4e,
	0x32, 0x32, 0x2f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x1a, 0x1c, 0x64, 0x69, 0x6e,
	0x68, 0x6e, 0x67, 0x6f, 0x63, 0x71, 0x75, 0x61, 0x6e, 0x31, 0x31, 0x32, 0x33, 0x37, 0x38, 0x40,
	0x67, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x32, 0x03, 0x32, 0x2e, 0x30, 0x1a, 0x0e,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x3a, 0x38, 0x30, 0x38, 0x30, 0x2a, 0x02,
	0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x23, 0x0a, 0x21, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x4b,
	0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x13, 0x08, 0x02, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x10, 0x0a, 0x0e, 0x0a,
	0x0a, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x72, 0x49, 0x0a,
	0x17, 0x4d, 0x6f, 0x72, 0x65, 0x20, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x20, 0x67, 0x52, 0x50, 0x43,
	0x2d, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2d, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x51, 0x75, 0x61, 0x6e, 0x44, 0x4e, 0x32, 0x32, 0x2f, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2d, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2d,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_management_system_server_proto_rawDescOnce sync.Once
	file_management_system_server_proto_rawDescData = file_management_system_server_proto_rawDesc
)

func file_management_system_server_proto_rawDescGZIP() []byte {
	file_management_system_server_proto_rawDescOnce.Do(func() {
		file_management_system_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_management_system_server_proto_rawDescData)
	})
	return file_management_system_server_proto_rawDescData
}

var file_management_system_server_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_management_system_server_proto_goTypes = []interface{}{
	(*PingResponse)(nil),          // 0: management.system.service.PingResponse
	(*DeleteServerRequest)(nil),   // 1: management.system.service.DeleteServerRequest
	(*UpdateServerRequest)(nil),   // 2: management.system.service.UpdateServerRequest
	(*CreateServerRequest)(nil),   // 3: management.system.service.CreateServerRequest
	(*Server)(nil),                // 4: management.system.service.Server
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 6: google.protobuf.Empty
}
var file_management_system_server_proto_depIdxs = []int32{
	5, // 0: management.system.service.Server.CreatedAt:type_name -> google.protobuf.Timestamp
	5, // 1: management.system.service.Server.UpdatedAt:type_name -> google.protobuf.Timestamp
	6, // 2: management.system.service.ManagementSystem.Ping:input_type -> google.protobuf.Empty
	3, // 3: management.system.service.ManagementSystem.CreateServer:input_type -> management.system.service.CreateServerRequest
	2, // 4: management.system.service.ManagementSystem.UpdateServer:input_type -> management.system.service.UpdateServerRequest
	1, // 5: management.system.service.ManagementSystem.DeleteServer:input_type -> management.system.service.DeleteServerRequest
	0, // 6: management.system.service.ManagementSystem.Ping:output_type -> management.system.service.PingResponse
	4, // 7: management.system.service.ManagementSystem.CreateServer:output_type -> management.system.service.Server
	4, // 8: management.system.service.ManagementSystem.UpdateServer:output_type -> management.system.service.Server
	6, // 9: management.system.service.ManagementSystem.DeleteServer:output_type -> google.protobuf.Empty
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_management_system_server_proto_init() }
func file_management_system_server_proto_init() {
	if File_management_system_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_management_system_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
		file_management_system_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteServerRequest); i {
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
		file_management_system_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateServerRequest); i {
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
		file_management_system_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateServerRequest); i {
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
		file_management_system_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
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
			RawDescriptor: file_management_system_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_management_system_server_proto_goTypes,
		DependencyIndexes: file_management_system_server_proto_depIdxs,
		MessageInfos:      file_management_system_server_proto_msgTypes,
	}.Build()
	File_management_system_server_proto = out.File
	file_management_system_server_proto_rawDesc = nil
	file_management_system_server_proto_goTypes = nil
	file_management_system_server_proto_depIdxs = nil
}
