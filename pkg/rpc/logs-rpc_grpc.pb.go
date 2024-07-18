// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v4.25.1
// source: logs-rpc.proto

package rpc

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
	CacheService_GetCache_FullMethodName = "/pb.CacheService/GetCache"
)

// CacheServiceClient is the client API for CacheService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CacheServiceClient interface {
	GetCache(ctx context.Context, in *CacheRequest, opts ...grpc.CallOption) (*CacheResponse, error)
}

type cacheServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCacheServiceClient(cc grpc.ClientConnInterface) CacheServiceClient {
	return &cacheServiceClient{cc}
}

func (c *cacheServiceClient) GetCache(ctx context.Context, in *CacheRequest, opts ...grpc.CallOption) (*CacheResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CacheResponse)
	err := c.cc.Invoke(ctx, CacheService_GetCache_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacheServiceServer is the server API for CacheService service.
// All implementations must embed UnimplementedCacheServiceServer
// for forward compatibility
type CacheServiceServer interface {
	GetCache(context.Context, *CacheRequest) (*CacheResponse, error)
	mustEmbedUnimplementedCacheServiceServer()
}

// UnimplementedCacheServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCacheServiceServer struct {
}

func (UnimplementedCacheServiceServer) GetCache(context.Context, *CacheRequest) (*CacheResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCache not implemented")
}
func (UnimplementedCacheServiceServer) mustEmbedUnimplementedCacheServiceServer() {}

// UnsafeCacheServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CacheServiceServer will
// result in compilation errors.
type UnsafeCacheServiceServer interface {
	mustEmbedUnimplementedCacheServiceServer()
}

func RegisterCacheServiceServer(s grpc.ServiceRegistrar, srv CacheServiceServer) {
	s.RegisterService(&CacheService_ServiceDesc, srv)
}

func _CacheService_GetCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CacheRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).GetCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheService_GetCache_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).GetCache(ctx, req.(*CacheRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CacheService_ServiceDesc is the grpc.ServiceDesc for CacheService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CacheService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CacheService",
	HandlerType: (*CacheServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCache",
			Handler:    _CacheService_GetCache_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "logs-rpc.proto",
}