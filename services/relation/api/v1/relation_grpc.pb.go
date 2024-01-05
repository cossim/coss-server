// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: api/v1/relation.proto

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
	RelationService_AddFriend_FullMethodName       = "/v1.RelationService/AddFriend"
	RelationService_ConfirmFriend_FullMethodName   = "/v1.RelationService/ConfirmFriend"
	RelationService_DeleteFriend_FullMethodName    = "/v1.RelationService/DeleteFriend"
	RelationService_AddBlacklist_FullMethodName    = "/v1.RelationService/AddBlacklist"
	RelationService_DeleteBlacklist_FullMethodName = "/v1.RelationService/DeleteBlacklist"
	RelationService_GetFriendList_FullMethodName   = "/v1.RelationService/GetFriendList"
	RelationService_GetBlacklist_FullMethodName    = "/v1.RelationService/GetBlacklist"
	RelationService_GetUserRelation_FullMethodName = "/v1.RelationService/GetUserRelation"
)

// RelationServiceClient is the client API for RelationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RelationServiceClient interface {
	AddFriend(ctx context.Context, in *AddFriendRequest, opts ...grpc.CallOption) (*AddFriendResponse, error)
	ConfirmFriend(ctx context.Context, in *ConfirmFriendRequest, opts ...grpc.CallOption) (*ConfirmFriendResponse, error)
	DeleteFriend(ctx context.Context, in *DeleteFriendRequest, opts ...grpc.CallOption) (*DeleteFriendResponse, error)
	AddBlacklist(ctx context.Context, in *AddBlacklistRequest, opts ...grpc.CallOption) (*AddBlacklistResponse, error)
	DeleteBlacklist(ctx context.Context, in *DeleteBlacklistRequest, opts ...grpc.CallOption) (*DeleteBlacklistResponse, error)
	GetFriendList(ctx context.Context, in *GetFriendListRequest, opts ...grpc.CallOption) (*GetFriendListResponse, error)
	GetBlacklist(ctx context.Context, in *GetBlacklistRequest, opts ...grpc.CallOption) (*GetBlacklistResponse, error)
	GetUserRelation(ctx context.Context, in *GetUserRelationRequest, opts ...grpc.CallOption) (*GetUserRelationResponse, error)
}

type relationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRelationServiceClient(cc grpc.ClientConnInterface) RelationServiceClient {
	return &relationServiceClient{cc}
}

