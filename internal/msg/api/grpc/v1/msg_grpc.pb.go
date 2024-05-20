// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: api/grpc/v1/msg.proto

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
	MsgService_SendUserMessage_FullMethodName                    = "/msg_v1.MsgService/SendUserMessage"
	MsgService_SendMultiUserMessage_FullMethodName               = "/msg_v1.MsgService/SendMultiUserMessage"
	MsgService_ConfirmDeleteUserMessageByDialogId_FullMethodName = "/msg_v1.MsgService/ConfirmDeleteUserMessageByDialogId"
	MsgService_DeleteUserMessageById_FullMethodName              = "/msg_v1.MsgService/DeleteUserMessageById"
	MsgService_DeleteUserMessageByIDs_FullMethodName             = "/msg_v1.MsgService/DeleteUserMessageByIDs"
)

// MsgServiceClient is the client API for MsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgServiceClient interface {
	// 发送私聊消息
	SendUserMessage(ctx context.Context, in *SendUserMsgRequest, opts ...grpc.CallOption) (*SendUserMsgResponse, error)
	// //发送私聊消息
	// rpc SendUserMessageRevert(MsgIdRequest)returns (SendUserMsgRevertResponse);
	// //群发私聊消息
	SendMultiUserMessage(ctx context.Context, in *SendMultiUserMsgRequest, opts ...grpc.CallOption) (*SendMultiUserMsgResponse, error)
	// //发送群消息
	// rpc SendGroupMessage(SendGroupMsgRequest) returns(SendGroupMsgResponse);
	// //发送群消息回滚
	// rpc SendGroupMessageRevert(MsgIdRequest)returns (SendGroupMsgRevertResponse);
	// //获取用户消息列表
	// rpc GetUserMessageList(GetUserMsgListRequest) returns(GetUserMsgListResponse);
	// //获取私聊对话最新消息
	// rpc GetUserLastMessageList(GetLastMsgListRequest)returns(UserMessages);
	// //获取群聊对话最新消息
	// rpc GetGroupLastMessageList(GetLastMsgListRequest) returns(GroupMessages);
	//
	// //根据好友id获取最后一条消息
	// rpc GetLastMsgsForUserWithFriends(UserMsgsRequest) returns (UserMessages);
	// //根据群组id获取最后一条消息
	// rpc GetLastMsgsForGroupsWithIDs(GroupMsgsRequest) returns (GroupMessages);
	// //根据对话id获取最后一条消息
	// rpc GetLastUserMsgsByDialogIds(GetLastMsgsByDialogIdsRequest) returns (GetLastMsgsResponse);
	// //编辑私聊消息
	// rpc EditUserMessage(EditUserMsgRequest) returns (UserMessage);
	// //撤回私聊消息
	// rpc DeleteUserMessage(DeleteUserMsgRequest) returns (UserMessage);
	// //根据对话id与msgId查询msgId之后的私聊消息
	// rpc GetUserMsgIdAfterMsgList(GetUserMsgIdAfterMsgListRequest) returns (GetUserMsgIdAfterMsgListResponse);
	//
	// //获取群消息列表
	// rpc GetGroupMessageList(GetGroupMsgListRequest) returns(GetGroupMsgListResponse);
	// //编辑群消息
	// rpc EditGroupMessage(EditGroupMsgRequest) returns (GroupMessage);
	// //撤回群消息
	// rpc DeleteGroupMessage(DeleteGroupMsgRequest) returns (GroupMessage);
	// //根据消息id获取私聊消息
	// rpc GetUserMessageById(GetUserMsgByIDRequest) returns (UserMessage);
	// //根据消息id获取群消息
	// rpc GetGroupMessageById(GetGroupMsgByIDRequest) returns (GroupMessage);
	// //根据多个消息id获取私聊消息
	// rpc GetUserMessagesByIds(GetUserMessagesByIdsRequest) returns (GetUserMessagesByIdsResponse);
	// //根据多个消息id获取群消息
	// rpc GetGroupMessagesByIds(GetGroupMessagesByIdsRequest) returns (GetGroupMessagesByIdsResponse);
	// //设置私聊消息标注状态
	// rpc SetUserMsgLabel(SetUserMsgLabelRequest) returns (SetUserMsgLabelResponse);
	// // 设置群消息标注状态
	// rpc SetGroupMsgLabel(SetGroupMsgLabelRequest) returns (SetGroupMsgLabelResponse);
	// //根据对话id获取私聊标注信息
	// rpc GetUserMsgLabelByDialogId(GetUserMsgLabelByDialogIdRequest) returns (GetUserMsgLabelByDialogIdResponse);
	// //根据对话id获取群消息标注信息
	// rpc GetGroupMsgLabelByDialogId(GetGroupMsgLabelByDialogIdRequest) returns (GetGroupMsgLabelByDialogIdResponse);
	// //根据对话id与msgId查询msgId之后的群消息
	// rpc GetGroupMsgIdAfterMsgList(GetGroupMsgIdAfterMsgListRequest) returns (GetGroupMsgIdAfterMsgListResponse);
	// //获取群聊未读消息
	// rpc GetGroupUnreadMessages(GetGroupUnreadMessagesRequest) returns (GetGroupUnreadMessagesResponse);
	//
	// // 一键已读用户消息
	// rpc ReadAllUserMsg(ReadAllUserMsgRequest) returns (ReadAllUserMsgResponse);
	//
	// // 一键已读群聊消息
	// rpc ReadAllGroupMsg(ReadAllGroupMsgRequest) returns (ReadAllGroupMsgResponse);
	//
	// //批量设置私聊消息id为已读
	// rpc SetUserMsgsReadStatus(SetUserMsgsReadStatusRequest) returns (SetUserMsgsReadStatusResponse);
	// //修改指定私聊消息的已读状态
	// rpc SetUserMsgReadStatus(SetUserMsgReadStatusRequest) returns (SetUserMsgReadStatusResponse);
	// //获取私聊对话未读消息
	// rpc GetUnreadUserMsgs(GetUnreadUserMsgsRequest) returns (GetUnreadUserMsgsResponse);
	// //确认根据对话id删除私聊消息
	ConfirmDeleteUserMessageByDialogId(ctx context.Context, in *DeleteUserMsgByDialogIdRequest, opts ...grpc.CallOption) (*DeleteUserMsgByDialogIdResponse, error)
	// //确认根据对话id删除群聊消息
	// rpc ConfirmDeleteGroupMessageByDialogId(DeleteGroupMsgByDialogIdRequest) returns (DeleteGroupMsgByDialogIdResponse);
	// //根据对话id删除私聊消息
	// rpc DeleteUserMessageByDialogId(DeleteUserMsgByDialogIdRequest) returns (DeleteUserMsgByDialogIdResponse);
	// //根据对话id删除群聊消息
	// rpc DeleteGroupMessageByDialogId(DeleteGroupMsgByDialogIdRequest) returns (DeleteGroupMsgByDialogIdResponse);
	// //根据对话id删除私聊消息回滚
	// rpc DeleteUserMessageByDialogIdRollback(DeleteUserMsgByDialogIdRequest) returns (DeleteUserMsgByDialogIdResponse);
	// //根据对话id删除群聊消息回滚
	// rpc DeleteGroupMessageByDialogIdRollback(DeleteUserMsgByDialogIdRequest) returns (DeleteGroupMsgByDialogIdResponse);
	// //根据消息id删除私聊消息
	DeleteUserMessageById(ctx context.Context, in *DeleteUserMsgByIDRequest, opts ...grpc.CallOption) (*DeleteUserMsgByIDResponse, error)
	// 根据消息ids删除私聊消息
	DeleteUserMessageByIDs(ctx context.Context, in *DeleteUserMessageByIdsRequest, opts ...grpc.CallOption) (*DeleteUserMsgByIDResponse, error)
}

type msgServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgServiceClient(cc grpc.ClientConnInterface) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) SendUserMessage(ctx context.Context, in *SendUserMsgRequest, opts ...grpc.CallOption) (*SendUserMsgResponse, error) {
	out := new(SendUserMsgResponse)
	err := c.cc.Invoke(ctx, MsgService_SendUserMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) SendMultiUserMessage(ctx context.Context, in *SendMultiUserMsgRequest, opts ...grpc.CallOption) (*SendMultiUserMsgResponse, error) {
	out := new(SendMultiUserMsgResponse)
	err := c.cc.Invoke(ctx, MsgService_SendMultiUserMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) ConfirmDeleteUserMessageByDialogId(ctx context.Context, in *DeleteUserMsgByDialogIdRequest, opts ...grpc.CallOption) (*DeleteUserMsgByDialogIdResponse, error) {
	out := new(DeleteUserMsgByDialogIdResponse)
	err := c.cc.Invoke(ctx, MsgService_ConfirmDeleteUserMessageByDialogId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) DeleteUserMessageById(ctx context.Context, in *DeleteUserMsgByIDRequest, opts ...grpc.CallOption) (*DeleteUserMsgByIDResponse, error) {
	out := new(DeleteUserMsgByIDResponse)
	err := c.cc.Invoke(ctx, MsgService_DeleteUserMessageById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) DeleteUserMessageByIDs(ctx context.Context, in *DeleteUserMessageByIdsRequest, opts ...grpc.CallOption) (*DeleteUserMsgByIDResponse, error) {
	out := new(DeleteUserMsgByIDResponse)
	err := c.cc.Invoke(ctx, MsgService_DeleteUserMessageByIDs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
// All implementations should embed UnimplementedMsgServiceServer
// for forward compatibility
type MsgServiceServer interface {
	// 发送私聊消息
	SendUserMessage(context.Context, *SendUserMsgRequest) (*SendUserMsgResponse, error)
	// //发送私聊消息
	// rpc SendUserMessageRevert(MsgIdRequest)returns (SendUserMsgRevertResponse);
	// //群发私聊消息
	SendMultiUserMessage(context.Context, *SendMultiUserMsgRequest) (*SendMultiUserMsgResponse, error)
	// //发送群消息
	// rpc SendGroupMessage(SendGroupMsgRequest) returns(SendGroupMsgResponse);
	// //发送群消息回滚
	// rpc SendGroupMessageRevert(MsgIdRequest)returns (SendGroupMsgRevertResponse);
	// //获取用户消息列表
	// rpc GetUserMessageList(GetUserMsgListRequest) returns(GetUserMsgListResponse);
	// //获取私聊对话最新消息
	// rpc GetUserLastMessageList(GetLastMsgListRequest)returns(UserMessages);
	// //获取群聊对话最新消息
	// rpc GetGroupLastMessageList(GetLastMsgListRequest) returns(GroupMessages);
	//
	// //根据好友id获取最后一条消息
	// rpc GetLastMsgsForUserWithFriends(UserMsgsRequest) returns (UserMessages);
	// //根据群组id获取最后一条消息
	// rpc GetLastMsgsForGroupsWithIDs(GroupMsgsRequest) returns (GroupMessages);
	// //根据对话id获取最后一条消息
	// rpc GetLastUserMsgsByDialogIds(GetLastMsgsByDialogIdsRequest) returns (GetLastMsgsResponse);
	// //编辑私聊消息
	// rpc EditUserMessage(EditUserMsgRequest) returns (UserMessage);
	// //撤回私聊消息
	// rpc DeleteUserMessage(DeleteUserMsgRequest) returns (UserMessage);
	// //根据对话id与msgId查询msgId之后的私聊消息
	// rpc GetUserMsgIdAfterMsgList(GetUserMsgIdAfterMsgListRequest) returns (GetUserMsgIdAfterMsgListResponse);
	//
	// //获取群消息列表
	// rpc GetGroupMessageList(GetGroupMsgListRequest) returns(GetGroupMsgListResponse);
	// //编辑群消息
	// rpc EditGroupMessage(EditGroupMsgRequest) returns (GroupMessage);
	// //撤回群消息
	// rpc DeleteGroupMessage(DeleteGroupMsgRequest) returns (GroupMessage);
	// //根据消息id获取私聊消息
	// rpc GetUserMessageById(GetUserMsgByIDRequest) returns (UserMessage);
	// //根据消息id获取群消息
	// rpc GetGroupMessageById(GetGroupMsgByIDRequest) returns (GroupMessage);
	// //根据多个消息id获取私聊消息
	// rpc GetUserMessagesByIds(GetUserMessagesByIdsRequest) returns (GetUserMessagesByIdsResponse);
	// //根据多个消息id获取群消息
	// rpc GetGroupMessagesByIds(GetGroupMessagesByIdsRequest) returns (GetGroupMessagesByIdsResponse);
	// //设置私聊消息标注状态
	// rpc SetUserMsgLabel(SetUserMsgLabelRequest) returns (SetUserMsgLabelResponse);
	// // 设置群消息标注状态
	// rpc SetGroupMsgLabel(SetGroupMsgLabelRequest) returns (SetGroupMsgLabelResponse);
	// //根据对话id获取私聊标注信息
	// rpc GetUserMsgLabelByDialogId(GetUserMsgLabelByDialogIdRequest) returns (GetUserMsgLabelByDialogIdResponse);
	// //根据对话id获取群消息标注信息
	// rpc GetGroupMsgLabelByDialogId(GetGroupMsgLabelByDialogIdRequest) returns (GetGroupMsgLabelByDialogIdResponse);
	// //根据对话id与msgId查询msgId之后的群消息
	// rpc GetGroupMsgIdAfterMsgList(GetGroupMsgIdAfterMsgListRequest) returns (GetGroupMsgIdAfterMsgListResponse);
	// //获取群聊未读消息
	// rpc GetGroupUnreadMessages(GetGroupUnreadMessagesRequest) returns (GetGroupUnreadMessagesResponse);
	//
	// // 一键已读用户消息
	// rpc ReadAllUserMsg(ReadAllUserMsgRequest) returns (ReadAllUserMsgResponse);
	//
	// // 一键已读群聊消息
	// rpc ReadAllGroupMsg(ReadAllGroupMsgRequest) returns (ReadAllGroupMsgResponse);
	//
	// //批量设置私聊消息id为已读
	// rpc SetUserMsgsReadStatus(SetUserMsgsReadStatusRequest) returns (SetUserMsgsReadStatusResponse);
	// //修改指定私聊消息的已读状态
	// rpc SetUserMsgReadStatus(SetUserMsgReadStatusRequest) returns (SetUserMsgReadStatusResponse);
	// //获取私聊对话未读消息
	// rpc GetUnreadUserMsgs(GetUnreadUserMsgsRequest) returns (GetUnreadUserMsgsResponse);
	// //确认根据对话id删除私聊消息
	ConfirmDeleteUserMessageByDialogId(context.Context, *DeleteUserMsgByDialogIdRequest) (*DeleteUserMsgByDialogIdResponse, error)
	// //确认根据对话id删除群聊消息
	// rpc ConfirmDeleteGroupMessageByDialogId(DeleteGroupMsgByDialogIdRequest) returns (DeleteGroupMsgByDialogIdResponse);
	// //根据对话id删除私聊消息
	// rpc DeleteUserMessageByDialogId(DeleteUserMsgByDialogIdRequest) returns (DeleteUserMsgByDialogIdResponse);
	// //根据对话id删除群聊消息
	// rpc DeleteGroupMessageByDialogId(DeleteGroupMsgByDialogIdRequest) returns (DeleteGroupMsgByDialogIdResponse);
	// //根据对话id删除私聊消息回滚
	// rpc DeleteUserMessageByDialogIdRollback(DeleteUserMsgByDialogIdRequest) returns (DeleteUserMsgByDialogIdResponse);
	// //根据对话id删除群聊消息回滚
	// rpc DeleteGroupMessageByDialogIdRollback(DeleteUserMsgByDialogIdRequest) returns (DeleteGroupMsgByDialogIdResponse);
	// //根据消息id删除私聊消息
	DeleteUserMessageById(context.Context, *DeleteUserMsgByIDRequest) (*DeleteUserMsgByIDResponse, error)
	// 根据消息ids删除私聊消息
	DeleteUserMessageByIDs(context.Context, *DeleteUserMessageByIdsRequest) (*DeleteUserMsgByIDResponse, error)
}

// UnimplementedMsgServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (UnimplementedMsgServiceServer) SendUserMessage(context.Context, *SendUserMsgRequest) (*SendUserMsgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendUserMessage not implemented")
}
func (UnimplementedMsgServiceServer) SendMultiUserMessage(context.Context, *SendMultiUserMsgRequest) (*SendMultiUserMsgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMultiUserMessage not implemented")
}
func (UnimplementedMsgServiceServer) ConfirmDeleteUserMessageByDialogId(context.Context, *DeleteUserMsgByDialogIdRequest) (*DeleteUserMsgByDialogIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmDeleteUserMessageByDialogId not implemented")
}
func (UnimplementedMsgServiceServer) DeleteUserMessageById(context.Context, *DeleteUserMsgByIDRequest) (*DeleteUserMsgByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserMessageById not implemented")
}
func (UnimplementedMsgServiceServer) DeleteUserMessageByIDs(context.Context, *DeleteUserMessageByIdsRequest) (*DeleteUserMsgByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserMessageByIDs not implemented")
}

// UnsafeMsgServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServiceServer will
// result in compilation errors.
type UnsafeMsgServiceServer interface {
	mustEmbedUnimplementedMsgServiceServer()
}

func RegisterMsgServiceServer(s grpc.ServiceRegistrar, srv MsgServiceServer) {
	s.RegisterService(&MsgService_ServiceDesc, srv)
}

func _MsgService_SendUserMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendUserMsgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).SendUserMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MsgService_SendUserMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).SendUserMessage(ctx, req.(*SendUserMsgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_SendMultiUserMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMultiUserMsgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).SendMultiUserMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MsgService_SendMultiUserMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).SendMultiUserMessage(ctx, req.(*SendMultiUserMsgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_ConfirmDeleteUserMessageByDialogId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserMsgByDialogIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).ConfirmDeleteUserMessageByDialogId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MsgService_ConfirmDeleteUserMessageByDialogId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).ConfirmDeleteUserMessageByDialogId(ctx, req.(*DeleteUserMsgByDialogIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_DeleteUserMessageById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserMsgByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).DeleteUserMessageById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MsgService_DeleteUserMessageById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).DeleteUserMessageById(ctx, req.(*DeleteUserMsgByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_DeleteUserMessageByIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserMessageByIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).DeleteUserMessageByIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MsgService_DeleteUserMessageByIDs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).DeleteUserMessageByIDs(ctx, req.(*DeleteUserMessageByIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MsgService_ServiceDesc is the grpc.ServiceDesc for MsgService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MsgService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "msg_v1.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendUserMessage",
			Handler:    _MsgService_SendUserMessage_Handler,
		},
		{
			MethodName: "SendMultiUserMessage",
			Handler:    _MsgService_SendMultiUserMessage_Handler,
		},
		{
			MethodName: "ConfirmDeleteUserMessageByDialogId",
			Handler:    _MsgService_ConfirmDeleteUserMessageByDialogId_Handler,
		},
		{
			MethodName: "DeleteUserMessageById",
			Handler:    _MsgService_DeleteUserMessageById_Handler,
		},
		{
			MethodName: "DeleteUserMessageByIDs",
			Handler:    _MsgService_DeleteUserMessageByIDs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/grpc/v1/msg.proto",
}
