// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.15.8
// source: api/v1/group_relation.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	GroupId        uint32 `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	ShowSession    uint32 `protobuf:"varint,3,opt,name=show_session,json=showSession,proto3" json:"show_session,omitempty"`
	RelationStatus uint32 `protobuf:"varint,4,opt,name=relation_status,json=relationStatus,proto3" json:"relation_status,omitempty"`
}

func (x *UserGroupRequest) Reset() {
	*x = UserGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_group_relation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGroupRequest) ProtoMessage() {}

func (x *UserGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_group_relation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGroupRequest.ProtoReflect.Descriptor instead.
func (*UserGroupRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_group_relation_proto_rawDescGZIP(), []int{0}
}

func (x *UserGroupRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserGroupRequest) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *UserGroupRequest) GetShowSession() uint32 {
	if x != nil {
		return x.ShowSession
	}
	return 0
}

func (x *UserGroupRequest) GetRelationStatus() uint32 {
	if x != nil {
		return x.RelationStatus
	}
	return 0
}

type UserGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserGroupResponse) Reset() {
	*x = UserGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_group_relation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGroupResponse) ProtoMessage() {}

func (x *UserGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_group_relation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGroupResponse.ProtoReflect.Descriptor instead.
func (*UserGroupResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_group_relation_proto_rawDescGZIP(), []int{1}
}

type GroupID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"group_id"
	GroupId uint32 `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id"`
}

func (x *GroupID) Reset() {
	*x = GroupID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_group_relation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupID) ProtoMessage() {}

func (x *GroupID) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_group_relation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupID.ProtoReflect.Descriptor instead.
func (*GroupID) Descriptor() ([]byte, []int) {
	return file_api_v1_group_relation_proto_rawDescGZIP(), []int{2}
}

func (x *GroupID) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

type UserIDs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"user_ids"
	UserIds []string `protobuf:"bytes,1,rep,name=user_ids,json=userIds,proto3" json:"user_ids"`
}

func (x *UserIDs) Reset() {
	*x = UserIDs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_group_relation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserIDs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIDs) ProtoMessage() {}

func (x *UserIDs) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_group_relation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIDs.ProtoReflect.Descriptor instead.
func (*UserIDs) Descriptor() ([]byte, []int) {
	return file_api_v1_group_relation_proto_rawDescGZIP(), []int{3}
}

func (x *UserIDs) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

var File_api_v1_group_relation_proto protoreflect.FileDescriptor

var file_api_v1_group_relation_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76,
	0x31, 0x22, 0x92, 0x01, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x68,
	0x6f, 0x77, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x73, 0x68, 0x6f, 0x77, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a,
	0x0f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x13, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x0a, 0x07, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x22, 0x24, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73, 0x12, 0x19, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x32, 0x83, 0x01, 0x0a, 0x14, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3e, 0x0a, 0x0f, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x14, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x44, 0x73, 0x12, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44,
	0x1a, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73, 0x42, 0x37, 0x5a,
	0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x73,
	0x69, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_group_relation_proto_rawDescOnce sync.Once
	file_api_v1_group_relation_proto_rawDescData = file_api_v1_group_relation_proto_rawDesc
)

func file_api_v1_group_relation_proto_rawDescGZIP() []byte {
	file_api_v1_group_relation_proto_rawDescOnce.Do(func() {
		file_api_v1_group_relation_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_group_relation_proto_rawDescData)
	})
	return file_api_v1_group_relation_proto_rawDescData
}

var file_api_v1_group_relation_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_v1_group_relation_proto_goTypes = []interface{}{
	(*UserGroupRequest)(nil),  // 0: v1.UserGroupRequest
	(*UserGroupResponse)(nil), // 1: v1.UserGroupResponse
	(*GroupID)(nil),           // 2: v1.GroupID
	(*UserIDs)(nil),           // 3: v1.UserIDs
}
var file_api_v1_group_relation_proto_depIdxs = []int32{
	0, // 0: v1.GroupRelationService.InsertUserGroup:input_type -> v1.UserGroupRequest
	2, // 1: v1.GroupRelationService.GetUserGroupIDs:input_type -> v1.GroupID
	1, // 2: v1.GroupRelationService.InsertUserGroup:output_type -> v1.UserGroupResponse
	3, // 3: v1.GroupRelationService.GetUserGroupIDs:output_type -> v1.UserIDs
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1_group_relation_proto_init() }
func file_api_v1_group_relation_proto_init() {
	if File_api_v1_group_relation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_group_relation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserGroupRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_group_relation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserGroupResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_group_relation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_group_relation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserIDs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_group_relation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_group_relation_proto_goTypes,
		DependencyIndexes: file_api_v1_group_relation_proto_depIdxs,
		MessageInfos:      file_api_v1_group_relation_proto_msgTypes,
	}.Build()
	File_api_v1_group_relation_proto = out.File
	file_api_v1_group_relation_proto_rawDesc = nil
	file_api_v1_group_relation_proto_goTypes = nil
	file_api_v1_group_relation_proto_depIdxs = nil
}
