// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: api/v1/group_relation.proto

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
	GroupRelationService_GetGroupUserIDs_FullMethodName                             = "/v1.GroupRelationService/GetGroupUserIDs"
	GroupRelationService_GetUserGroupIDs_FullMethodName                             = "/v1.GroupRelationService/GetUserGroupIDs"
	GroupRelationService_GetUserGroupList_FullMethodName                            = "/v1.GroupRelationService/GetUserGroupList"
	GroupRelationService_GetGroupAdminIds_FullMethodName                            = "/v1.GroupRelationService/GetGroupAdminIds"
	GroupRelationService_GetUserManageGroupID_FullMethodName                        = "/v1.GroupRelationService/GetUserManageGroupID"
	GroupRelationService_RemoveFromGroup_FullMethodName                             = "/v1.GroupRelationService/RemoveFromGroup"
	GroupRelationService_LeaveGroup_FullMethodName                                  = "/v1.GroupRelationService/LeaveGroup"
	GroupRelationService_LeaveGroupRevert_FullMethodName                            = "/v1.GroupRelationService/LeaveGroupRevert"
	GroupRelationService_GetGroupRelation_FullMethodName                            = "/v1.GroupRelationService/GetGroupRelation"
	GroupRelationService_GetBatchGroupRelation_FullMethodName                       = "/v1.GroupRelationService/GetBatchGroupRelation"
	GroupRelationService_DeleteGroupRelationByGroupId_FullMethodName                = "/v1.GroupRelationService/DeleteGroupRelationByGroupId"
	GroupRelationService_DeleteGroupRelationByGroupIdRevert_FullMethodName          = "/v1.GroupRelationService/DeleteGroupRelationByGroupIdRevert"
	GroupRelationService_DeleteGroupRelationByGroupIdAndUserID_FullMethodName       = "/v1.GroupRelationService/DeleteGroupRelationByGroupIdAndUserID"
	GroupRelationService_DeleteGroupRelationByGroupIdAndUserIDRevert_FullMethodName = "/v1.GroupRelationService/DeleteGroupRelationByGroupIdAndUserIDRevert"
	GroupRelationService_CreateGroupAndInviteUsers_FullMethodName                   = "/v1.GroupRelationService/CreateGroupAndInviteUsers"
	GroupRelationService_CreateGroupAndInviteUsersRevert_FullMethodName             = "/v1.GroupRelationService/CreateGroupAndInviteUsersRevert"
	GroupRelationService_SetGroupSilentNotification_FullMethodName                  = "/v1.GroupRelationService/SetGroupSilentNotification"
	GroupRelationService_RemoveGroupRelationByGroupIdAndUserIDs_FullMethodName      = "/v1.GroupRelationService/RemoveGroupRelationByGroupIdAndUserIDs"
	GroupRelationService_SetGroupOpenBurnAfterReading_FullMethodName                = "/v1.GroupRelationService/SetGroupOpenBurnAfterReading"
)

