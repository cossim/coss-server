// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: api/v1/user_relation.proto

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
	UserRelationService_AddFriend_FullMethodName                   = "/v1.UserRelationService/AddFriend"
	UserRelationService_ManageFriend_FullMethodName                = "/v1.UserRelationService/ManageFriend"
	UserRelationService_ManageFriendRevert_FullMethodName          = "/v1.UserRelationService/ManageFriendRevert"
	UserRelationService_DeleteFriend_FullMethodName                = "/v1.UserRelationService/DeleteFriend"
	UserRelationService_DeleteFriendRevert_FullMethodName          = "/v1.UserRelationService/DeleteFriendRevert"
	UserRelationService_AddBlacklist_FullMethodName                = "/v1.UserRelationService/AddBlacklist"
	UserRelationService_DeleteBlacklist_FullMethodName             = "/v1.UserRelationService/DeleteBlacklist"
	UserRelationService_GetFriendList_FullMethodName               = "/v1.UserRelationService/GetFriendList"
	UserRelationService_GetBlacklist_FullMethodName                = "/v1.UserRelationService/GetBlacklist"
	UserRelationService_GetUserRelation_FullMethodName             = "/v1.UserRelationService/GetUserRelation"
	UserRelationService_GetUserRelationByUserIds_FullMethodName    = "/v1.UserRelationService/GetUserRelationByUserIds"
	UserRelationService_SetFriendSilentNotification_FullMethodName = "/v1.UserRelationService/SetFriendSilentNotification"
	UserRelationService_SetUserOpenBurnAfterReading_FullMethodName = "/v1.UserRelationService/SetUserOpenBurnAfterReading"
	UserRelationService_SetFriendRemark_FullMethodName             = "/v1.UserRelationService/SetFriendRemark"
)

