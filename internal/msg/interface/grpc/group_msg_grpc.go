package grpc

import (
	"context"
	"fmt"
	v1 "github.com/cossim/coss-server/internal/msg/api/grpc/v1"
	"github.com/cossim/coss-server/internal/msg/domain/entity"
	"github.com/cossim/coss-server/internal/msg/infrastructure/persistence"
	"github.com/cossim/coss-server/pkg/code"
	"github.com/cossim/coss-server/pkg/utils/time"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (s *Handler) SetGroupMessageRead(ctx context.Context, in *v1.SetGroupMessagesReadRequest) (*v1.SetGroupMessageReadResponse, error) {
	resp := &v1.SetGroupMessageReadResponse{}
	var reads []*entity.GroupMessageRead
	msgids := make([]uint32, 0)

	if len(in.List) > 0 {
		reads = make([]*entity.GroupMessageRead, len(in.List))
		for i, _ := range in.List {
			reads[i] = &entity.GroupMessageRead{
				DialogId: uint(in.List[i].DialogId),
				MsgId:    uint(in.List[i].MsgId),
				UserId:   in.List[i].UserId,
				GroupID:  uint(in.List[i].GroupId),
				ReadAt:   in.List[i].ReadAt,
			}
			msgids = append(msgids, in.List[i].MsgId)
		}
	}
	err := s.db.Transaction(func(tx *gorm.DB) error {
		npo := persistence.NewRepositories(tx)
		msgs, err := npo.Mr.GetGroupMsgsByIDs(msgids)
		//修改已读数量
		for _, v := range msgs {
			v.ReadCount++
			err := npo.Mr.UpdateGroupMsgColumn(uint32(v.ID), "read_count", v.ReadCount)
			if err != nil {
				return err
			}
		}
		err = npo.Gmrr.SetGroupMsgReadByMsgs(reads)
		if err != nil {
			return status.Error(codes.Code(code.GroupErrSetGroupMsgReadFailed.Code()), err.Error())
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *Handler) GetGroupMessageReaders(ctx context.Context, in *v1.GetGroupMessageReadersRequest) (*v1.GetGroupMessageReadersResponse, error) {
	resp := &v1.GetGroupMessageReadersResponse{}
	msgs, err := s.gmrr.GetGroupMsgReadUserIdsByMsgID(in.MsgId)
	if err != nil {
		return nil, status.Error(codes.Code(code.GroupErrGetGroupMsgReadersFailed.Code()), err.Error())
	}
	resp.UserIds = msgs

	return resp, err
}

func (s *Handler) GetGroupMessageReadByMsgIdAndUserId(ctx context.Context, in *v1.GetGroupMessageReadByMsgIdAndUserIdRequest) (*v1.GetGroupMessageReadByMsgIdAndUserIdResponse, error) {
	resp := &v1.GetGroupMessageReadByMsgIdAndUserIdResponse{}
	msg, err := s.gmrr.GetGroupMsgReadByMsgIDAndUserID(in.MsgId, in.UserId)
	if err != nil {
		return nil, status.Error(codes.Code(code.GroupErrGetGroupMsgReadByMsgIdAndUserIdFailed.Code()), err.Error())
	}
	resp.ReadAt = msg.ReadAt
	return resp, err
}

func (s *Handler) ReadAllGroupMsg(ctx context.Context, request *v1.ReadAllGroupMsgRequest) (*v1.ReadAllGroupMsgResponse, error) {
	resp := &v1.ReadAllGroupMsgResponse{}

	msgids, err := s.gmrr.GetGroupMsgUserReadIdsByDialogID(request.DialogId, request.UserId)
	if err != nil {
		return resp, status.Error(codes.Code(code.GroupErrGetGroupMsgReadByMsgIdAndUserIdFailed.Code()), err.Error())
	}

	fmt.Println("msgids => ", msgids)

	list, err := s.mr.GetGroupUnreadMsgList(request.DialogId, msgids)
	if err != nil {
		return resp, err
	}
	var reads []*entity.GroupMessageRead

	fmt.Println("list => ", list)

	if len(list) > 0 {
		reads = make([]*entity.GroupMessageRead, len(list))
		for k, v := range list {
			reads[k] = &entity.GroupMessageRead{
				DialogId: v.DialogId,
				MsgId:    v.ID,
				UserId:   request.UserId,
				GroupID:  v.GroupID,
				ReadAt:   time.Now(),
			}
			resp.UnreadGroupMessage = append(resp.UnreadGroupMessage, &v1.UnreadGroupMessage{
				MsgId:  uint32(v.ID),
				UserId: v.UserID,
			})
		}
	}

	if err := s.gmrr.SetGroupMsgReadByMsgs(reads); err != nil {
		return resp, status.Error(codes.Code(code.GroupErrGetGroupMsgReadByMsgIdAndUserIdFailed.Code()), err.Error())
	}

	return resp, nil
}
