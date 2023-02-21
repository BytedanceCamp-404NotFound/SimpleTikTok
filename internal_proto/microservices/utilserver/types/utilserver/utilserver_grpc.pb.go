// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/utilserver.proto

package utilserver

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

// UtilserverClient is the client API for Utilserver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UtilserverClient interface {
	GetMinio(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*MinioResponse, error)
}

type utilserverClient struct {
	cc grpc.ClientConnInterface
}

func NewUtilserverClient(cc grpc.ClientConnInterface) UtilserverClient {
	return &utilserverClient{cc}
}

func (c *utilserverClient) GetMinio(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*MinioResponse, error) {
	out := new(MinioResponse)
	err := c.cc.Invoke(ctx, "/utilserver.Utilserver/getMinio", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UtilserverServer is the server API for Utilserver service.
// All implementations must embed UnimplementedUtilserverServer
// for forward compatibility
type UtilserverServer interface {
	GetMinio(context.Context, *IdRequest) (*MinioResponse, error)
	mustEmbedUnimplementedUtilserverServer()
}

// UnimplementedUtilserverServer must be embedded to have forward compatible implementations.
type UnimplementedUtilserverServer struct {
}

func (UnimplementedUtilserverServer) GetMinio(context.Context, *IdRequest) (*MinioResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMinio not implemented")
}
func (UnimplementedUtilserverServer) mustEmbedUnimplementedUtilserverServer() {}

// UnsafeUtilserverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UtilserverServer will
// result in compilation errors.
type UnsafeUtilserverServer interface {
	mustEmbedUnimplementedUtilserverServer()
}

func RegisterUtilserverServer(s grpc.ServiceRegistrar, srv UtilserverServer) {
	s.RegisterService(&Utilserver_ServiceDesc, srv)
}

func _Utilserver_GetMinio_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UtilserverServer).GetMinio(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/utilserver.Utilserver/getMinio",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UtilserverServer).GetMinio(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Utilserver_ServiceDesc is the grpc.ServiceDesc for Utilserver service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Utilserver_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "utilserver.Utilserver",
	HandlerType: (*UtilserverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getMinio",
			Handler:    _Utilserver_GetMinio_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/utilserver.proto",
}
