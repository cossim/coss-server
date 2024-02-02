// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: api/v1/user.proto

package v1

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
	UserService_UserLogin_FullMethodName               = "/v1.UserService/UserLogin"
	UserService_UserRegister_FullMethodName            = "/v1.UserService/UserRegister"
	UserService_UserInfo_FullMethodName                = "/v1.UserService/UserInfo"
	UserService_GetBatchUserInfo_FullMethodName        = "/v1.UserService/GetBatchUserInfo"
	UserService_GetUserInfoByEmail_FullMethodName      = "/v1.UserService/GetUserInfoByEmail"
	UserService_GetUserPublicKey_FullMethodName        = "/v1.UserService/GetUserPublicKey"
	UserService_SetUserPublicKey_FullMethodName        = "/v1.UserService/SetUserPublicKey"
	UserService_ModifyUserInfo_FullMethodName          = "/v1.UserService/ModifyUserInfo"
	UserService_ModifyUserPassword_FullMethodName      = "/v1.UserService/ModifyUserPassword"
	UserService_GetUserPasswordByUserId_FullMethodName = "/v1.UserService/GetUserPasswordByUserId"
	UserService_SetUserSecretBundle_FullMethodName     = "/v1.UserService/SetUserSecretBundle"
	UserService_GetUserSecretBundle_FullMethodName     = "/v1.UserService/GetUserSecretBundle"
	UserService_ActivateUser_FullMethodName            = "/v1.UserService/ActivateUser"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	// 用户登录
	UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
	// 用户注册
	UserRegister(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error)
	// 获取用户信息
	UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	// 批量获取用户信息
	GetBatchUserInfo(ctx context.Context, in *GetBatchUserInfoRequest, opts ...grpc.CallOption) (*GetBatchUserInfoResponse, error)
	// 根据email获取用户信息
	GetUserInfoByEmail(ctx context.Context, in *GetUserInfoByEmailRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	GetUserPublicKey(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*GetUserPublicKeyResponse, error)
	SetUserPublicKey(ctx context.Context, in *SetPublicKeyRequest, opts ...grpc.CallOption) (*UserResponse, error)
	ModifyUserInfo(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserResponse, error)
	ModifyUserPassword(ctx context.Context, in *ModifyUserPasswordRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GetUserPasswordByUserId(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*GetUserPasswordByUserIdResponse, error)
	// 设置用户密钥包
	SetUserSecretBundle(ctx context.Context, in *SetUserSecretBundleRequest, opts ...grpc.CallOption) (*SetUserSecretBundleResponse, error)
	// 获取用户密钥包
	GetUserSecretBundle(ctx context.Context, in *GetUserSecretBundleRequest, opts ...grpc.CallOption) (*GetUserSecretBundleResponse, error)
	ActivateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	out := new(UserLoginResponse)
	err := c.cc.Invoke(ctx, UserService_UserLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserRegister(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error) {
	out := new(UserRegisterResponse)
	err := c.cc.Invoke(ctx, UserService_UserRegister_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, UserService_UserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetBatchUserInfo(ctx context.Context, in *GetBatchUserInfoRequest, opts ...grpc.CallOption) (*GetBatchUserInfoResponse, error) {
	out := new(GetBatchUserInfoResponse)
	err := c.cc.Invoke(ctx, UserService_GetBatchUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserInfoByEmail(ctx context.Context, in *GetUserInfoByEmailRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, UserService_GetUserInfoByEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserPublicKey(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*GetUserPublicKeyResponse, error) {
	out := new(GetUserPublicKeyResponse)
	err := c.cc.Invoke(ctx, UserService_GetUserPublicKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SetUserPublicKey(ctx context.Context, in *SetPublicKeyRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UserService_SetUserPublicKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ModifyUserInfo(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UserService_ModifyUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ModifyUserPassword(ctx context.Context, in *ModifyUserPasswordRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UserService_ModifyUserPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserPasswordByUserId(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*GetUserPasswordByUserIdResponse, error) {
	out := new(GetUserPasswordByUserIdResponse)
	err := c.cc.Invoke(ctx, UserService_GetUserPasswordByUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SetUserSecretBundle(ctx context.Context, in *SetUserSecretBundleRequest, opts ...grpc.CallOption) (*SetUserSecretBundleResponse, error) {
	out := new(SetUserSecretBundleResponse)
	err := c.cc.Invoke(ctx, UserService_SetUserSecretBundle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserSecretBundle(ctx context.Context, in *GetUserSecretBundleRequest, opts ...grpc.CallOption) (*GetUserSecretBundleResponse, error) {
	out := new(GetUserSecretBundleResponse)
	err := c.cc.Invoke(ctx, UserService_GetUserSecretBundle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ActivateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UserService_ActivateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	// 用户登录
	UserLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error)
	// 用户注册
	UserRegister(context.Context, *UserRegisterRequest) (*UserRegisterResponse, error)
	// 获取用户信息
	UserInfo(context.Context, *UserInfoRequest) (*UserInfoResponse, error)
	// 批量获取用户信息
	GetBatchUserInfo(context.Context, *GetBatchUserInfoRequest) (*GetBatchUserInfoResponse, error)
	// 根据email获取用户信息
	GetUserInfoByEmail(context.Context, *GetUserInfoByEmailRequest) (*UserInfoResponse, error)
	GetUserPublicKey(context.Context, *UserRequest) (*GetUserPublicKeyResponse, error)
	SetUserPublicKey(context.Context, *SetPublicKeyRequest) (*UserResponse, error)
	ModifyUserInfo(context.Context, *User) (*UserResponse, error)
	ModifyUserPassword(context.Context, *ModifyUserPasswordRequest) (*UserResponse, error)
	GetUserPasswordByUserId(context.Context, *UserRequest) (*GetUserPasswordByUserIdResponse, error)
	// 设置用户密钥包
	SetUserSecretBundle(context.Context, *SetUserSecretBundleRequest) (*SetUserSecretBundleResponse, error)
	// 获取用户密钥包
	GetUserSecretBundle(context.Context, *GetUserSecretBundleRequest) (*GetUserSecretBundleResponse, error)
	ActivateUser(context.Context, *UserRequest) (*UserResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) UserLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedUserServiceServer) UserRegister(context.Context, *UserRegisterRequest) (*UserRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegister not implemented")
}
func (UnimplementedUserServiceServer) UserInfo(context.Context, *UserInfoRequest) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
func (UnimplementedUserServiceServer) GetBatchUserInfo(context.Context, *GetBatchUserInfoRequest) (*GetBatchUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBatchUserInfo not implemented")
}
func (UnimplementedUserServiceServer) GetUserInfoByEmail(context.Context, *GetUserInfoByEmailRequest) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfoByEmail not implemented")
}
func (UnimplementedUserServiceServer) GetUserPublicKey(context.Context, *UserRequest) (*GetUserPublicKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPublicKey not implemented")
}
func (UnimplementedUserServiceServer) SetUserPublicKey(context.Context, *SetPublicKeyRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserPublicKey not implemented")
}
func (UnimplementedUserServiceServer) ModifyUserInfo(context.Context, *User) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyUserInfo not implemented")
}
func (UnimplementedUserServiceServer) ModifyUserPassword(context.Context, *ModifyUserPasswordRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyUserPassword not implemented")
}
func (UnimplementedUserServiceServer) GetUserPasswordByUserId(context.Context, *UserRequest) (*GetUserPasswordByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPasswordByUserId not implemented")
}
func (UnimplementedUserServiceServer) SetUserSecretBundle(context.Context, *SetUserSecretBundleRequest) (*SetUserSecretBundleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserSecretBundle not implemented")
}
func (UnimplementedUserServiceServer) GetUserSecretBundle(context.Context, *GetUserSecretBundleRequest) (*GetUserSecretBundleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserSecretBundle not implemented")
}
func (UnimplementedUserServiceServer) ActivateUser(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateUser not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserLogin(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserRegister(ctx, req.(*UserRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserInfo(ctx, req.(*UserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetBatchUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBatchUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetBatchUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetBatchUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetBatchUserInfo(ctx, req.(*GetBatchUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserInfoByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoByEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserInfoByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserInfoByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserInfoByEmail(ctx, req.(*GetUserInfoByEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserPublicKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserPublicKey(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SetUserPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPublicKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SetUserPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_SetUserPublicKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SetUserPublicKey(ctx, req.(*SetPublicKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ModifyUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ModifyUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ModifyUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ModifyUserInfo(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ModifyUserPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyUserPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ModifyUserPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ModifyUserPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ModifyUserPassword(ctx, req.(*ModifyUserPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserPasswordByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserPasswordByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserPasswordByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserPasswordByUserId(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SetUserSecretBundle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserSecretBundleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SetUserSecretBundle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_SetUserSecretBundle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SetUserSecretBundle(ctx, req.(*SetUserSecretBundleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserSecretBundle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserSecretBundleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserSecretBundle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserSecretBundle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserSecretBundle(ctx, req.(*GetUserSecretBundleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ActivateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ActivateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ActivateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ActivateUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserLogin",
			Handler:    _UserService_UserLogin_Handler,
		},
		{
			MethodName: "UserRegister",
			Handler:    _UserService_UserRegister_Handler,
		},
		{
			MethodName: "UserInfo",
			Handler:    _UserService_UserInfo_Handler,
		},
		{
			MethodName: "GetBatchUserInfo",
			Handler:    _UserService_GetBatchUserInfo_Handler,
		},
		{
			MethodName: "GetUserInfoByEmail",
			Handler:    _UserService_GetUserInfoByEmail_Handler,
		},
		{
			MethodName: "GetUserPublicKey",
			Handler:    _UserService_GetUserPublicKey_Handler,
		},
		{
			MethodName: "SetUserPublicKey",
			Handler:    _UserService_SetUserPublicKey_Handler,
		},
		{
			MethodName: "ModifyUserInfo",
			Handler:    _UserService_ModifyUserInfo_Handler,
		},
		{
			MethodName: "ModifyUserPassword",
			Handler:    _UserService_ModifyUserPassword_Handler,
		},
		{
			MethodName: "GetUserPasswordByUserId",
			Handler:    _UserService_GetUserPasswordByUserId_Handler,
		},
		{
			MethodName: "SetUserSecretBundle",
			Handler:    _UserService_SetUserSecretBundle_Handler,
		},
		{
			MethodName: "GetUserSecretBundle",
			Handler:    _UserService_GetUserSecretBundle_Handler,
		},
		{
			MethodName: "ActivateUser",
			Handler:    _UserService_ActivateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/user.proto",
}
