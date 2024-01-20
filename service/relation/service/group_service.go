package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/cossim/coss-server/pkg/code"
	v1 "github.com/cossim/coss-server/service/relation/api/v1"
	"github.com/cossim/coss-server/service/relation/domain/entity"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

func (s *Service) GetGroupUserIDs(ctx context.Context, id *v1.GroupIDRequest) (*v1.UserIdsResponse, error) {
	resp := &v1.UserIdsResponse{}

	// 调用持久层方法获取用户群关系列表
	userGroupIDs, err := s.grr.GetGroupUserIDs(id.GetGroupId())
	if err != nil {
		return resp, status.Error(codes.Code(code.GroupErrGetBatchGroupInfoByIDsFailed.Code()), err.Error())
	}

	resp.UserIds = userGroupIDs
	return resp, nil
}

func (s *Service) GetUserGroupIDs(ctx context.Context, request *v1.GetUserGroupIDsRequest) (*v1.GetUserGroupIDsResponse, error) {
	resp := &v1.GetUserGroupIDsResponse{}

	ds, err := s.grr.GetUserJoinedGroupIDs(request.UserId)
	if err != nil {
		return resp, status.Error(codes.Code(code.RelationErrGetGroupIDsFailed.Code()), err.Error())
	}

	for _, v := range ds {
		resp.GroupId = append(resp.GroupId, v)
	}
	return resp, nil
}

func (s *Service) GetUserGroupRequestList(ctx context.Context, request *v1.GetUserGroupRequestListRequest) (*v1.GetUserGroupRequestListResponse, error) {
	resp := &v1.GetUserGroupRequestListResponse{}

	ids, err := s.grr.GetUserGroupIDs(request.UserId)
	if err != nil {
		return resp, status.Error(codes.Code(code.GroupErrGetBatchGroupInfoByIDsFailed.Code()), err.Error())
	}

	reqList, err := s.grr.GetJoinRequestBatchListByID(ids)
	if err != nil {
		return resp, status.Error(codes.Code(code.RelationUserErrGetRequestListFailed.Code()), err.Error())
	}

	for _, v := range reqList {
		if v.UserID == request.UserId {
			continue
		}
		resp.GroupJoinRequestList = append(resp.GroupJoinRequestList, &v1.GroupJoinRequestList{
			GroupId:   uint32(v.GroupID),
			UserId:    v.UserID,
			Msg:       v.Remark,
			Status:    v1.GroupRelationStatus(v.Status),
			CreatedAt: v.CreatedAt,
		})
	}
	return resp, nil
}

func (s *Service) JoinGroup(ctx context.Context, request *v1.JoinGroupRequest) (*v1.JoinGroupResponse, error) {
	resp := &v1.JoinGroupResponse{}

	userGroup := &entity.GroupRelation{
		UserID:   request.UserId,
		GroupID:  uint(request.GroupId),
		Remark:   request.Msg,
		Status:   entity.GroupStatusApplying,
		Identity: entity.GroupIdentity(request.Identify),
	}

	// 检查是否已经存在加入申请
	relation, err := s.grr.GetUserGroupByID(request.GroupId, request.UserId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return resp, status.Error(codes.Code(code.RelationGroupErrRequestFailed.Code()), err.Error())
	}

	if relation != nil && relation.Status == entity.GroupStatusApplying {
		return resp, status.Error(codes.Code(code.RelationGroupErrRequestAlreadyPending.Code()), code.RelationGroupErrRequestFailed.Message())
	}
	//如果是群主，则直接加入群组
	if request.Identify == v1.GroupIdentity_IDENTITY_OWNER {
		userGroup.Status = entity.GroupStatusJoined
	}

	//return resp, status.Error(codes.Aborted, formatErrorMessage(errors.New("测试回滚")))

	// 插入加入申请
	_, err = s.grr.CreateRelation(userGroup)
	if err != nil {
		//return resp, status.Error(codes.Code(code.RelationGroupErrRequestFailed.Code()), err.Error())
		return resp, status.Error(codes.Aborted, err.Error())
	}

	return resp, nil
}

func (s *Service) JoinGroupRevert(ctx context.Context, request *v1.JoinGroupRequest) (*v1.JoinGroupResponse, error) {
	resp := &v1.JoinGroupResponse{}
	if err := s.grr.DeleteRelationByID(request.GroupId, request.UserId); err != nil {
		return resp, status.Error(codes.Code(code.RelationGroupErrRequestFailed.Code()), err.Error())
	}
	return resp, nil
}

