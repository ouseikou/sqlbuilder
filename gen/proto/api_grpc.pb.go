// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: proto/api.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SqlBuilderApi_Generate_FullMethodName = "/proto.SqlBuilderApi/Generate"
)

// SqlBuilderApiClient is the client API for SqlBuilderApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SqlBuilderApiClient interface {
	Generate(ctx context.Context, in *BuilderRequest, opts ...grpc.CallOption) (*CommonResponse, error)
}

type sqlBuilderApiClient struct {
	cc grpc.ClientConnInterface
}

func NewSqlBuilderApiClient(cc grpc.ClientConnInterface) SqlBuilderApiClient {
	return &sqlBuilderApiClient{cc}
}

func (c *sqlBuilderApiClient) Generate(ctx context.Context, in *BuilderRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, SqlBuilderApi_Generate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SqlBuilderApiServer is the server API for SqlBuilderApi service.
// All implementations must embed UnimplementedSqlBuilderApiServer
// for forward compatibility.
type SqlBuilderApiServer interface {
	Generate(context.Context, *BuilderRequest) (*CommonResponse, error)
	mustEmbedUnimplementedSqlBuilderApiServer()
}

// UnimplementedSqlBuilderApiServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSqlBuilderApiServer struct{}

func (UnimplementedSqlBuilderApiServer) Generate(context.Context, *BuilderRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}
func (UnimplementedSqlBuilderApiServer) mustEmbedUnimplementedSqlBuilderApiServer() {}
func (UnimplementedSqlBuilderApiServer) testEmbeddedByValue()                       {}

// UnsafeSqlBuilderApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SqlBuilderApiServer will
// result in compilation errors.
type UnsafeSqlBuilderApiServer interface {
	mustEmbedUnimplementedSqlBuilderApiServer()
}

func RegisterSqlBuilderApiServer(s grpc.ServiceRegistrar, srv SqlBuilderApiServer) {
	// If the following call pancis, it indicates UnimplementedSqlBuilderApiServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SqlBuilderApi_ServiceDesc, srv)
}

func _SqlBuilderApi_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuilderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SqlBuilderApiServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SqlBuilderApi_Generate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SqlBuilderApiServer).Generate(ctx, req.(*BuilderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SqlBuilderApi_ServiceDesc is the grpc.ServiceDesc for SqlBuilderApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SqlBuilderApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SqlBuilderApi",
	HandlerType: (*SqlBuilderApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _SqlBuilderApi_Generate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api.proto",
}