// UserRelationServiceClient is the client API for UserRelationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserRelationServiceClient interface {
	// 添加好友
	AddFriend(ctx context.Context, in *AddFriendRequest, opts ...grpc.CallOption) (*AddFriendResponse, error)
	// 管理好友
	ManageFriend(ctx context.Context, in *ManageFriendRequest, opts ...grpc.CallOption) (*ManageFriendResponse, error)
	// 管理好友回滚
	ManageFriendRevert(ctx context.Context, in *ManageFriendRequest, opts ...grpc.CallOption) (*ManageFriendResponse, error)
	// 删除好友
	DeleteFriend(ctx context.Context, in *DeleteFriendRequest, opts ...grpc.CallOption) (*DeleteFriendResponse, error)
	// 删除好友回滚
	DeleteFriendRevert(ctx context.Context, in *DeleteFriendRequest, opts ...grpc.CallOption) (*DeleteFriendResponse, error)
	// 添加黑名单
	AddBlacklist(ctx context.Context, in *AddBlacklistRequest, opts ...grpc.CallOption) (*AddBlacklistResponse, error)
	// 移出黑名单
	DeleteBlacklist(ctx context.Context, in *DeleteBlacklistRequest, opts ...grpc.CallOption) (*DeleteBlacklistResponse, error)
	// 获取好友列表
	GetFriendList(ctx context.Context, in *GetFriendListRequest, opts ...grpc.CallOption) (*GetFriendListResponse, error)
	// 获取黑名单列表
	GetBlacklist(ctx context.Context, in *GetBlacklistRequest, opts ...grpc.CallOption) (*GetBlacklistResponse, error)
	// 获取用户关系
	GetUserRelation(ctx context.Context, in *GetUserRelationRequest, opts ...grpc.CallOption) (*GetUserRelationResponse, error)
	// 批量获取用户关系
	GetUserRelationByUserIds(ctx context.Context, in *GetUserRelationByUserIdsRequest, opts ...grpc.CallOption) (*GetUserRelationByUserIdsResponse, error)
	// 设置好友静默通知
	SetFriendSilentNotification(ctx context.Context, in *SetFriendSilentNotificationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 设置阅后即焚开关
	SetUserOpenBurnAfterReading(ctx context.Context, in *SetUserOpenBurnAfterReadingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 设置好友备注
	SetFriendRemark(ctx context.Context, in *SetFriendRemarkRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type userRelationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserRelationServiceClient(cc grpc.ClientConnInterface) UserRelationServiceClient {
	return &userRelationServiceClient{cc}
}

func (c *userRelationServiceClient) AddFriend(ctx context.Context, in *AddFriendRequest, opts ...grpc.CallOption) (*AddFriendResponse, error) {
	out := new(AddFriendResponse)
	err := c.cc.Invoke(ctx, UserRelationService_AddFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRelationServiceClient) ManageFriend(ctx context.Context, in *ManageFriendRequest, opts ...grpc.CallOption) (*ManageFriendResponse, error) {
	out := new(ManageFriendResponse)
	err := c.cc.Invoke(ctx, UserRelationService_ManageFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRelationServiceClient) ManageFriendRevert(ctx context.Context, in *ManageFriendRequest, opts ...grpc.CallOption) (*ManageFriendResponse, error) {
	out := new(ManageFriendResponse)
	err := c.cc.Invoke(ctx, UserRelationService_ManageFriendRevert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRelationServiceClient) DeleteFriend(ctx context.Context, in *DeleteFriendRequest, opts ...grpc.CallOption) (*DeleteFriendResponse, error) {
	out := new(DeleteFriendResponse)
	err := c.cc.Invoke(ctx, UserRelationService_DeleteFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRelationServiceClient) DeleteFriendRevert(ctx context.Context, in *DeleteFriendRequest, opts ...grpc.CallOption) (*DeleteFriendResponse, error) {
	out := new(DeleteFriendResponse)
	err := c.cc.Invoke(ctx, UserRelationService_DeleteFriendRevert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRelationServiceClient) AddBlacklist(ctx context.Context, in *AddBlacklistRequest, opts ...grpc.CallOption) (*AddBlacklistResponse, error) {
	out := new(AddBlacklistResponse)
	err := c.cc.Invoke(ctx, UserRelationService_AddBlacklist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRelationServiceClient) DeleteBlacklist(ctx context.Context, in *DeleteBlacklistRequest, opts ...grpc.CallOption) (*DeleteBlacklistResponse, error) {
	out := new(DeleteBlacklistResponse)
	err := c.cc.Invoke(ctx, UserRelationService_DeleteBlacklist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRelationServiceClient) GetFriendList(ctx context.Context, in *GetFriendListRequest, opts ...grpc.CallOption) (*GetFriendListResponse, error) {
	out := new(GetFriendListResponse)
	err := c.cc.Invoke(ctx, UserRelationService_GetFriendList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRelationServiceClient) GetBlacklist(ctx context.Context, in *GetBlacklistRequest, opts ...grpc.CallOption) (*GetBlacklistResponse, error) {
	out := new(GetBlacklistResponse)
	err := c.cc.Invoke(ctx, UserRelationService_GetBlacklist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRelationServiceClient) GetUserRelation(ctx context.Context, in *GetUserRelationRequest, opts ...grpc.CallOption) (*GetUserRelationResponse, error) {
	out := new(GetUserRelationResponse)
	err := c.cc.Invoke(ctx, UserRelationService_GetUserRelation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRelationServiceClient) GetUserRelationByUserIds(ctx context.Context, in *GetUserRelationByUserIdsRequest, opts ...grpc.CallOption) (*GetUserRelationByUserIdsResponse, error) {
	out := new(GetUserRelationByUserIdsResponse)
	err := c.cc.Invoke(ctx, UserRelationService_GetUserRelationByUserIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRelationServiceClient) SetFriendSilentNotification(ctx context.Context, in *SetFriendSilentNotificationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserRelationService_SetFriendSilentNotification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRelationServiceClient) SetUserOpenBurnAfterReading(ctx context.Context, in *SetUserOpenBurnAfterReadingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserRelationService_SetUserOpenBurnAfterReading_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRelationServiceClient) SetFriendRemark(ctx context.Context, in *SetFriendRemarkRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserRelationService_SetFriendRemark_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserRelationServiceServer is the server API for UserRelationService service.
// All implementations must embed UnimplementedUserRelationServiceServer
// for forward compatibility
type UserRelationServiceServer interface {
	// 添加好友
	AddFriend(context.Context, *AddFriendRequest) (*AddFriendResponse, error)
	// 管理好友
	ManageFriend(context.Context, *ManageFriendRequest) (*ManageFriendResponse, error)
	// 管理好友回滚
	ManageFriendRevert(context.Context, *ManageFriendRequest) (*ManageFriendResponse, error)
	// 删除好友
	DeleteFriend(context.Context, *DeleteFriendRequest) (*DeleteFriendResponse, error)
	// 删除好友回滚
	DeleteFriendRevert(context.Context, *DeleteFriendRequest) (*DeleteFriendResponse, error)
	// 添加黑名单
	AddBlacklist(context.Context, *AddBlacklistRequest) (*AddBlacklistResponse, error)
	// 移出黑名单
	DeleteBlacklist(context.Context, *DeleteBlacklistRequest) (*DeleteBlacklistResponse, error)
	// 获取好友列表
	GetFriendList(context.Context, *GetFriendListRequest) (*GetFriendListResponse, error)
	// 获取黑名单列表
	GetBlacklist(context.Context, *GetBlacklistRequest) (*GetBlacklistResponse, error)
	// 获取用户关系
	GetUserRelation(context.Context, *GetUserRelationRequest) (*GetUserRelationResponse, error)
	// 批量获取用户关系
	GetUserRelationByUserIds(context.Context, *GetUserRelationByUserIdsRequest) (*GetUserRelationByUserIdsResponse, error)
	// 设置好友静默通知
	SetFriendSilentNotification(context.Context, *SetFriendSilentNotificationRequest) (*emptypb.Empty, error)
	// 设置阅后即焚开关
	SetUserOpenBurnAfterReading(context.Context, *SetUserOpenBurnAfterReadingRequest) (*emptypb.Empty, error)
	// 设置好友备注
	SetFriendRemark(context.Context, *SetFriendRemarkRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedUserRelationServiceServer()
}

// UnimplementedUserRelationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserRelationServiceServer struct {
}

func (UnimplementedUserRelationServiceServer) AddFriend(context.Context, *AddFriendRequest) (*AddFriendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFriend not implemented")
}
func (UnimplementedUserRelationServiceServer) ManageFriend(context.Context, *ManageFriendRequest) (*ManageFriendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ManageFriend not implemented")
}
func (UnimplementedUserRelationServiceServer) ManageFriendRevert(context.Context, *ManageFriendRequest) (*ManageFriendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ManageFriendRevert not implemented")
}
func (UnimplementedUserRelationServiceServer) DeleteFriend(context.Context, *DeleteFriendRequest) (*DeleteFriendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFriend not implemented")
}
func (UnimplementedUserRelationServiceServer) DeleteFriendRevert(context.Context, *DeleteFriendRequest) (*DeleteFriendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFriendRevert not implemented")
}
func (UnimplementedUserRelationServiceServer) AddBlacklist(context.Context, *AddBlacklistRequest) (*AddBlacklistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBlacklist not implemented")
}
func (UnimplementedUserRelationServiceServer) DeleteBlacklist(context.Context, *DeleteBlacklistRequest) (*DeleteBlacklistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBlacklist not implemented")
}
func (UnimplementedUserRelationServiceServer) GetFriendList(context.Context, *GetFriendListRequest) (*GetFriendListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriendList not implemented")
}
func (UnimplementedUserRelationServiceServer) GetBlacklist(context.Context, *GetBlacklistRequest) (*GetBlacklistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlacklist not implemented")
}
func (UnimplementedUserRelationServiceServer) GetUserRelation(context.Context, *GetUserRelationRequest) (*GetUserRelationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserRelation not implemented")
}
func (UnimplementedUserRelationServiceServer) GetUserRelationByUserIds(context.Context, *GetUserRelationByUserIdsRequest) (*GetUserRelationByUserIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserRelationByUserIds not implemented")
}
func (UnimplementedUserRelationServiceServer) SetFriendSilentNotification(context.Context, *SetFriendSilentNotificationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetFriendSilentNotification not implemented")
}
func (UnimplementedUserRelationServiceServer) SetUserOpenBurnAfterReading(context.Context, *SetUserOpenBurnAfterReadingRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserOpenBurnAfterReading not implemented")
}
func (UnimplementedUserRelationServiceServer) SetFriendRemark(context.Context, *SetFriendRemarkRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetFriendRemark not implemented")
}
func (UnimplementedUserRelationServiceServer) mustEmbedUnimplementedUserRelationServiceServer() {}

// UnsafeUserRelationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserRelationServiceServer will
// result in compilation errors.
type UnsafeUserRelationServiceServer interface {
	mustEmbedUnimplementedUserRelationServiceServer()
}

func RegisterUserRelationServiceServer(s grpc.ServiceRegistrar, srv UserRelationServiceServer) {
	s.RegisterService(&UserRelationService_ServiceDesc, srv)
}

func _UserRelationService_AddFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRelationServiceServer).AddFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserRelationService_AddFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRelationServiceServer).AddFriend(ctx, req.(*AddFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRelationService_ManageFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManageFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRelationServiceServer).ManageFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserRelationService_ManageFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRelationServiceServer).ManageFriend(ctx, req.(*ManageFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRelationService_ManageFriendRevert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManageFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRelationServiceServer).ManageFriendRevert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserRelationService_ManageFriendRevert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRelationServiceServer).ManageFriendRevert(ctx, req.(*ManageFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRelationService_DeleteFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRelationServiceServer).DeleteFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserRelationService_DeleteFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRelationServiceServer).DeleteFriend(ctx, req.(*DeleteFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRelationService_DeleteFriendRevert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRelationServiceServer).DeleteFriendRevert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserRelationService_DeleteFriendRevert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRelationServiceServer).DeleteFriendRevert(ctx, req.(*DeleteFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRelationService_AddBlacklist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBlacklistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRelationServiceServer).AddBlacklist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserRelationService_AddBlacklist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRelationServiceServer).AddBlacklist(ctx, req.(*AddBlacklistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRelationService_DeleteBlacklist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBlacklistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRelationServiceServer).DeleteBlacklist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserRelationService_DeleteBlacklist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRelationServiceServer).DeleteBlacklist(ctx, req.(*DeleteBlacklistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRelationService_GetFriendList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFriendListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRelationServiceServer).GetFriendList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserRelationService_GetFriendList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRelationServiceServer).GetFriendList(ctx, req.(*GetFriendListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRelationService_GetBlacklist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlacklistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRelationServiceServer).GetBlacklist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserRelationService_GetBlacklist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRelationServiceServer).GetBlacklist(ctx, req.(*GetBlacklistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRelationService_GetUserRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRelationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRelationServiceServer).GetUserRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserRelationService_GetUserRelation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRelationServiceServer).GetUserRelation(ctx, req.(*GetUserRelationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRelationService_GetUserRelationByUserIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRelationByUserIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRelationServiceServer).GetUserRelationByUserIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserRelationService_GetUserRelationByUserIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRelationServiceServer).GetUserRelationByUserIds(ctx, req.(*GetUserRelationByUserIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRelationService_SetFriendSilentNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetFriendSilentNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRelationServiceServer).SetFriendSilentNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserRelationService_SetFriendSilentNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRelationServiceServer).SetFriendSilentNotification(ctx, req.(*SetFriendSilentNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRelationService_SetUserOpenBurnAfterReading_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserOpenBurnAfterReadingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRelationServiceServer).SetUserOpenBurnAfterReading(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserRelationService_SetUserOpenBurnAfterReading_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRelationServiceServer).SetUserOpenBurnAfterReading(ctx, req.(*SetUserOpenBurnAfterReadingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRelationService_SetFriendRemark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetFriendRemarkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRelationServiceServer).SetFriendRemark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserRelationService_SetFriendRemark_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRelationServiceServer).SetFriendRemark(ctx, req.(*SetFriendRemarkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserRelationService_ServiceDesc is the grpc.ServiceDesc for UserRelationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserRelationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.UserRelationService",
	HandlerType: (*UserRelationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFriend",
			Handler:    _UserRelationService_AddFriend_Handler,
		},
		{
			MethodName: "ManageFriend",
			Handler:    _UserRelationService_ManageFriend_Handler,
		},
		{
			MethodName: "ManageFriendRevert",
			Handler:    _UserRelationService_ManageFriendRevert_Handler,
		},
		{
			MethodName: "DeleteFriend",
			Handler:    _UserRelationService_DeleteFriend_Handler,
		},
		{
			MethodName: "DeleteFriendRevert",
			Handler:    _UserRelationService_DeleteFriendRevert_Handler,
		},
		{
			MethodName: "AddBlacklist",
			Handler:    _UserRelationService_AddBlacklist_Handler,
		},
		{
			MethodName: "DeleteBlacklist",
			Handler:    _UserRelationService_DeleteBlacklist_Handler,
		},
		{
			MethodName: "GetFriendList",
			Handler:    _UserRelationService_GetFriendList_Handler,
		},
		{
			MethodName: "GetBlacklist",
			Handler:    _UserRelationService_GetBlacklist_Handler,
		},
		{
			MethodName: "GetUserRelation",
			Handler:    _UserRelationService_GetUserRelation_Handler,
		},
		{
			MethodName: "GetUserRelationByUserIds",
			Handler:    _UserRelationService_GetUserRelationByUserIds_Handler,
		},
		{
			MethodName: "SetFriendSilentNotification",
			Handler:    _UserRelationService_SetFriendSilentNotification_Handler,
		},
		{
			MethodName: "SetUserOpenBurnAfterReading",
			Handler:    _UserRelationService_SetUserOpenBurnAfterReading_Handler,
		},
		{
			MethodName: "SetFriendRemark",
			Handler:    _UserRelationService_SetFriendRemark_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/user_relation.proto",
}
