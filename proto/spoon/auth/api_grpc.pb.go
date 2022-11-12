// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: spoon/auth/api.proto

package auth

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

// AuthenticationSessionServiceClient is the client API for AuthenticationSessionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticationSessionServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error)
}

type authenticationSessionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationSessionServiceClient(cc grpc.ClientConnInterface) AuthenticationSessionServiceClient {
	return &authenticationSessionServiceClient{cc}
}

func (c *authenticationSessionServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthenticationSessionService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationSessionServiceClient) Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error) {
	out := new(RefreshResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthenticationSessionService/Refresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationSessionServiceServer is the server API for AuthenticationSessionService service.
// All implementations must embed UnimplementedAuthenticationSessionServiceServer
// for forward compatibility
type AuthenticationSessionServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error)
	mustEmbedUnimplementedAuthenticationSessionServiceServer()
}

// UnimplementedAuthenticationSessionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticationSessionServiceServer struct {
}

func (UnimplementedAuthenticationSessionServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthenticationSessionServiceServer) Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (UnimplementedAuthenticationSessionServiceServer) mustEmbedUnimplementedAuthenticationSessionServiceServer() {
}

// UnsafeAuthenticationSessionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticationSessionServiceServer will
// result in compilation errors.
type UnsafeAuthenticationSessionServiceServer interface {
	mustEmbedUnimplementedAuthenticationSessionServiceServer()
}

func RegisterAuthenticationSessionServiceServer(s grpc.ServiceRegistrar, srv AuthenticationSessionServiceServer) {
	s.RegisterService(&AuthenticationSessionService_ServiceDesc, srv)
}

func _AuthenticationSessionService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationSessionServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthenticationSessionService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationSessionServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationSessionService_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationSessionServiceServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthenticationSessionService/Refresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationSessionServiceServer).Refresh(ctx, req.(*RefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthenticationSessionService_ServiceDesc is the grpc.ServiceDesc for AuthenticationSessionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthenticationSessionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthenticationSessionService",
	HandlerType: (*AuthenticationSessionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthenticationSessionService_Login_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _AuthenticationSessionService_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spoon/auth/api.proto",
}

// AuthenticationManagementServiceClient is the client API for AuthenticationManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticationManagementServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
}

type authenticationManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationManagementServiceClient(cc grpc.ClientConnInterface) AuthenticationManagementServiceClient {
	return &authenticationManagementServiceClient{cc}
}

func (c *authenticationManagementServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthenticationManagementService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationManagementServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthenticationManagementService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationManagementServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthenticationManagementService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationManagementServiceServer is the server API for AuthenticationManagementService service.
// All implementations must embed UnimplementedAuthenticationManagementServiceServer
// for forward compatibility
type AuthenticationManagementServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	mustEmbedUnimplementedAuthenticationManagementServiceServer()
}

// UnimplementedAuthenticationManagementServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticationManagementServiceServer struct {
}

func (UnimplementedAuthenticationManagementServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedAuthenticationManagementServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedAuthenticationManagementServiceServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedAuthenticationManagementServiceServer) mustEmbedUnimplementedAuthenticationManagementServiceServer() {
}

// UnsafeAuthenticationManagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticationManagementServiceServer will
// result in compilation errors.
type UnsafeAuthenticationManagementServiceServer interface {
	mustEmbedUnimplementedAuthenticationManagementServiceServer()
}

func RegisterAuthenticationManagementServiceServer(s grpc.ServiceRegistrar, srv AuthenticationManagementServiceServer) {
	s.RegisterService(&AuthenticationManagementService_ServiceDesc, srv)
}

func _AuthenticationManagementService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationManagementServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthenticationManagementService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationManagementServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationManagementService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationManagementServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthenticationManagementService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationManagementServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationManagementService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationManagementServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthenticationManagementService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationManagementServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthenticationManagementService_ServiceDesc is the grpc.ServiceDesc for AuthenticationManagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthenticationManagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthenticationManagementService",
	HandlerType: (*AuthenticationManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _AuthenticationManagementService_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _AuthenticationManagementService_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _AuthenticationManagementService_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spoon/auth/api.proto",
}

// AuthenticationAuditServiceClient is the client API for AuthenticationAuditService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticationAuditServiceClient interface {
	GetPublicKey(ctx context.Context, in *GetPublicKeyRequest, opts ...grpc.CallOption) (*GetPublicKeyResponse, error)
}

type authenticationAuditServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationAuditServiceClient(cc grpc.ClientConnInterface) AuthenticationAuditServiceClient {
	return &authenticationAuditServiceClient{cc}
}

func (c *authenticationAuditServiceClient) GetPublicKey(ctx context.Context, in *GetPublicKeyRequest, opts ...grpc.CallOption) (*GetPublicKeyResponse, error) {
	out := new(GetPublicKeyResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthenticationAuditService/GetPublicKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationAuditServiceServer is the server API for AuthenticationAuditService service.
// All implementations must embed UnimplementedAuthenticationAuditServiceServer
// for forward compatibility
type AuthenticationAuditServiceServer interface {
	GetPublicKey(context.Context, *GetPublicKeyRequest) (*GetPublicKeyResponse, error)
	mustEmbedUnimplementedAuthenticationAuditServiceServer()
}

// UnimplementedAuthenticationAuditServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticationAuditServiceServer struct {
}

func (UnimplementedAuthenticationAuditServiceServer) GetPublicKey(context.Context, *GetPublicKeyRequest) (*GetPublicKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicKey not implemented")
}
func (UnimplementedAuthenticationAuditServiceServer) mustEmbedUnimplementedAuthenticationAuditServiceServer() {
}

// UnsafeAuthenticationAuditServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticationAuditServiceServer will
// result in compilation errors.
type UnsafeAuthenticationAuditServiceServer interface {
	mustEmbedUnimplementedAuthenticationAuditServiceServer()
}

func RegisterAuthenticationAuditServiceServer(s grpc.ServiceRegistrar, srv AuthenticationAuditServiceServer) {
	s.RegisterService(&AuthenticationAuditService_ServiceDesc, srv)
}

func _AuthenticationAuditService_GetPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublicKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationAuditServiceServer).GetPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthenticationAuditService/GetPublicKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationAuditServiceServer).GetPublicKey(ctx, req.(*GetPublicKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthenticationAuditService_ServiceDesc is the grpc.ServiceDesc for AuthenticationAuditService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthenticationAuditService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthenticationAuditService",
	HandlerType: (*AuthenticationAuditServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPublicKey",
			Handler:    _AuthenticationAuditService_GetPublicKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spoon/auth/api.proto",
}
