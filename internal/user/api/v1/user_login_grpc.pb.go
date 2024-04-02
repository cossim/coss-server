// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: api/grpc/v1/user_login.proto

package v1

import (
	context "context"
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
	UserLoginService_InsertUserLogin_FullMethodName                 = "/user_v1.UserLoginService/InsertUserLogin"
	UserLoginService_GetUserLoginByToken_FullMethodName             = "/user_v1.UserLoginService/GetUserLoginByToken"
	UserLoginService_GetUserLoginByDriverIdAndUserId_FullMethodName = "/user_v1.UserLoginService/GetUserLoginByDriverIdAndUserId"
	UserLoginService_UpdateUserLoginTokenByDriverId_FullMethodName  = "/user_v1.UserLoginService/UpdateUserLoginTokenByDriverId"
	UserLoginService_GetUserDriverTokenByUserId_FullMethodName      = "/user_v1.UserLoginService/GetUserDriverTokenByUserId"
	UserLoginService_GetUserLoginByUserId_FullMethodName            = "/user_v1.UserLoginService/GetUserLoginByUserId"
)

// UserLoginServiceClient is the client API for UserLoginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserLoginServiceClient interface {
	InsertUserLogin(ctx context.Context, in *UserLogin, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetUserLoginByToken(ctx context.Context, in *GetUserLoginByTokenRequest, opts ...grpc.CallOption) (*UserLogin, error)
	GetUserLoginByDriverIdAndUserId(ctx context.Context, in *DriverIdAndUserId, opts ...grpc.CallOption) (*UserLogin, error)
	UpdateUserLoginTokenByDriverId(ctx context.Context, in *TokenUpdate, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetUserDriverTokenByUserId(ctx context.Context, in *GetUserDriverTokenByUserIdRequest, opts ...grpc.CallOption) (*GetUserDriverTokenByUserIdResponse, error)
	// 根据用户ID获取登录信息
	GetUserLoginByUserId(ctx context.Context, in *GetUserLoginByUserIdRequest, opts ...grpc.CallOption) (*UserLogin, error)
}

type userLoginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserLoginServiceClient(cc grpc.ClientConnInterface) UserLoginServiceClient {
	return &userLoginServiceClient{cc}
}

func (c *userLoginServiceClient) InsertUserLogin(ctx context.Context, in *UserLogin, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserLoginService_InsertUserLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLoginServiceClient) GetUserLoginByToken(ctx context.Context, in *GetUserLoginByTokenRequest, opts ...grpc.CallOption) (*UserLogin, error) {
	out := new(UserLogin)
	err := c.cc.Invoke(ctx, UserLoginService_GetUserLoginByToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLoginServiceClient) GetUserLoginByDriverIdAndUserId(ctx context.Context, in *DriverIdAndUserId, opts ...grpc.CallOption) (*UserLogin, error) {
	out := new(UserLogin)
	err := c.cc.Invoke(ctx, UserLoginService_GetUserLoginByDriverIdAndUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLoginServiceClient) UpdateUserLoginTokenByDriverId(ctx context.Context, in *TokenUpdate, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserLoginService_UpdateUserLoginTokenByDriverId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLoginServiceClient) GetUserDriverTokenByUserId(ctx context.Context, in *GetUserDriverTokenByUserIdRequest, opts ...grpc.CallOption) (*GetUserDriverTokenByUserIdResponse, error) {
	out := new(GetUserDriverTokenByUserIdResponse)
	err := c.cc.Invoke(ctx, UserLoginService_GetUserDriverTokenByUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLoginServiceClient) GetUserLoginByUserId(ctx context.Context, in *GetUserLoginByUserIdRequest, opts ...grpc.CallOption) (*UserLogin, error) {
	out := new(UserLogin)
	err := c.cc.Invoke(ctx, UserLoginService_GetUserLoginByUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserLoginServiceServer is the server API for UserLoginService service.
// All implementations must embed UnimplementedUserLoginServiceServer
// for forward compatibility
type UserLoginServiceServer interface {
	InsertUserLogin(context.Context, *UserLogin) (*emptypb.Empty, error)
	GetUserLoginByToken(context.Context, *GetUserLoginByTokenRequest) (*UserLogin, error)
	GetUserLoginByDriverIdAndUserId(context.Context, *DriverIdAndUserId) (*UserLogin, error)
	UpdateUserLoginTokenByDriverId(context.Context, *TokenUpdate) (*emptypb.Empty, error)
	GetUserDriverTokenByUserId(context.Context, *GetUserDriverTokenByUserIdRequest) (*GetUserDriverTokenByUserIdResponse, error)
	// 根据用户ID获取登录信息
	GetUserLoginByUserId(context.Context, *GetUserLoginByUserIdRequest) (*UserLogin, error)
	mustEmbedUnimplementedUserLoginServiceServer()
}

// UnimplementedUserLoginServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserLoginServiceServer struct {
}

func (UnimplementedUserLoginServiceServer) InsertUserLogin(context.Context, *UserLogin) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertUserLogin not implemented")
}
func (UnimplementedUserLoginServiceServer) GetUserLoginByToken(context.Context, *GetUserLoginByTokenRequest) (*UserLogin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserLoginByToken not implemented")
}
func (UnimplementedUserLoginServiceServer) GetUserLoginByDriverIdAndUserId(context.Context, *DriverIdAndUserId) (*UserLogin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserLoginByDriverIdAndUserId not implemented")
}
func (UnimplementedUserLoginServiceServer) UpdateUserLoginTokenByDriverId(context.Context, *TokenUpdate) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserLoginTokenByDriverId not implemented")
}
func (UnimplementedUserLoginServiceServer) GetUserDriverTokenByUserId(context.Context, *GetUserDriverTokenByUserIdRequest) (*GetUserDriverTokenByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserDriverTokenByUserId not implemented")
}
func (UnimplementedUserLoginServiceServer) GetUserLoginByUserId(context.Context, *GetUserLoginByUserIdRequest) (*UserLogin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserLoginByUserId not implemented")
}
func (UnimplementedUserLoginServiceServer) mustEmbedUnimplementedUserLoginServiceServer() {}

// UnsafeUserLoginServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserLoginServiceServer will
// result in compilation errors.
type UnsafeUserLoginServiceServer interface {
	mustEmbedUnimplementedUserLoginServiceServer()
}

func RegisterUserLoginServiceServer(s grpc.ServiceRegistrar, srv UserLoginServiceServer) {
	s.RegisterService(&UserLoginService_ServiceDesc, srv)
}

func _UserLoginService_InsertUserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLogin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLoginServiceServer).InsertUserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserLoginService_InsertUserLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLoginServiceServer).InsertUserLogin(ctx, req.(*UserLogin))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserLoginService_GetUserLoginByToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserLoginByTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLoginServiceServer).GetUserLoginByToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserLoginService_GetUserLoginByToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLoginServiceServer).GetUserLoginByToken(ctx, req.(*GetUserLoginByTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserLoginService_GetUserLoginByDriverIdAndUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DriverIdAndUserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLoginServiceServer).GetUserLoginByDriverIdAndUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserLoginService_GetUserLoginByDriverIdAndUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLoginServiceServer).GetUserLoginByDriverIdAndUserId(ctx, req.(*DriverIdAndUserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserLoginService_UpdateUserLoginTokenByDriverId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLoginServiceServer).UpdateUserLoginTokenByDriverId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserLoginService_UpdateUserLoginTokenByDriverId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLoginServiceServer).UpdateUserLoginTokenByDriverId(ctx, req.(*TokenUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserLoginService_GetUserDriverTokenByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserDriverTokenByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLoginServiceServer).GetUserDriverTokenByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserLoginService_GetUserDriverTokenByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLoginServiceServer).GetUserDriverTokenByUserId(ctx, req.(*GetUserDriverTokenByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserLoginService_GetUserLoginByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserLoginByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLoginServiceServer).GetUserLoginByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserLoginService_GetUserLoginByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLoginServiceServer).GetUserLoginByUserId(ctx, req.(*GetUserLoginByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserLoginService_ServiceDesc is the grpc.ServiceDesc for UserLoginService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserLoginService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_v1.UserLoginService",
	HandlerType: (*UserLoginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertUserLogin",
			Handler:    _UserLoginService_InsertUserLogin_Handler,
		},
		{
			MethodName: "GetUserLoginByToken",
			Handler:    _UserLoginService_GetUserLoginByToken_Handler,
		},
		{
			MethodName: "GetUserLoginByDriverIdAndUserId",
			Handler:    _UserLoginService_GetUserLoginByDriverIdAndUserId_Handler,
		},
		{
			MethodName: "UpdateUserLoginTokenByDriverId",
			Handler:    _UserLoginService_UpdateUserLoginTokenByDriverId_Handler,
		},
		{
			MethodName: "GetUserDriverTokenByUserId",
			Handler:    _UserLoginService_GetUserDriverTokenByUserId_Handler,
		},
		{
			MethodName: "GetUserLoginByUserId",
			Handler:    _UserLoginService_GetUserLoginByUserId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/grpc/v1/user_login.proto",
}