func (s *Service) ManageJoinGroup(ctx context.Context, request *v1.ManageJoinGroupRequest) (*v1.ManageJoinGroupResponse, error) {
	resp := &v1.ManageJoinGroupResponse{}

	//return resp, status.Error(codes.Aborted, formatErrorMessage(errors.New("测试回滚")))

	// 检查是否已经存在加入申请
	relation, err := s.grr.GetUserGroupByID(request.GroupId, request.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, status.Error(codes.Code(code.RelationGroupErrNoJoinRequestRecords.Code()), code.RelationGroupErrNoJoinRequestRecords.Message())
		}
		return resp, status.Error(codes.Code(code.RelationGroupErrRequestFailed.Code()), err.Error())
	}

	if relation == nil {
		return resp, status.Error(codes.Code(code.RelationGroupErrNoJoinRequestRecords.Code()), code.RelationGroupErrNoJoinRequestRecords.Message())
	}

	if relation.Status == entity.GroupStatusJoined {
		return resp, status.Error(codes.Code(code.RelationGroupErrAlreadyInGroup.Code()), code.RelationGroupErrAlreadyInGroup.Message())
	}

	relation.Status = entity.GroupRelationStatus(request.Status)
	// 插入加入申请
	_, err = s.grr.UpdateRelation(relation)
	if err != nil {
		//return resp, status.Error(codes.Code(code.RelationGroupErrRequestFailed.Code()), err.Error())
		return resp, status.Error(codes.Aborted, fmt.Sprintf("Failed to revert join request : %v", err))
	}

	return resp, nil
}

func (s *Service) ManageJoinGroupRevert(ctx context.Context, request *v1.ManageJoinGroupRequest) (*v1.ManageJoinGroupResponse, error) {
	resp := &v1.ManageJoinGroupResponse{}
	if err := s.grr.UpdateRelationColumnByGroupAndUser(request.GroupId, request.UserId, "status", request.Status); err != nil {
		return resp, status.Error(codes.Code(code.RelationGroupErrRequestFailed.Code()), err.Error())
	}
	return resp, nil
}

func (s *Service) RemoveFromGroup(ctx context.Context, request *v1.RemoveFromGroupRequest) (*v1.RemoveFromGroupResponse, error) {
	resp := &v1.RemoveFromGroupResponse{}

	if err := s.grr.DeleteRelationByID(request.GroupId, request.UserId); err != nil {
		return resp, status.Error(codes.Code(code.RelationGroupErrRemoveUserFromGroupFailed.Code()), err.Error())
	}

	return resp, nil
}

func (s *Service) LeaveGroup(ctx context.Context, request *v1.LeaveGroupRequest) (*v1.LeaveGroupResponse, error) {
	resp := &v1.LeaveGroupResponse{}
	if err := s.grr.DeleteRelationByID(request.GroupId, request.UserId); err != nil {
		return resp, status.Error(codes.Aborted, err.Error())
	}
	return resp, nil
}

func (s *Service) LeaveGroupRevert(ctx context.Context, request *v1.LeaveGroupRequest) (*v1.LeaveGroupResponse, error) {
	fmt.Println("revert leave group req => ", request)
	resp := &v1.LeaveGroupResponse{}

	if err := s.grr.UpdateRelationColumnByGroupAndUser(request.GroupId, request.UserId, "deleted_at", 0); err != nil {
		return resp, status.Error(codes.Code(code.RelationGroupErrLeaveGroupFailed.Code()), err.Error())
	}

	return resp, nil
}

func (s *Service) GetGroupJoinRequestList(ctx context.Context, request *v1.GetGroupJoinRequestListRequest) (*v1.GroupJoinRequestListResponse, error) {
	resp := &v1.GroupJoinRequestListResponse{}

	var ids []uint32
	for _, v := range request.GroupIds {
		ids = append(ids, v.GroupId)
	}

	joins, err := s.grr.GetJoinRequestBatchListByID(ids)
	if err != nil {
		return resp, status.Error(codes.Code(code.RelationUserErrGetRequestListFailed.Code()), err.Error())
	}

	for _, join := range joins {
		resp.GroupJoinRequestList = append(resp.GroupJoinRequestList, &v1.GroupJoinRequestList{
			UserId: join.UserID,
			Msg:    join.Remark,
			Status: v1.GroupRelationStatus(join.Status),
		})
	}

	return resp, nil
}