// GroupRelationServiceClient is the client API for GroupRelationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupRelationServiceClient interface {
	// 获取群聊成员ID列表
	GetGroupUserIDs(ctx context.Context, in *GroupIDRequest, opts ...grpc.CallOption) (*UserIdsResponse, error)
	// 获取用户的所有群聊ID列表
	GetUserGroupIDs(ctx context.Context, in *GetUserGroupIDsRequest, opts ...grpc.CallOption) (*GetUserGroupIDsResponse, error)
	// 获取用户的所有群聊列表信息
	GetUserGroupList(ctx context.Context, in *GetUserGroupListRequest, opts ...grpc.CallOption) (*GetUserGroupListResponse, error)
	// 获取群聊管理员ID列表
	GetGroupAdminIds(ctx context.Context, in *GroupIDRequest, opts ...grpc.CallOption) (*UserIdsResponse, error)
	// 获取用户管理的群聊ID列表
	GetUserManageGroupID(ctx context.Context, in *GetUserManageGroupIDRequest, opts ...grpc.CallOption) (*GetUserManageGroupIDResponse, error)
	// 从群聊中移除成员
	RemoveFromGroup(ctx context.Context, in *RemoveFromGroupRequest, opts ...grpc.CallOption) (*RemoveFromGroupResponse, error)
	// 退出群聊
	LeaveGroup(ctx context.Context, in *LeaveGroupRequest, opts ...grpc.CallOption) (*LeaveGroupResponse, error)
	// 退出群聊回滚操作
	LeaveGroupRevert(ctx context.Context, in *LeaveGroupRequest, opts ...grpc.CallOption) (*LeaveGroupResponse, error)
	// 获取用户与群聊关系信息
	GetGroupRelation(ctx context.Context, in *GetGroupRelationRequest, opts ...grpc.CallOption) (*GetGroupRelationResponse, error)
	// 批量获取用户与群聊关系信息
	GetBatchGroupRelation(ctx context.Context, in *GetBatchGroupRelationRequest, opts ...grpc.CallOption) (*GetBatchGroupRelationResponse, error)
	// 根据群聊ID删除群聊的所有关系
	DeleteGroupRelationByGroupId(ctx context.Context, in *GroupIDRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// DeleteGroupRelationByGroupId回滚操作
	DeleteGroupRelationByGroupIdRevert(ctx context.Context, in *GroupIDRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 根据群聊ID和用户ID删除用户的群聊关系
	DeleteGroupRelationByGroupIdAndUserID(ctx context.Context, in *DeleteGroupRelationByGroupIdAndUserIDRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// DeleteGroupRelationByGroupIdAndUserID回滚操作
	DeleteGroupRelationByGroupIdAndUserIDRevert(ctx context.Context, in *DeleteGroupRelationByGroupIdAndUserIDRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 创建群聊并邀请多个用户进群（包括自己）
	CreateGroupAndInviteUsers(ctx context.Context, in *CreateGroupAndInviteUsersRequest, opts ...grpc.CallOption) (*CreateGroupAndInviteUsersResponse, error)
	// CreateGroupAndInviteUsers 回滚操作
	CreateGroupAndInviteUsersRevert(ctx context.Context, in *CreateGroupAndInviteUsersRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 设置群聊为静默通知状态
	SetGroupSilentNotification(ctx context.Context, in *SetGroupSilentNotificationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 移除多个用户
	RemoveGroupRelationByGroupIdAndUserIDs(ctx context.Context, in *RemoveGroupRelationByGroupIdAndUserIDsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 设置阅后即焚开关
	SetGroupOpenBurnAfterReading(ctx context.Context, in *SetGroupOpenBurnAfterReadingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type groupRelationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupRelationServiceClient(cc grpc.ClientConnInterface) GroupRelationServiceClient {
	return &groupRelationServiceClient{cc}
}

func (c *groupRelationServiceClient) GetGroupUserIDs(ctx context.Context, in *GroupIDRequest, opts ...grpc.CallOption) (*UserIdsResponse, error) {
	out := new(UserIdsResponse)
	err := c.cc.Invoke(ctx, GroupRelationService_GetGroupUserIDs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupRelationServiceClient) GetUserGroupIDs(ctx context.Context, in *GetUserGroupIDsRequest, opts ...grpc.CallOption) (*GetUserGroupIDsResponse, error) {
	out := new(GetUserGroupIDsResponse)
	err := c.cc.Invoke(ctx, GroupRelationService_GetUserGroupIDs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupRelationServiceClient) GetUserGroupList(ctx context.Context, in *GetUserGroupListRequest, opts ...grpc.CallOption) (*GetUserGroupListResponse, error) {
	out := new(GetUserGroupListResponse)
	err := c.cc.Invoke(ctx, GroupRelationService_GetUserGroupList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupRelationServiceClient) GetGroupAdminIds(ctx context.Context, in *GroupIDRequest, opts ...grpc.CallOption) (*UserIdsResponse, error) {
	out := new(UserIdsResponse)
	err := c.cc.Invoke(ctx, GroupRelationService_GetGroupAdminIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupRelationServiceClient) GetUserManageGroupID(ctx context.Context, in *GetUserManageGroupIDRequest, opts ...grpc.CallOption) (*GetUserManageGroupIDResponse, error) {
	out := new(GetUserManageGroupIDResponse)
	err := c.cc.Invoke(ctx, GroupRelationService_GetUserManageGroupID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupRelationServiceClient) RemoveFromGroup(ctx context.Context, in *RemoveFromGroupRequest, opts ...grpc.CallOption) (*RemoveFromGroupResponse, error) {
	out := new(RemoveFromGroupResponse)
	err := c.cc.Invoke(ctx, GroupRelationService_RemoveFromGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupRelationServiceClient) LeaveGroup(ctx context.Context, in *LeaveGroupRequest, opts ...grpc.CallOption) (*LeaveGroupResponse, error) {
	out := new(LeaveGroupResponse)
	err := c.cc.Invoke(ctx, GroupRelationService_LeaveGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupRelationServiceClient) LeaveGroupRevert(ctx context.Context, in *LeaveGroupRequest, opts ...grpc.CallOption) (*LeaveGroupResponse, error) {
	out := new(LeaveGroupResponse)
	err := c.cc.Invoke(ctx, GroupRelationService_LeaveGroupRevert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupRelationServiceClient) GetGroupRelation(ctx context.Context, in *GetGroupRelationRequest, opts ...grpc.CallOption) (*GetGroupRelationResponse, error) {
	out := new(GetGroupRelationResponse)
	err := c.cc.Invoke(ctx, GroupRelationService_GetGroupRelation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupRelationServiceClient) GetBatchGroupRelation(ctx context.Context, in *GetBatchGroupRelationRequest, opts ...grpc.CallOption) (*GetBatchGroupRelationResponse, error) {
	out := new(GetBatchGroupRelationResponse)
	err := c.cc.Invoke(ctx, GroupRelationService_GetBatchGroupRelation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupRelationServiceClient) DeleteGroupRelationByGroupId(ctx context.Context, in *GroupIDRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GroupRelationService_DeleteGroupRelationByGroupId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupRelationServiceClient) DeleteGroupRelationByGroupIdRevert(ctx context.Context, in *GroupIDRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GroupRelationService_DeleteGroupRelationByGroupIdRevert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupRelationServiceClient) DeleteGroupRelationByGroupIdAndUserID(ctx context.Context, in *DeleteGroupRelationByGroupIdAndUserIDRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GroupRelationService_DeleteGroupRelationByGroupIdAndUserID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupRelationServiceClient) DeleteGroupRelationByGroupIdAndUserIDRevert(ctx context.Context, in *DeleteGroupRelationByGroupIdAndUserIDRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GroupRelationService_DeleteGroupRelationByGroupIdAndUserIDRevert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupRelationServiceClient) CreateGroupAndInviteUsers(ctx context.Context, in *CreateGroupAndInviteUsersRequest, opts ...grpc.CallOption) (*CreateGroupAndInviteUsersResponse, error) {
	out := new(CreateGroupAndInviteUsersResponse)
	err := c.cc.Invoke(ctx, GroupRelationService_CreateGroupAndInviteUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupRelationServiceClient) CreateGroupAndInviteUsersRevert(ctx context.Context, in *CreateGroupAndInviteUsersRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GroupRelationService_CreateGroupAndInviteUsersRevert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupRelationServiceClient) SetGroupSilentNotification(ctx context.Context, in *SetGroupSilentNotificationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GroupRelationService_SetGroupSilentNotification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupRelationServiceClient) RemoveGroupRelationByGroupIdAndUserIDs(ctx context.Context, in *RemoveGroupRelationByGroupIdAndUserIDsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GroupRelationService_RemoveGroupRelationByGroupIdAndUserIDs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupRelationServiceClient) SetGroupOpenBurnAfterReading(ctx context.Context, in *SetGroupOpenBurnAfterReadingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GroupRelationService_SetGroupOpenBurnAfterReading_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupRelationServiceServer is the server API for GroupRelationService service.
// All implementations must embed UnimplementedGroupRelationServiceServer
// for forward compatibility
type GroupRelationServiceServer interface {
	// 获取群聊成员ID列表
	GetGroupUserIDs(context.Context, *GroupIDRequest) (*UserIdsResponse, error)
	// 获取用户的所有群聊ID列表
	GetUserGroupIDs(context.Context, *GetUserGroupIDsRequest) (*GetUserGroupIDsResponse, error)
	// 获取用户的所有群聊列表信息
	GetUserGroupList(context.Context, *GetUserGroupListRequest) (*GetUserGroupListResponse, error)
	// 获取群聊管理员ID列表
	GetGroupAdminIds(context.Context, *GroupIDRequest) (*UserIdsResponse, error)
	// 获取用户管理的群聊ID列表
	GetUserManageGroupID(context.Context, *GetUserManageGroupIDRequest) (*GetUserManageGroupIDResponse, error)
	// 从群聊中移除成员
	RemoveFromGroup(context.Context, *RemoveFromGroupRequest) (*RemoveFromGroupResponse, error)
	// 退出群聊
	LeaveGroup(context.Context, *LeaveGroupRequest) (*LeaveGroupResponse, error)
	// 退出群聊回滚操作
	LeaveGroupRevert(context.Context, *LeaveGroupRequest) (*LeaveGroupResponse, error)
	// 获取用户与群聊关系信息
	GetGroupRelation(context.Context, *GetGroupRelationRequest) (*GetGroupRelationResponse, error)
	// 批量获取用户与群聊关系信息
	GetBatchGroupRelation(context.Context, *GetBatchGroupRelationRequest) (*GetBatchGroupRelationResponse, error)
	// 根据群聊ID删除群聊的所有关系
	DeleteGroupRelationByGroupId(context.Context, *GroupIDRequest) (*emptypb.Empty, error)
	// DeleteGroupRelationByGroupId回滚操作
	DeleteGroupRelationByGroupIdRevert(context.Context, *GroupIDRequest) (*emptypb.Empty, error)
	// 根据群聊ID和用户ID删除用户的群聊关系
	DeleteGroupRelationByGroupIdAndUserID(context.Context, *DeleteGroupRelationByGroupIdAndUserIDRequest) (*emptypb.Empty, error)
	// DeleteGroupRelationByGroupIdAndUserID回滚操作
	DeleteGroupRelationByGroupIdAndUserIDRevert(context.Context, *DeleteGroupRelationByGroupIdAndUserIDRequest) (*emptypb.Empty, error)
	// 创建群聊并邀请多个用户进群（包括自己）
	CreateGroupAndInviteUsers(context.Context, *CreateGroupAndInviteUsersRequest) (*CreateGroupAndInviteUsersResponse, error)
	// CreateGroupAndInviteUsers 回滚操作
	CreateGroupAndInviteUsersRevert(context.Context, *CreateGroupAndInviteUsersRequest) (*emptypb.Empty, error)
	// 设置群聊为静默通知状态
	SetGroupSilentNotification(context.Context, *SetGroupSilentNotificationRequest) (*emptypb.Empty, error)
	// 移除多个用户
	RemoveGroupRelationByGroupIdAndUserIDs(context.Context, *RemoveGroupRelationByGroupIdAndUserIDsRequest) (*emptypb.Empty, error)
	// 设置阅后即焚开关
	SetGroupOpenBurnAfterReading(context.Context, *SetGroupOpenBurnAfterReadingRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedGroupRelationServiceServer()
}

// UnimplementedGroupRelationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGroupRelationServiceServer struct {
}

func (UnimplementedGroupRelationServiceServer) GetGroupUserIDs(context.Context, *GroupIDRequest) (*UserIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupUserIDs not implemented")
}
func (UnimplementedGroupRelationServiceServer) GetUserGroupIDs(context.Context, *GetUserGroupIDsRequest) (*GetUserGroupIDsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserGroupIDs not implemented")
}
func (UnimplementedGroupRelationServiceServer) GetUserGroupList(context.Context, *GetUserGroupListRequest) (*GetUserGroupListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserGroupList not implemented")
}
func (UnimplementedGroupRelationServiceServer) GetGroupAdminIds(context.Context, *GroupIDRequest) (*UserIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupAdminIds not implemented")
}
func (UnimplementedGroupRelationServiceServer) GetUserManageGroupID(context.Context, *GetUserManageGroupIDRequest) (*GetUserManageGroupIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserManageGroupID not implemented")
}
func (UnimplementedGroupRelationServiceServer) RemoveFromGroup(context.Context, *RemoveFromGroupRequest) (*RemoveFromGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromGroup not implemented")
}
func (UnimplementedGroupRelationServiceServer) LeaveGroup(context.Context, *LeaveGroupRequest) (*LeaveGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveGroup not implemented")
}
func (UnimplementedGroupRelationServiceServer) LeaveGroupRevert(context.Context, *LeaveGroupRequest) (*LeaveGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveGroupRevert not implemented")
}
func (UnimplementedGroupRelationServiceServer) GetGroupRelation(context.Context, *GetGroupRelationRequest) (*GetGroupRelationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupRelation not implemented")
}
func (UnimplementedGroupRelationServiceServer) GetBatchGroupRelation(context.Context, *GetBatchGroupRelationRequest) (*GetBatchGroupRelationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBatchGroupRelation not implemented")
}
func (UnimplementedGroupRelationServiceServer) DeleteGroupRelationByGroupId(context.Context, *GroupIDRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroupRelationByGroupId not implemented")
}
func (UnimplementedGroupRelationServiceServer) DeleteGroupRelationByGroupIdRevert(context.Context, *GroupIDRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroupRelationByGroupIdRevert not implemented")
}
func (UnimplementedGroupRelationServiceServer) DeleteGroupRelationByGroupIdAndUserID(context.Context, *DeleteGroupRelationByGroupIdAndUserIDRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroupRelationByGroupIdAndUserID not implemented")
}
func (UnimplementedGroupRelationServiceServer) DeleteGroupRelationByGroupIdAndUserIDRevert(context.Context, *DeleteGroupRelationByGroupIdAndUserIDRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroupRelationByGroupIdAndUserIDRevert not implemented")
}
func (UnimplementedGroupRelationServiceServer) CreateGroupAndInviteUsers(context.Context, *CreateGroupAndInviteUsersRequest) (*CreateGroupAndInviteUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroupAndInviteUsers not implemented")
}
func (UnimplementedGroupRelationServiceServer) CreateGroupAndInviteUsersRevert(context.Context, *CreateGroupAndInviteUsersRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroupAndInviteUsersRevert not implemented")
}
func (UnimplementedGroupRelationServiceServer) SetGroupSilentNotification(context.Context, *SetGroupSilentNotificationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGroupSilentNotification not implemented")
}
func (UnimplementedGroupRelationServiceServer) RemoveGroupRelationByGroupIdAndUserIDs(context.Context, *RemoveGroupRelationByGroupIdAndUserIDsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveGroupRelationByGroupIdAndUserIDs not implemented")
}
func (UnimplementedGroupRelationServiceServer) SetGroupOpenBurnAfterReading(context.Context, *SetGroupOpenBurnAfterReadingRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGroupOpenBurnAfterReading not implemented")
}
func (UnimplementedGroupRelationServiceServer) mustEmbedUnimplementedGroupRelationServiceServer() {}

// UnsafeGroupRelationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupRelationServiceServer will
// result in compilation errors.
type UnsafeGroupRelationServiceServer interface {
	mustEmbedUnimplementedGroupRelationServiceServer()
}

func RegisterGroupRelationServiceServer(s grpc.ServiceRegistrar, srv GroupRelationServiceServer) {
	s.RegisterService(&GroupRelationService_ServiceDesc, srv)
}

func _GroupRelationService_GetGroupUserIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupRelationServiceServer).GetGroupUserIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupRelationService_GetGroupUserIDs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupRelationServiceServer).GetGroupUserIDs(ctx, req.(*GroupIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupRelationService_GetUserGroupIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserGroupIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupRelationServiceServer).GetUserGroupIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupRelationService_GetUserGroupIDs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupRelationServiceServer).GetUserGroupIDs(ctx, req.(*GetUserGroupIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupRelationService_GetUserGroupList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserGroupListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupRelationServiceServer).GetUserGroupList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupRelationService_GetUserGroupList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupRelationServiceServer).GetUserGroupList(ctx, req.(*GetUserGroupListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupRelationService_GetGroupAdminIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupRelationServiceServer).GetGroupAdminIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupRelationService_GetGroupAdminIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupRelationServiceServer).GetGroupAdminIds(ctx, req.(*GroupIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupRelationService_GetUserManageGroupID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserManageGroupIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupRelationServiceServer).GetUserManageGroupID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupRelationService_GetUserManageGroupID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupRelationServiceServer).GetUserManageGroupID(ctx, req.(*GetUserManageGroupIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupRelationService_RemoveFromGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFromGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupRelationServiceServer).RemoveFromGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupRelationService_RemoveFromGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupRelationServiceServer).RemoveFromGroup(ctx, req.(*RemoveFromGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupRelationService_LeaveGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupRelationServiceServer).LeaveGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupRelationService_LeaveGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupRelationServiceServer).LeaveGroup(ctx, req.(*LeaveGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupRelationService_LeaveGroupRevert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupRelationServiceServer).LeaveGroupRevert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupRelationService_LeaveGroupRevert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupRelationServiceServer).LeaveGroupRevert(ctx, req.(*LeaveGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupRelationService_GetGroupRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupRelationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupRelationServiceServer).GetGroupRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupRelationService_GetGroupRelation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupRelationServiceServer).GetGroupRelation(ctx, req.(*GetGroupRelationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupRelationService_GetBatchGroupRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBatchGroupRelationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupRelationServiceServer).GetBatchGroupRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupRelationService_GetBatchGroupRelation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupRelationServiceServer).GetBatchGroupRelation(ctx, req.(*GetBatchGroupRelationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupRelationService_DeleteGroupRelationByGroupId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupRelationServiceServer).DeleteGroupRelationByGroupId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupRelationService_DeleteGroupRelationByGroupId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupRelationServiceServer).DeleteGroupRelationByGroupId(ctx, req.(*GroupIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupRelationService_DeleteGroupRelationByGroupIdRevert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupRelationServiceServer).DeleteGroupRelationByGroupIdRevert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupRelationService_DeleteGroupRelationByGroupIdRevert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupRelationServiceServer).DeleteGroupRelationByGroupIdRevert(ctx, req.(*GroupIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupRelationService_DeleteGroupRelationByGroupIdAndUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupRelationByGroupIdAndUserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupRelationServiceServer).DeleteGroupRelationByGroupIdAndUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupRelationService_DeleteGroupRelationByGroupIdAndUserID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupRelationServiceServer).DeleteGroupRelationByGroupIdAndUserID(ctx, req.(*DeleteGroupRelationByGroupIdAndUserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupRelationService_DeleteGroupRelationByGroupIdAndUserIDRevert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupRelationByGroupIdAndUserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupRelationServiceServer).DeleteGroupRelationByGroupIdAndUserIDRevert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupRelationService_DeleteGroupRelationByGroupIdAndUserIDRevert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupRelationServiceServer).DeleteGroupRelationByGroupIdAndUserIDRevert(ctx, req.(*DeleteGroupRelationByGroupIdAndUserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupRelationService_CreateGroupAndInviteUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupAndInviteUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupRelationServiceServer).CreateGroupAndInviteUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupRelationService_CreateGroupAndInviteUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupRelationServiceServer).CreateGroupAndInviteUsers(ctx, req.(*CreateGroupAndInviteUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupRelationService_CreateGroupAndInviteUsersRevert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupAndInviteUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupRelationServiceServer).CreateGroupAndInviteUsersRevert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupRelationService_CreateGroupAndInviteUsersRevert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupRelationServiceServer).CreateGroupAndInviteUsersRevert(ctx, req.(*CreateGroupAndInviteUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupRelationService_SetGroupSilentNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetGroupSilentNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupRelationServiceServer).SetGroupSilentNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupRelationService_SetGroupSilentNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupRelationServiceServer).SetGroupSilentNotification(ctx, req.(*SetGroupSilentNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupRelationService_RemoveGroupRelationByGroupIdAndUserIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveGroupRelationByGroupIdAndUserIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupRelationServiceServer).RemoveGroupRelationByGroupIdAndUserIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupRelationService_RemoveGroupRelationByGroupIdAndUserIDs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupRelationServiceServer).RemoveGroupRelationByGroupIdAndUserIDs(ctx, req.(*RemoveGroupRelationByGroupIdAndUserIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupRelationService_SetGroupOpenBurnAfterReading_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetGroupOpenBurnAfterReadingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupRelationServiceServer).SetGroupOpenBurnAfterReading(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupRelationService_SetGroupOpenBurnAfterReading_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupRelationServiceServer).SetGroupOpenBurnAfterReading(ctx, req.(*SetGroupOpenBurnAfterReadingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GroupRelationService_ServiceDesc is the grpc.ServiceDesc for GroupRelationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GroupRelationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.GroupRelationService",
	HandlerType: (*GroupRelationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGroupUserIDs",
			Handler:    _GroupRelationService_GetGroupUserIDs_Handler,
		},
		{
			MethodName: "GetUserGroupIDs",
			Handler:    _GroupRelationService_GetUserGroupIDs_Handler,
		},
		{
			MethodName: "GetUserGroupList",
			Handler:    _GroupRelationService_GetUserGroupList_Handler,
		},
		{
			MethodName: "GetGroupAdminIds",
			Handler:    _GroupRelationService_GetGroupAdminIds_Handler,
		},
		{
			MethodName: "GetUserManageGroupID",
			Handler:    _GroupRelationService_GetUserManageGroupID_Handler,
		},
		{
			MethodName: "RemoveFromGroup",
			Handler:    _GroupRelationService_RemoveFromGroup_Handler,
		},
		{
			MethodName: "LeaveGroup",
			Handler:    _GroupRelationService_LeaveGroup_Handler,
		},
		{
			MethodName: "LeaveGroupRevert",
			Handler:    _GroupRelationService_LeaveGroupRevert_Handler,
		},
		{
			MethodName: "GetGroupRelation",
			Handler:    _GroupRelationService_GetGroupRelation_Handler,
		},
		{
			MethodName: "GetBatchGroupRelation",
			Handler:    _GroupRelationService_GetBatchGroupRelation_Handler,
		},
		{
			MethodName: "DeleteGroupRelationByGroupId",
			Handler:    _GroupRelationService_DeleteGroupRelationByGroupId_Handler,
		},
		{
			MethodName: "DeleteGroupRelationByGroupIdRevert",
			Handler:    _GroupRelationService_DeleteGroupRelationByGroupIdRevert_Handler,
		},
		{
			MethodName: "DeleteGroupRelationByGroupIdAndUserID",
			Handler:    _GroupRelationService_DeleteGroupRelationByGroupIdAndUserID_Handler,
		},
		{
			MethodName: "DeleteGroupRelationByGroupIdAndUserIDRevert",
			Handler:    _GroupRelationService_DeleteGroupRelationByGroupIdAndUserIDRevert_Handler,
		},
		{
			MethodName: "CreateGroupAndInviteUsers",
			Handler:    _GroupRelationService_CreateGroupAndInviteUsers_Handler,
		},
		{
			MethodName: "CreateGroupAndInviteUsersRevert",
			Handler:    _GroupRelationService_CreateGroupAndInviteUsersRevert_Handler,
		},
		{
			MethodName: "SetGroupSilentNotification",
			Handler:    _GroupRelationService_SetGroupSilentNotification_Handler,
		},
		{
			MethodName: "RemoveGroupRelationByGroupIdAndUserIDs",
			Handler:    _GroupRelationService_RemoveGroupRelationByGroupIdAndUserIDs_Handler,
		},
		{
			MethodName: "SetGroupOpenBurnAfterReading",
			Handler:    _GroupRelationService_SetGroupOpenBurnAfterReading_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/group_relation.proto",
}
