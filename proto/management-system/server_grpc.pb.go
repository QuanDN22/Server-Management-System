// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: management-system/server.proto

package management_system

import (
	context "context"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ManagementSystem_Ping_FullMethodName         = "/management.system.service.ManagementSystem/Ping"
	ManagementSystem_CreateServer_FullMethodName = "/management.system.service.ManagementSystem/CreateServer"
	ManagementSystem_UpdateServer_FullMethodName = "/management.system.service.ManagementSystem/UpdateServer"
	ManagementSystem_DeleteServer_FullMethodName = "/management.system.service.ManagementSystem/DeleteServer"
	ManagementSystem_ImportServer_FullMethodName = "/management.system.service.ManagementSystem/ImportServer"
)

// ManagementSystemClient is the client API for ManagementSystem service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagementSystemClient interface {
	// Ping server
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PingResponse, error)
	// Create server
	CreateServer(ctx context.Context, in *CreateServerRequest, opts ...grpc.CallOption) (*Server, error)
	// Update server
	UpdateServer(ctx context.Context, in *UpdateServerRequest, opts ...grpc.CallOption) (*Server, error)
	// Delete server
	DeleteServer(ctx context.Context, in *DeleteServerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Import server
	ImportServer(ctx context.Context, opts ...grpc.CallOption) (ManagementSystem_ImportServerClient, error)
}

type managementSystemClient struct {
	cc grpc.ClientConnInterface
}

func NewManagementSystemClient(cc grpc.ClientConnInterface) ManagementSystemClient {
	return &managementSystemClient{cc}
}

func (c *managementSystemClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, ManagementSystem_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementSystemClient) CreateServer(ctx context.Context, in *CreateServerRequest, opts ...grpc.CallOption) (*Server, error) {
	out := new(Server)
	err := c.cc.Invoke(ctx, ManagementSystem_CreateServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementSystemClient) UpdateServer(ctx context.Context, in *UpdateServerRequest, opts ...grpc.CallOption) (*Server, error) {
	out := new(Server)
	err := c.cc.Invoke(ctx, ManagementSystem_UpdateServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementSystemClient) DeleteServer(ctx context.Context, in *DeleteServerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ManagementSystem_DeleteServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementSystemClient) ImportServer(ctx context.Context, opts ...grpc.CallOption) (ManagementSystem_ImportServerClient, error) {
	stream, err := c.cc.NewStream(ctx, &ManagementSystem_ServiceDesc.Streams[0], ManagementSystem_ImportServer_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &managementSystemImportServerClient{stream}
	return x, nil
}

type ManagementSystem_ImportServerClient interface {
	Send(*ImportServerRequest) error
	CloseAndRecv() (*httpbody.HttpBody, error)
	grpc.ClientStream
}

type managementSystemImportServerClient struct {
	grpc.ClientStream
}

func (x *managementSystemImportServerClient) Send(m *ImportServerRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *managementSystemImportServerClient) CloseAndRecv() (*httpbody.HttpBody, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(httpbody.HttpBody)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ManagementSystemServer is the server API for ManagementSystem service.
// All implementations must embed UnimplementedManagementSystemServer
// for forward compatibility
type ManagementSystemServer interface {
	// Ping server
	Ping(context.Context, *emptypb.Empty) (*PingResponse, error)
	// Create server
	CreateServer(context.Context, *CreateServerRequest) (*Server, error)
	// Update server
	UpdateServer(context.Context, *UpdateServerRequest) (*Server, error)
	// Delete server
	DeleteServer(context.Context, *DeleteServerRequest) (*emptypb.Empty, error)
	// Import server
	ImportServer(ManagementSystem_ImportServerServer) error
	mustEmbedUnimplementedManagementSystemServer()
}

// UnimplementedManagementSystemServer must be embedded to have forward compatible implementations.
type UnimplementedManagementSystemServer struct {
}

func (UnimplementedManagementSystemServer) Ping(context.Context, *emptypb.Empty) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedManagementSystemServer) CreateServer(context.Context, *CreateServerRequest) (*Server, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateServer not implemented")
}
func (UnimplementedManagementSystemServer) UpdateServer(context.Context, *UpdateServerRequest) (*Server, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateServer not implemented")
}
func (UnimplementedManagementSystemServer) DeleteServer(context.Context, *DeleteServerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteServer not implemented")
}
func (UnimplementedManagementSystemServer) ImportServer(ManagementSystem_ImportServerServer) error {
	return status.Errorf(codes.Unimplemented, "method ImportServer not implemented")
}
func (UnimplementedManagementSystemServer) mustEmbedUnimplementedManagementSystemServer() {}

// UnsafeManagementSystemServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagementSystemServer will
// result in compilation errors.
type UnsafeManagementSystemServer interface {
	mustEmbedUnimplementedManagementSystemServer()
}

func RegisterManagementSystemServer(s grpc.ServiceRegistrar, srv ManagementSystemServer) {
	s.RegisterService(&ManagementSystem_ServiceDesc, srv)
}

func _ManagementSystem_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementSystemServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagementSystem_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementSystemServer).Ping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementSystem_CreateServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementSystemServer).CreateServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagementSystem_CreateServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementSystemServer).CreateServer(ctx, req.(*CreateServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementSystem_UpdateServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementSystemServer).UpdateServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagementSystem_UpdateServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementSystemServer).UpdateServer(ctx, req.(*UpdateServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementSystem_DeleteServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementSystemServer).DeleteServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagementSystem_DeleteServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementSystemServer).DeleteServer(ctx, req.(*DeleteServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementSystem_ImportServer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ManagementSystemServer).ImportServer(&managementSystemImportServerServer{stream})
}

type ManagementSystem_ImportServerServer interface {
	SendAndClose(*httpbody.HttpBody) error
	Recv() (*ImportServerRequest, error)
	grpc.ServerStream
}

type managementSystemImportServerServer struct {
	grpc.ServerStream
}

func (x *managementSystemImportServerServer) SendAndClose(m *httpbody.HttpBody) error {
	return x.ServerStream.SendMsg(m)
}

func (x *managementSystemImportServerServer) Recv() (*ImportServerRequest, error) {
	m := new(ImportServerRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ManagementSystem_ServiceDesc is the grpc.ServiceDesc for ManagementSystem service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ManagementSystem_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "management.system.service.ManagementSystem",
	HandlerType: (*ManagementSystemServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _ManagementSystem_Ping_Handler,
		},
		{
			MethodName: "CreateServer",
			Handler:    _ManagementSystem_CreateServer_Handler,
		},
		{
			MethodName: "UpdateServer",
			Handler:    _ManagementSystem_UpdateServer_Handler,
		},
		{
			MethodName: "DeleteServer",
			Handler:    _ManagementSystem_DeleteServer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ImportServer",
			Handler:       _ManagementSystem_ImportServer_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "management-system/server.proto",
}
