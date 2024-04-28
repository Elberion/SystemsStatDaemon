// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: system_stats.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SystemStat_Internal_FullMethodName = "/system_stats.SystemStat/Internal"
)

// SystemStatClient is the client API for SystemStat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SystemStatClient interface {
	Internal(ctx context.Context, in *CollectSettings, opts ...grpc.CallOption) (SystemStat_InternalClient, error)
}

type systemStatClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemStatClient(cc grpc.ClientConnInterface) SystemStatClient {
	return &systemStatClient{cc}
}

func (c *systemStatClient) Internal(ctx context.Context, in *CollectSettings, opts ...grpc.CallOption) (SystemStat_InternalClient, error) {
	stream, err := c.cc.NewStream(ctx, &SystemStat_ServiceDesc.Streams[0], SystemStat_Internal_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &systemStatInternalClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SystemStat_InternalClient interface {
	Recv() (*SystemStats, error)
	grpc.ClientStream
}

type systemStatInternalClient struct {
	grpc.ClientStream
}

func (x *systemStatInternalClient) Recv() (*SystemStats, error) {
	m := new(SystemStats)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SystemStatServer is the server API for SystemStat service.
// All implementations must embed UnimplementedSystemStatServer
// for forward compatibility
type SystemStatServer interface {
	Internal(*CollectSettings, SystemStat_InternalServer) error
	mustEmbedUnimplementedSystemStatServer()
}

// UnimplementedSystemStatServer must be embedded to have forward compatible implementations.
type UnimplementedSystemStatServer struct {
}

func (UnimplementedSystemStatServer) Internal(*CollectSettings, SystemStat_InternalServer) error {
	return status.Errorf(codes.Unimplemented, "method Internal not implemented")
}
func (UnimplementedSystemStatServer) mustEmbedUnimplementedSystemStatServer() {}

// UnsafeSystemStatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemStatServer will
// result in compilation errors.
type UnsafeSystemStatServer interface {
	mustEmbedUnimplementedSystemStatServer()
}

func RegisterSystemStatServer(s grpc.ServiceRegistrar, srv SystemStatServer) {
	s.RegisterService(&SystemStat_ServiceDesc, srv)
}

func _SystemStat_Internal_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CollectSettings)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SystemStatServer).Internal(m, &systemStatInternalServer{stream})
}

type SystemStat_InternalServer interface {
	Send(*SystemStats) error
	grpc.ServerStream
}

type systemStatInternalServer struct {
	grpc.ServerStream
}

func (x *systemStatInternalServer) Send(m *SystemStats) error {
	return x.ServerStream.SendMsg(m)
}

// SystemStat_ServiceDesc is the grpc.ServiceDesc for SystemStat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SystemStat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "system_stats.SystemStat",
	HandlerType: (*SystemStatServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Internal",
			Handler:       _SystemStat_Internal_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "system_stats.proto",
}