func (s *Service) GetGroupRelation(ctx context.Context, request *v1.GetGroupRelationRequest) (*v1.GetGroupRelationResponse, error) {
	resp := &v1.GetGroupRelationResponse{}

	relation, err := s.grr.GetUserGroupByID(request.GroupId, request.UserId)
	if err != nil {
		return resp, status.Error(codes.Code(code.RelationGroupErrGroupRelationFailed.Code()), err.Error())
	}

	resp.GroupId = uint32(relation.GroupID)
	resp.UserId = relation.UserID
	resp.Identity = v1.GroupIdentity(uint32(relation.Identity))
	resp.MuteEndTime = relation.MuteEndTime
	resp.Status = v1.GroupRelationStatus(uint32(relation.Status))
	return resp, nil
}

func (s *Service) DeleteGroupRelationByGroupId(ctx context.Context, in *v1.GroupIDRequest) (*emptypb.Empty, error) {
	err := s.grr.DeleteGroupRelationByID(in.GroupId)
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.Aborted, err.Error())
	}
	return &emptypb.Empty{}, nil
}

func (s *Service) DeleteGroupRelationByGroupIdRevert(ctx context.Context, request *v1.GroupIDRequest) (*emptypb.Empty, error) {
	resp := &emptypb.Empty{}
	fmt.Println("DeleteGroupRelationByGroupIdRevert req => ", request)

	if err := s.grr.UpdateGroupRelationByGroupID(request.GroupId, map[string]interface{}{
		"deleted_at": 0,
	}); err != nil {
		return resp, status.Error(codes.Code(code.GroupErrDeleteGroupFailed.Code()), err.Error())
	}
	return resp, nil
}

func (s *Service) GetGroupAdminIds(ctx context.Context, in *v1.GroupIDRequest) (*v1.UserIdsResponse, error) {
	var resp = &v1.UserIdsResponse{}
	ids, err := s.grr.GetGroupAdminIds(in.GroupId)
	if err != nil {
		return resp, err
	}
	resp.UserIds = ids
	return resp, nil
}

func (s *Service) GetUserManageGroupID(ctx context.Context, request *v1.GetUserManageGroupIDRequest) (*v1.GetUserManageGroupIDResponse, error) {
	resp := &v1.GetUserManageGroupIDResponse{}

	ids, err := s.grr.GetUserManageGroupIDs(request.UserId)
	if err != nil {
		return resp, status.Error(codes.Code(code.GroupErrGetBatchGroupInfoByIDsFailed.Code()), err.Error())
	}

	for _, id := range ids {
		resp.GroupIDs = append(resp.GroupIDs, &v1.GroupIDRequest{GroupId: id})
	}

	return resp, nil
}

func (s *Service) DeleteGroupRelationByGroupIdAndUserID(ctx context.Context, request *v1.DeleteGroupRelationByGroupIdAndUserIDRequest) (*emptypb.Empty, error) {
	resp := &emptypb.Empty{}

	//return resp, status.Error(codes.Aborted, formatErrorMessage(errors.New("测试回滚")))

	if err := s.grr.DeleteUserGroupRelationByGroupIDAndUserID(request.GroupID, request.UserID); err != nil {
		//return resp, status.Error(codes.Code(code.GroupErrDeleteUserGroupRelationFailed.Code()), err.Error())
		return resp, status.Error(codes.Aborted, fmt.Sprintf(code.GroupErrDeleteUserGroupRelationFailed.Message()+" : "+err.Error()))
	}
	return resp, nil
}

func (s *Service) DeleteGroupRelationByGroupIdAndUserIDRevert(ctx context.Context, request *v1.DeleteGroupRelationByGroupIdAndUserIDRequest) (*emptypb.Empty, error) {
	resp := &emptypb.Empty{}
	if err := s.grr.UpdateRelationColumnByGroupAndUser(request.GroupID, request.UserID, "deleted_at", 0); err != nil {
		//return resp, status.Error(codes.Code(code.GroupErrDeleteUserGroupRelationFailed.Code()), err.Error())
		return resp, status.Error(codes.Code(code.GroupErrDeleteUserGroupRelationRevertFailed.Code()), err.Error())
	}
	if err := s.grr.UpdateRelationColumnByGroupAndUser(request.GroupID, request.UserID, "status", request.Status); err != nil {
		//return resp, status.Error(codes.Code(code.GroupErrDeleteUserGroupRelationFailed.Code()), err.Error())
		return resp, status.Error(codes.Code(code.GroupErrDeleteUserGroupRelationRevertFailed.Code()), err.Error())
	}
	return resp, nil
}
