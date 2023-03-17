// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: user.proto

package __

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

// UserApiClient is the client API for UserApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserApiClient interface {
	// rpc Heartbeat(BeatsPerMinuteRequest) returns (stream BeatsPerMinuteResponse){};
	Echo(ctx context.Context, in *ResponseRequest, opts ...grpc.CallOption) (*ResponseRequest, error)
}

type userApiClient struct {
	cc grpc.ClientConnInterface
}

func NewUserApiClient(cc grpc.ClientConnInterface) UserApiClient {
	return &userApiClient{cc}
}

func (c *userApiClient) Echo(ctx context.Context, in *ResponseRequest, opts ...grpc.CallOption) (*ResponseRequest, error) {
	out := new(ResponseRequest)
	err := c.cc.Invoke(ctx, "/main.UserApi/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserApiServer is the server API for UserApi service.
// All implementations must embed UnimplementedUserApiServer
// for forward compatibility
type UserApiServer interface {
	// rpc Heartbeat(BeatsPerMinuteRequest) returns (stream BeatsPerMinuteResponse){};
	Echo(context.Context, *ResponseRequest) (*ResponseRequest, error)
	mustEmbedUnimplementedUserApiServer()
}

// UnimplementedUserApiServer must be embedded to have forward compatible implementations.
type UnimplementedUserApiServer struct {
}

func (UnimplementedUserApiServer) Echo(context.Context, *ResponseRequest) (*ResponseRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedUserApiServer) mustEmbedUnimplementedUserApiServer() {}

// UnsafeUserApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserApiServer will
// result in compilation errors.
type UnsafeUserApiServer interface {
	mustEmbedUnimplementedUserApiServer()
}

func RegisterUserApiServer(s grpc.ServiceRegistrar, srv UserApiServer) {
	s.RegisterService(&UserApi_ServiceDesc, srv)
}

func _UserApi_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResponseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserApiServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.UserApi/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserApiServer).Echo(ctx, req.(*ResponseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserApi_ServiceDesc is the grpc.ServiceDesc for UserApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.UserApi",
	HandlerType: (*UserApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _UserApi_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