func (c *relationServiceClient) AddFriend(ctx context.Context, in *AddFriendRequest, opts ...grpc.CallOption) (*AddFriendResponse, error) {
	out := new(AddFriendResponse)
	err := c.cc.Invoke(ctx, RelationService_AddFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) ConfirmFriend(ctx context.Context, in *ConfirmFriendRequest, opts ...grpc.CallOption) (*ConfirmFriendResponse, error) {
	out := new(ConfirmFriendResponse)
	err := c.cc.Invoke(ctx, RelationService_ConfirmFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) DeleteFriend(ctx context.Context, in *DeleteFriendRequest, opts ...grpc.CallOption) (*DeleteFriendResponse, error) {
	out := new(DeleteFriendResponse)
	err := c.cc.Invoke(ctx, RelationService_DeleteFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) AddBlacklist(ctx context.Context, in *AddBlacklistRequest, opts ...grpc.CallOption) (*AddBlacklistResponse, error) {
	out := new(AddBlacklistResponse)
	err := c.cc.Invoke(ctx, RelationService_AddBlacklist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) DeleteBlacklist(ctx context.Context, in *DeleteBlacklistRequest, opts ...grpc.CallOption) (*DeleteBlacklistResponse, error) {
	out := new(DeleteBlacklistResponse)
	err := c.cc.Invoke(ctx, RelationService_DeleteBlacklist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) GetFriendList(ctx context.Context, in *GetFriendListRequest, opts ...grpc.CallOption) (*GetFriendListResponse, error) {
	out := new(GetFriendListResponse)
	err := c.cc.Invoke(ctx, RelationService_GetFriendList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) GetBlacklist(ctx context.Context, in *GetBlacklistRequest, opts ...grpc.CallOption) (*GetBlacklistResponse, error) {
	out := new(GetBlacklistResponse)
	err := c.cc.Invoke(ctx, RelationService_GetBlacklist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) GetUserRelation(ctx context.Context, in *GetUserRelationRequest, opts ...grpc.CallOption) (*GetUserRelationResponse, error) {
	out := new(GetUserRelationResponse)
	err := c.cc.Invoke(ctx, RelationService_GetUserRelation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RelationServiceServer is the server API for RelationService service.
// All implementations must embed UnimplementedRelationServiceServer
// for forward compatibility
type RelationServiceServer interface {
	AddFriend(context.Context, *AddFriendRequest) (*AddFriendResponse, error)
	ConfirmFriend(context.Context, *ConfirmFriendRequest) (*ConfirmFriendResponse, error)
	DeleteFriend(context.Context, *DeleteFriendRequest) (*DeleteFriendResponse, error)
	AddBlacklist(context.Context, *AddBlacklistRequest) (*AddBlacklistResponse, error)
	DeleteBlacklist(context.Context, *DeleteBlacklistRequest) (*DeleteBlacklistResponse, error)
	GetFriendList(context.Context, *GetFriendListRequest) (*GetFriendListResponse, error)
	GetBlacklist(context.Context, *GetBlacklistRequest) (*GetBlacklistResponse, error)
	GetUserRelation(context.Context, *GetUserRelationRequest) (*GetUserRelationResponse, error)
	mustEmbedUnimplementedRelationServiceServer()
}

// UnimplementedRelationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRelationServiceServer struct {
}

func (UnimplementedRelationServiceServer) AddFriend(context.Context, *AddFriendRequest) (*AddFriendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFriend not implemented")
}
func (UnimplementedRelationServiceServer) ConfirmFriend(context.Context, *ConfirmFriendRequest) (*ConfirmFriendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmFriend not implemented")
}
func (UnimplementedRelationServiceServer) DeleteFriend(context.Context, *DeleteFriendRequest) (*DeleteFriendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFriend not implemented")
}
func (UnimplementedRelationServiceServer) AddBlacklist(context.Context, *AddBlacklistRequest) (*AddBlacklistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBlacklist not implemented")
}
func (UnimplementedRelationServiceServer) DeleteBlacklist(context.Context, *DeleteBlacklistRequest) (*DeleteBlacklistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBlacklist not implemented")
}
func (UnimplementedRelationServiceServer) GetFriendList(context.Context, *GetFriendListRequest) (*GetFriendListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriendList not implemented")
}
func (UnimplementedRelationServiceServer) GetBlacklist(context.Context, *GetBlacklistRequest) (*GetBlacklistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlacklist not implemented")
}
func (UnimplementedRelationServiceServer) GetUserRelation(context.Context, *GetUserRelationRequest) (*GetUserRelationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserRelation not implemented")
}
func (UnimplementedRelationServiceServer) mustEmbedUnimplementedRelationServiceServer() {}

// UnsafeRelationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RelationServiceServer will
// result in compilation errors.
type UnsafeRelationServiceServer interface {
	mustEmbedUnimplementedRelationServiceServer()
}

func RegisterRelationServiceServer(s grpc.ServiceRegistrar, srv RelationServiceServer) {
	s.RegisterService(&RelationService_ServiceDesc, srv)
}

func _RelationService_AddFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).AddFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_AddFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).AddFriend(ctx, req.(*AddFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_ConfirmFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).ConfirmFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_ConfirmFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).ConfirmFriend(ctx, req.(*ConfirmFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_DeleteFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).DeleteFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_DeleteFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).DeleteFriend(ctx, req.(*DeleteFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_AddBlacklist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBlacklistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).AddBlacklist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_AddBlacklist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).AddBlacklist(ctx, req.(*AddBlacklistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_DeleteBlacklist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBlacklistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).DeleteBlacklist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_DeleteBlacklist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).DeleteBlacklist(ctx, req.(*DeleteBlacklistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_GetFriendList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFriendListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).GetFriendList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_GetFriendList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).GetFriendList(ctx, req.(*GetFriendListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_GetBlacklist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlacklistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).GetBlacklist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_GetBlacklist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).GetBlacklist(ctx, req.(*GetBlacklistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_GetUserRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRelationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).GetUserRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_GetUserRelation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).GetUserRelation(ctx, req.(*GetUserRelationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RelationService_ServiceDesc is the grpc.ServiceDesc for RelationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RelationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.RelationService",
	HandlerType: (*RelationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFriend",
			Handler:    _RelationService_AddFriend_Handler,
		},
		{
			MethodName: "ConfirmFriend",
			Handler:    _RelationService_ConfirmFriend_Handler,
		},
		{
			MethodName: "DeleteFriend",
			Handler:    _RelationService_DeleteFriend_Handler,
		},
		{
			MethodName: "AddBlacklist",
			Handler:    _RelationService_AddBlacklist_Handler,
		},
		{
			MethodName: "DeleteBlacklist",
			Handler:    _RelationService_DeleteBlacklist_Handler,
		},
		{
			MethodName: "GetFriendList",
			Handler:    _RelationService_GetFriendList_Handler,
		},
		{
			MethodName: "GetBlacklist",
			Handler:    _RelationService_GetBlacklist_Handler,
		},
		{
			MethodName: "GetUserRelation",
			Handler:    _RelationService_GetUserRelation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/relation.proto",
}
