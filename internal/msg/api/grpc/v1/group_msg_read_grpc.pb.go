// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: api/grpc/v1/group_msg_read.proto

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
	GroupMessageService_SetGroupMessageRead_FullMethodName                 = "/msg_v1.GroupMessageService/SetGroupMessageRead"
	GroupMessageService_GetGroupMessageReaders_FullMethodName              = "/msg_v1.GroupMessageService/GetGroupMessageReaders"
	GroupMessageService_GetGroupMessageReadByMsgIdAndUserId_FullMethodName = "/msg_v1.GroupMessageService/GetGroupMessageReadByMsgIdAndUserId"
)

// GroupMessageServiceClient is the client API for GroupMessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupMessageServiceClient interface {
	// 设置群聊消息已读
	SetGroupMessageRead(ctx context.Context, in *SetGroupMessagesReadRequest, opts ...grpc.CallOption) (*SetGroupMessageReadResponse, error)
	// 获取指定消息已读人员
	GetGroupMessageReaders(ctx context.Context, in *GetGroupMessageReadersRequest, opts ...grpc.CallOption) (*GetGroupMessageReadersResponse, error)
	// 根据消息id与用户id查询是否已经读取过
	GetGroupMessageReadByMsgIdAndUserId(ctx context.Context, in *GetGroupMessageReadByMsgIdAndUserIdRequest, opts ...grpc.CallOption) (*GetGroupMessageReadByMsgIdAndUserIdResponse, error)
}

type groupMessageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupMessageServiceClient(cc grpc.ClientConnInterface) GroupMessageServiceClient {
	return &groupMessageServiceClient{cc}
}

func (c *groupMessageServiceClient) SetGroupMessageRead(ctx context.Context, in *SetGroupMessagesReadRequest, opts ...grpc.CallOption) (*SetGroupMessageReadResponse, error) {
	out := new(SetGroupMessageReadResponse)
	err := c.cc.Invoke(ctx, GroupMessageService_SetGroupMessageRead_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupMessageServiceClient) GetGroupMessageReaders(ctx context.Context, in *GetGroupMessageReadersRequest, opts ...grpc.CallOption) (*GetGroupMessageReadersResponse, error) {
	out := new(GetGroupMessageReadersResponse)
	err := c.cc.Invoke(ctx, GroupMessageService_GetGroupMessageReaders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupMessageServiceClient) GetGroupMessageReadByMsgIdAndUserId(ctx context.Context, in *GetGroupMessageReadByMsgIdAndUserIdRequest, opts ...grpc.CallOption) (*GetGroupMessageReadByMsgIdAndUserIdResponse, error) {
	out := new(GetGroupMessageReadByMsgIdAndUserIdResponse)
	err := c.cc.Invoke(ctx, GroupMessageService_GetGroupMessageReadByMsgIdAndUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupMessageServiceServer is the server API for GroupMessageService service.
// All implementations should embed UnimplementedGroupMessageServiceServer
// for forward compatibility
type GroupMessageServiceServer interface {
	// 设置群聊消息已读
	SetGroupMessageRead(context.Context, *SetGroupMessagesReadRequest) (*SetGroupMessageReadResponse, error)
	// 获取指定消息已读人员
	GetGroupMessageReaders(context.Context, *GetGroupMessageReadersRequest) (*GetGroupMessageReadersResponse, error)
	// 根据消息id与用户id查询是否已经读取过
	GetGroupMessageReadByMsgIdAndUserId(context.Context, *GetGroupMessageReadByMsgIdAndUserIdRequest) (*GetGroupMessageReadByMsgIdAndUserIdResponse, error)
}

// UnimplementedGroupMessageServiceServer should be embedded to have forward compatible implementations.
type UnimplementedGroupMessageServiceServer struct {
}

func (UnimplementedGroupMessageServiceServer) SetGroupMessageRead(context.Context, *SetGroupMessagesReadRequest) (*SetGroupMessageReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGroupMessageRead not implemented")
}
func (UnimplementedGroupMessageServiceServer) GetGroupMessageReaders(context.Context, *GetGroupMessageReadersRequest) (*GetGroupMessageReadersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupMessageReaders not implemented")
}
func (UnimplementedGroupMessageServiceServer) GetGroupMessageReadByMsgIdAndUserId(context.Context, *GetGroupMessageReadByMsgIdAndUserIdRequest) (*GetGroupMessageReadByMsgIdAndUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupMessageReadByMsgIdAndUserId not implemented")
}

// UnsafeGroupMessageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupMessageServiceServer will
// result in compilation errors.
type UnsafeGroupMessageServiceServer interface {
	mustEmbedUnimplementedGroupMessageServiceServer()
}

func RegisterGroupMessageServiceServer(s grpc.ServiceRegistrar, srv GroupMessageServiceServer) {
	s.RegisterService(&GroupMessageService_ServiceDesc, srv)
}

func _GroupMessageService_SetGroupMessageRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetGroupMessagesReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupMessageServiceServer).SetGroupMessageRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupMessageService_SetGroupMessageRead_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupMessageServiceServer).SetGroupMessageRead(ctx, req.(*SetGroupMessagesReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupMessageService_GetGroupMessageReaders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupMessageReadersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupMessageServiceServer).GetGroupMessageReaders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupMessageService_GetGroupMessageReaders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupMessageServiceServer).GetGroupMessageReaders(ctx, req.(*GetGroupMessageReadersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupMessageService_GetGroupMessageReadByMsgIdAndUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupMessageReadByMsgIdAndUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupMessageServiceServer).GetGroupMessageReadByMsgIdAndUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupMessageService_GetGroupMessageReadByMsgIdAndUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupMessageServiceServer).GetGroupMessageReadByMsgIdAndUserId(ctx, req.(*GetGroupMessageReadByMsgIdAndUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GroupMessageService_ServiceDesc is the grpc.ServiceDesc for GroupMessageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GroupMessageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "msg_v1.GroupMessageService",
	HandlerType: (*GroupMessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetGroupMessageRead",
			Handler:    _GroupMessageService_SetGroupMessageRead_Handler,
		},
		{
			MethodName: "GetGroupMessageReaders",
			Handler:    _GroupMessageService_GetGroupMessageReaders_Handler,
		},
		{
			MethodName: "GetGroupMessageReadByMsgIdAndUserId",
			Handler:    _GroupMessageService_GetGroupMessageReadByMsgIdAndUserId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/grpc/v1/group_msg_read.proto",
}
