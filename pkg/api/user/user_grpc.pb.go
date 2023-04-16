// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/user/user.proto

package user

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
	UserAPI_Login_FullMethodName    = "/user.UserAPI/Login"
	UserAPI_Register_FullMethodName = "/user.UserAPI/Register"
	UserAPI_GetUser_FullMethodName  = "/user.UserAPI/GetUser"
	UserAPI_Search_FullMethodName   = "/user.UserAPI/Search"
	UserAPI_AddUsers_FullMethodName = "/user.UserAPI/AddUsers"
)

// UserAPIClient is the client API for UserAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserAPIClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	AddUsers(ctx context.Context, in *AddUsersRequest, opts ...grpc.CallOption) (*AddUsersResponse, error)
}

type userAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewUserAPIClient(cc grpc.ClientConnInterface) UserAPIClient {
	return &userAPIClient{cc}
}

func (c *userAPIClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, UserAPI_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, UserAPI_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, UserAPI_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, UserAPI_Search_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIClient) AddUsers(ctx context.Context, in *AddUsersRequest, opts ...grpc.CallOption) (*AddUsersResponse, error) {
	out := new(AddUsersResponse)
	err := c.cc.Invoke(ctx, UserAPI_AddUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserAPIServer is the server API for UserAPI service.
// All implementations must embed UnimplementedUserAPIServer
// for forward compatibility
type UserAPIServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	AddUsers(context.Context, *AddUsersRequest) (*AddUsersResponse, error)
	mustEmbedUnimplementedUserAPIServer()
}

// UnimplementedUserAPIServer must be embedded to have forward compatible implementations.
type UnimplementedUserAPIServer struct {
}

func (UnimplementedUserAPIServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserAPIServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserAPIServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserAPIServer) Search(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedUserAPIServer) AddUsers(context.Context, *AddUsersRequest) (*AddUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUsers not implemented")
}
func (UnimplementedUserAPIServer) mustEmbedUnimplementedUserAPIServer() {}

// UnsafeUserAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserAPIServer will
// result in compilation errors.
type UnsafeUserAPIServer interface {
	mustEmbedUnimplementedUserAPIServer()
}

func RegisterUserAPIServer(s grpc.ServiceRegistrar, srv UserAPIServer) {
	s.RegisterService(&UserAPI_ServiceDesc, srv)
}

func _UserAPI_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserAPI_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAPI_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserAPI_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAPI_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserAPI_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAPI_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserAPI_Search_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAPI_AddUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).AddUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserAPI_AddUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).AddUsers(ctx, req.(*AddUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserAPI_ServiceDesc is the grpc.ServiceDesc for UserAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserAPI",
	HandlerType: (*UserAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _UserAPI_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _UserAPI_Register_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserAPI_GetUser_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _UserAPI_Search_Handler,
		},
		{
			MethodName: "AddUsers",
			Handler:    _UserAPI_AddUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/user/user.proto",
}