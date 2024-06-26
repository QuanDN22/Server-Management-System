// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: monitor/monitor.proto

package monitor

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Monitor_GetUpTime_FullMethodName = "/monitor.service.monitor/GetUpTime"
)

// MonitorClient is the client API for Monitor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MonitorClient interface {
	GetUpTime(ctx context.Context, in *UptimeRequest, opts ...grpc.CallOption) (*UptimeResponse, error)
}

type monitorClient struct {
	cc grpc.ClientConnInterface
}

func NewMonitorClient(cc grpc.ClientConnInterface) MonitorClient {
	return &monitorClient{cc}
}

func (c *monitorClient) GetUpTime(ctx context.Context, in *UptimeRequest, opts ...grpc.CallOption) (*UptimeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UptimeResponse)
	err := c.cc.Invoke(ctx, Monitor_GetUpTime_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MonitorServer is the server API for Monitor service.
// All implementations must embed UnimplementedMonitorServer
// for forward compatibility
type MonitorServer interface {
	GetUpTime(context.Context, *UptimeRequest) (*UptimeResponse, error)
	mustEmbedUnimplementedMonitorServer()
}

// UnimplementedMonitorServer must be embedded to have forward compatible implementations.
type UnimplementedMonitorServer struct {
}

func (UnimplementedMonitorServer) GetUpTime(context.Context, *UptimeRequest) (*UptimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUpTime not implemented")
}
func (UnimplementedMonitorServer) mustEmbedUnimplementedMonitorServer() {}

// UnsafeMonitorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MonitorServer will
// result in compilation errors.
type UnsafeMonitorServer interface {
	mustEmbedUnimplementedMonitorServer()
}

func RegisterMonitorServer(s grpc.ServiceRegistrar, srv MonitorServer) {
	s.RegisterService(&Monitor_ServiceDesc, srv)
}

func _Monitor_GetUpTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UptimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServer).GetUpTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Monitor_GetUpTime_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServer).GetUpTime(ctx, req.(*UptimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Monitor_ServiceDesc is the grpc.ServiceDesc for Monitor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Monitor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "monitor.service.monitor",
	HandlerType: (*MonitorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUpTime",
			Handler:    _Monitor_GetUpTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "monitor/monitor.proto",
}
