// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: auth.proto

package auth

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
	AuthSrv_Register_FullMethodName = "/AuthSrv/register"
	AuthSrv_Login_FullMethodName    = "/AuthSrv/login"
	AuthSrv_Refresh_FullMethodName  = "/AuthSrv/refresh"
)

// AuthSrvClient is the client API for AuthSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthSrvClient interface {
	Register(ctx context.Context, in *RegisterIn, opts ...grpc.CallOption) (*RegisterOut, error)
	Login(ctx context.Context, in *LoginIn, opts ...grpc.CallOption) (*LoginOut, error)
	Refresh(ctx context.Context, in *RefreshIn, opts ...grpc.CallOption) (*RefreshOut, error)
}

type authSrvClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthSrvClient(cc grpc.ClientConnInterface) AuthSrvClient {
	return &authSrvClient{cc}
}

func (c *authSrvClient) Register(ctx context.Context, in *RegisterIn, opts ...grpc.CallOption) (*RegisterOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterOut)
	err := c.cc.Invoke(ctx, AuthSrv_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authSrvClient) Login(ctx context.Context, in *LoginIn, opts ...grpc.CallOption) (*LoginOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginOut)
	err := c.cc.Invoke(ctx, AuthSrv_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authSrvClient) Refresh(ctx context.Context, in *RefreshIn, opts ...grpc.CallOption) (*RefreshOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RefreshOut)
	err := c.cc.Invoke(ctx, AuthSrv_Refresh_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthSrvServer is the server API for AuthSrv service.
// All implementations must embed UnimplementedAuthSrvServer
// for forward compatibility.
type AuthSrvServer interface {
	Register(context.Context, *RegisterIn) (*RegisterOut, error)
	Login(context.Context, *LoginIn) (*LoginOut, error)
	Refresh(context.Context, *RefreshIn) (*RefreshOut, error)
	mustEmbedUnimplementedAuthSrvServer()
}

// UnimplementedAuthSrvServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthSrvServer struct{}

func (UnimplementedAuthSrvServer) Register(context.Context, *RegisterIn) (*RegisterOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthSrvServer) Login(context.Context, *LoginIn) (*LoginOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthSrvServer) Refresh(context.Context, *RefreshIn) (*RefreshOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (UnimplementedAuthSrvServer) mustEmbedUnimplementedAuthSrvServer() {}
func (UnimplementedAuthSrvServer) testEmbeddedByValue()                 {}

// UnsafeAuthSrvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthSrvServer will
// result in compilation errors.
type UnsafeAuthSrvServer interface {
	mustEmbedUnimplementedAuthSrvServer()
}

func RegisterAuthSrvServer(s grpc.ServiceRegistrar, srv AuthSrvServer) {
	// If the following call pancis, it indicates UnimplementedAuthSrvServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AuthSrv_ServiceDesc, srv)
}

func _AuthSrv_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthSrvServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthSrv_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthSrvServer).Register(ctx, req.(*RegisterIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthSrv_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthSrvServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthSrv_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthSrvServer).Login(ctx, req.(*LoginIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthSrv_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthSrvServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthSrv_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthSrvServer).Refresh(ctx, req.(*RefreshIn))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthSrv_ServiceDesc is the grpc.ServiceDesc for AuthSrv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthSrv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AuthSrv",
	HandlerType: (*AuthSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "register",
			Handler:    _AuthSrv_Register_Handler,
		},
		{
			MethodName: "login",
			Handler:    _AuthSrv_Login_Handler,
		},
		{
			MethodName: "refresh",
			Handler:    _AuthSrv_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
