package grpc

import (
	"context"
	"fmt"
	"github.com/cossim/coss-server/internal/msg/api/grpc/dataTransformers"
	v1 "github.com/cossim/coss-server/internal/msg/api/grpc/v1"
	"github.com/cossim/coss-server/internal/msg/domain/entity"
	"github.com/cossim/coss-server/internal/msg/infrastructure/persistence"
	"github.com/cossim/coss-server/pkg/code"
	"github.com/cossim/coss-server/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"time"
)

func (s *Handler) SendUserMessage(ctx context.Context, request *v1.SendUserMsgRequest) (*v1.SendUserMsgResponse, error) {
	resp := &v1.SendUserMsgResponse{}

	msg, err := s.mr.InsertUserMessage(request.GetSenderId(), request.GetReceiverId(), request.GetContent(), entity.UserMessageType(request.GetType()), uint(request.GetReplyId()), uint(request.GetDialogId()), entity.BurnAfterReadingType(request.IsBurnAfterReadingType))
	if err != nil {
		return resp, status.Error(codes.Code(code.MsgErrInsertUserMessageFailed.Code()), err.Error())
	}
	resp.MsgId = uint32(msg.ID)

	//if s.cacheEnable {
	//	fmt.Println("DeleteLastMessage => ", request.DialogId)
	//	if err := s.cache.DeleteLastMessage(ctx, request.DialogId); err != nil {
	//		log.Printf("delete last message failed: %v", err)
	//	}
	//}

	return resp, err
}

func (s *Handler) SendGroupMessage(ctx context.Context, request *v1.SendGroupMsgRequest) (*v1.SendGroupMsgResponse, error) {
	resp := &v1.SendGroupMsgResponse{}

	ums, err := s.mr.InsertGroupMessage(request.GetUserId(), uint(request.GetGroupId()), request.GetContent(), entity.UserMessageType(request.GetType()), uint(request.GetReplyId()), uint(request.GetDialogId()), entity.BurnAfterReadingType(request.IsBurnAfterReadingType), request.AtUsers, entity.AtAllUserType(request.AtAllUser))
	if err != nil {
		return resp, status.Error(codes.Code(code.MsgErrInsertGroupMessageFailed.Code()), err.Error())
	}

	//if s.cacheEnable {
	//	fmt.Println("DeleteLastMessage => ", request.DialogId)
	//	if err := s.cache.DeleteLastMessage(ctx, request.DialogId); err != nil {
	//		log.Printf("delete last message failed: %v", err)
	//	}
	//}

	return &v1.SendGroupMsgResponse{
		MsgId:   uint32(ums.ID),
		GroupId: uint32(ums.GroupID),
	}, nil
}

func (s *Handler) SendUserMessageRevert(ctx context.Context, request *v1.MsgIdRequest) (*v1.SendUserMsgRevertResponse, error) {
	resp := &v1.SendUserMsgRevertResponse{}
	if err := s.mr.PhysicalDeleteUserMessage(request.MsgId); err != nil {
		return resp, status.Error(codes.Code(code.MsgErrDeleteUserMessageFailed.Code()), err.Error())
	}
	return resp, nil
}

func (s *Handler) SendGroupMessageRevert(ctx context.Context, request *v1.MsgIdRequest) (*v1.SendGroupMsgRevertResponse, error) {
	resp := &v1.SendGroupMsgRevertResponse{}
	if err := s.mr.PhysicalDeleteGroupMessage(request.MsgId); err != nil {
		return resp, status.Error(codes.Code(code.MsgErrDeleteGroupMessageFailed.Code()), err.Error())
	}
	return resp, nil
}

func (s *Handler) GetUserMessageList(ctx context.Context, request *v1.GetUserMsgListRequest) (*v1.GetUserMsgListResponse, error) {
	resp := &v1.GetUserMsgListResponse{}

	if request.MsgId != 0 {
		//list, total, err := s.mr.GetUserMsgIdBeforeMsgList(request.DialogId, uint32(request.MsgId), int(request.PageSize))
		//if err != nil {
		//	return nil, err
		//}

		res, err := s.mr.Find(ctx, &entity.UserMsgQuery{
			DialogIds: []uint32{request.DialogId},
			MsgIds:    []uint32{uint32(request.MsgId)},
			PageSize:  int64(request.PageSize),
		})
		if err != nil {
			return nil, err
		}

		for _, m := range res.Messages {
			resp.UserMessages = append(resp.UserMessages, &v1.UserMessage{
				Id:         uint32(m.ID),
				SenderId:   m.SendID,
				ReceiverId: m.ReceiveID,
				Content:    m.Content,
				Type:       uint32(int32(m.Type)),
				ReplyId:    uint64(m.ReplyId),
				IsRead:     int32(m.IsRead),
				ReadAt:     m.ReadAt,
				CreatedAt:  m.CreatedAt,
				DialogId:   uint32(m.DialogId),
			})
		}
		resp.Total = int32(res.TotalCount)
		return resp, nil
	}
	//ums, total, current := s.mr.GetUserMsgList(request.DialogId, request.UserId, request.GetContent(), entity.UserMessageType(request.GetType()), int(request.GetPageNum()), int(request.GetPageSize()))

	res, err := s.mr.Find(ctx, &entity.UserMsgQuery{
		DialogIds: []uint32{request.DialogId},
		MsgType:   entity.UserMessageType(request.Type),
		PageNum:   int64(request.PageNum),
		PageSize:  int64(request.PageSize),
		Content:   request.Content,
		SendID:    request.UserId,
		StartAt:   int64(request.StartAt),
		EndAt:     int64(request.EndAt),
	})
	if err != nil {
		return nil, err
	}

	for _, m := range res.Messages {
		resp.UserMessages = append(resp.UserMessages, &v1.UserMessage{
			Id:         uint32(m.ID),
			SenderId:   m.SendID,
			ReceiverId: m.ReceiveID,
			Content:    m.Content,
			Type:       uint32(int32(m.Type)),
			ReplyId:    uint64(m.ReplyId),
			IsRead:     int32(m.IsRead),
			ReadAt:     m.ReadAt,
			CreatedAt:  m.CreatedAt,
			DialogId:   uint32(m.DialogId),
		})
	}
	resp.Total = int32(res.TotalCount)
	resp.CurrentPage = int32(res.CurrentPage)

	return resp, nil
}

func (s *Handler) GetLastMsgsForUserWithFriends(ctx context.Context, request *v1.UserMsgsRequest) (*v1.UserMessages, error) {
	resp := &v1.UserMessages{}
	msgs, err := s.mr.GetLastMsgsForUserWithFriends(request.UserId, request.FriendId)
	if err != nil {
		return resp, status.Error(codes.Code(code.MsgErrGetLastMsgsForUserWithFriends.Code()), err.Error())
	}
	nmsgs := make([]*v1.UserMessage, 0)
	for _, m := range msgs {
		nmsgs = append(nmsgs, &v1.UserMessage{
			Id:        uint32(m.ID),
			Content:   m.Content,
			Type:      uint32(m.Type),
			ReplyId:   uint64(m.ReplyId),
			ReadAt:    m.ReadAt,
			CreatedAt: m.CreatedAt,
		})
	}
	resp.UserMessages = nmsgs
	return resp, nil
}

func (s *Handler) GetLastMsgsForGroupsWithIDs(ctx context.Context, request *v1.GroupMsgsRequest) (*v1.GroupMessages, error) {
	resp := &v1.GroupMessages{}
	ids := make([]uint, 0)
	if len(request.GroupId) > 0 {
		for _, id := range request.GroupId {
			ids = append(ids, uint(id))
		}
	}
	msgs, err := s.mr.GetLastMsgsForGroupsWithIDs(ids)
	if err != nil {
		return resp, status.Error(codes.Code(code.MsgErrGetLastMsgsForGroupsWithIDs.Code()), err.Error())
	}
	nmsgs := make([]*v1.GroupMessage, 0)
	for _, m := range msgs {
		nmsgs = append(nmsgs, &v1.GroupMessage{
			Id:        uint32(m.ID),
			UserId:    m.UserID,
			Content:   m.Content,
			Type:      uint32(m.Type),
			ReplyId:   uint32(m.ReplyId),
			ReadCount: int32(m.ReadCount),
			CreatedAt: m.CreatedAt,
		})
	}
	resp.GroupMessages = nmsgs
	return resp, nil
}

func (s *Handler) GetLastMsgsByDialogIds(ctx context.Context, request *v1.GetLastMsgsByDialogIdsRequest) (*v1.GetLastMsgsResponse, error) {
	resp := &v1.GetLastMsgsResponse{}

	ids := make([]uint, 0)
	if len(request.DialogIds) > 0 {
		for _, id := range request.DialogIds {
			ids = append(ids, uint(id))
		}
	}

	//获取群聊对话最后一条消息
	result1, err := s.mr.GetLastGroupMsgsByDialogIDs(ids)
	if err != nil {
		return resp, status.Error(codes.Code(code.MsgErrGetLastMsgsByDialogIds.Code()), err.Error())
	}

	result2, err := s.mr.GetLastUserMsgsByDialogIDs(ids)
	if err != nil {
		return resp, status.Error(codes.Code(code.MsgErrGetLastMsgsByDialogIds.Code()), err.Error())
	}

	for _, m := range result1 {
		//查询是否已读
		read, err := s.gmrr.GetGroupMsgReadByMsgIDAndUserID(uint32(m.ID), request.UserId)
		if err != nil && err != gorm.ErrRecordNotFound {
			return resp, status.Error(codes.Code(code.MsgErrGetLastMsgsByDialogIds.Code()), err.Error())
		}

		isRead := v1.ReadType_NotRead
		if read.ReadAt != 0 {
			isRead = v1.ReadType_IsRead
		}

		resp.LastMsgs = append(resp.LastMsgs, &v1.LastMsg{
			Id:                     uint32(m.ID),
			Type:                   uint32(m.Type),
			CreatedAt:              m.CreatedAt,
			Content:                m.Content,
			SenderId:               m.UserID,
			DialogId:               uint32(m.DialogId),
			IsBurnAfterReadingType: v1.BurnAfterReadingType(m.IsBurnAfterReading),
			AtUsers:                m.AtUsers,
			AtAllUser:              v1.AtAllUserType(m.AtAllUser),
			IsLabel:                v1.MsgLabel(m.IsLabel),
			ReplyId:                uint32(m.ReplyId),
			IsRead:                 int32(isRead),
			GroupId:                uint32(m.GroupID),
			ReadAt:                 read.ReadAt,
		})
	}

	for _, m := range result2 {
		resp.LastMsgs = append(resp.LastMsgs, &v1.LastMsg{
			Id:                     uint32(m.ID),
			Type:                   uint32(m.Type),
			CreatedAt:              m.CreatedAt,
			Content:                m.Content,
			SenderId:               m.SendID,
			DialogId:               uint32(m.DialogId),
			ReceiverId:             m.ReceiveID,
			IsBurnAfterReadingType: v1.BurnAfterReadingType(m.IsBurnAfterReading),
			IsLabel:                v1.MsgLabel(m.IsLabel),
			ReplyId:                uint32(m.ReplyId),
			IsRead:                 int32(m.IsRead),
			ReadAt:                 m.ReadAt,
		})
	}

	//if s.cacheEnable {
	//	for _, v := range resp.LastMsgs {
	//		if err := s.cache.SetLastMessage(ctx, v.DialogId, v, cache.MsgExpireTime); err != nil {
	//			log.Printf("set last message to cache failed, err: %v", err)
	//		}
	//	}
	//}

	return resp, nil
}

func (s *Handler) EditUserMessage(ctx context.Context, request *v1.EditUserMsgRequest) (*v1.UserMessage, error) {
	var resp = &v1.UserMessage{}
	msg := entity.UserMessage{
		BaseModel: entity.BaseModel{
			ID: uint(request.UserMessage.Id),
		},
		Content:   request.UserMessage.Content,
		ReplyId:   uint(request.UserMessage.ReplyId),
		SendID:    request.UserMessage.SenderId,
		ReceiveID: request.UserMessage.ReceiverId,
		Type:      entity.UserMessageType(request.UserMessage.Type),
		IsLabel:   uint(request.UserMessage.IsLabel),
	}
	if err := s.mr.UpdateUserMessage(msg); err != nil {
		return resp, status.Error(codes.Code(code.MsgErrEditUserMessageFailed.Code()), err.Error())
	}
	//
	//if s.cacheEnable {
	//	if err := s.cache.DeleteLastMessage(ctx, request.UserMessage.DialogId); err != nil {
	//		log.Printf("delete last message failed: %v", err)
	//	}
	//}

	return resp, nil
}

func (s *Handler) DeleteUserMessage(ctx context.Context, request *v1.DeleteUserMsgRequest) (*v1.UserMessage, error) {
	var resp = &v1.UserMessage{}
	err := s.mr.LogicalDeleteUserMessage(request.MsgId)
	if err != nil {
		return resp, status.Error(codes.Code(code.MsgErrDeleteUserMessageFailed.Code()), err.Error())
	}

	//if s.cacheEnable {
	//	if err := s.cache.DeleteLastMessage(ctx, request.DialogId); err != nil {
	//		log.Printf("delete last message failed: %v", err)
	//	}
	//}

	return &v1.UserMessage{
		Id: request.MsgId,
	}, nil
}

func (s *Handler) EditGroupMessage(ctx context.Context, request *v1.EditGroupMsgRequest) (*v1.GroupMessage, error) {
	var resp = &v1.GroupMessage{}
	msg := entity.GroupMessage{
		BaseModel: entity.BaseModel{
			ID: uint(request.GroupMessage.Id),
		},
		Content: request.GroupMessage.Content,
		ReplyId: uint(request.GroupMessage.ReplyId),
		UserID:  request.GroupMessage.UserId,
		GroupID: uint(request.GroupMessage.GroupId),
		Type:    entity.UserMessageType(request.GroupMessage.Type),
		IsLabel: uint(request.GroupMessage.IsLabel),
	}
	if err := s.mr.UpdateGroupMessage(msg); err != nil {
		return resp, status.Error(codes.Code(code.MsgErrEditGroupMessageFailed.Code()), err.Error())
	}
	resp = &v1.GroupMessage{
		Id:        uint32(msg.ID),
		UserId:    msg.UserID,
		Content:   msg.Content,
		Type:      uint32(int32(msg.Type)),
		ReplyId:   uint32(msg.ReplyId),
		GroupId:   uint32(msg.GroupID),
		ReadCount: int32(msg.ReadCount),
	}
	//
	//if s.cacheEnable {
	//	if err := s.cache.DeleteLastMessage(ctx, request.GroupMessage.DialogId); err != nil {
	//		log.Printf("delete last message failed: %v", err)
	//	}
	//}

	return resp, nil
}

func (s *Handler) DeleteGroupMessage(ctx context.Context, request *v1.DeleteGroupMsgRequest) (*v1.GroupMessage, error) {
	var resp = &v1.GroupMessage{}
	if err := s.mr.LogicalDeleteGroupMessage(request.MsgId); err != nil {
		return resp, status.Error(codes.Code(code.MsgErrDeleteGroupMessageFailed.Code()), err.Error())
	}

	//if s.cacheEnable {
	//	if err := s.cache.DeleteLastMessage(ctx, request.DialogId); err != nil {
	//		log.Printf("delete last message failed: %v", err)
	//	}
	//}

	return &v1.GroupMessage{
		Id: request.MsgId,
	}, nil
}

func (s *Handler) GetUserMessageById(ctx context.Context, in *v1.GetUserMsgByIDRequest) (*v1.UserMessage, error) {
	var resp = &v1.UserMessage{}
	msg, err := s.mr.GetUserMsgByID(in.MsgId)
	if err != nil {
		return resp, status.Error(codes.Code(code.GetMsgErrGetUserMsgByIDFailed.Code()), err.Error())
	}

	resp = &v1.UserMessage{
		DialogId:   uint32(msg.DialogId),
		Id:         uint32(msg.ID),
		Content:    msg.Content,
		Type:       uint32(int32(msg.Type)),
		ReplyId:    uint64(msg.ReplyId),
		SenderId:   msg.SendID,
		ReceiverId: msg.ReceiveID,
		CreatedAt:  msg.CreatedAt,
	}
	return resp, nil
}

func (s *Handler) GetGroupMessageById(ctx context.Context, in *v1.GetGroupMsgByIDRequest) (*v1.GroupMessage, error) {
	var resp = &v1.GroupMessage{}
	msg, err := s.mr.GetGroupMsgByID(in.MsgId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return resp, status.Error(codes.Code(code.GetMsgErrGetGroupMsgByIDFailed.Code()), err.Error())
		}
		return resp, err
	}
	return &v1.GroupMessage{
		Id:        uint32(msg.ID),
		UserId:    msg.UserID,
		DialogId:  uint32(msg.DialogId),
		Content:   msg.Content,
		Type:      uint32(int32(msg.Type)),
		ReplyId:   uint32(msg.ReplyId),
		GroupId:   uint32(msg.GroupID),
		ReadCount: int32(msg.ReadCount),
		CreatedAt: msg.CreatedAt,
	}, nil
}

func (s *Handler) SetUserMsgLabel(ctx context.Context, in *v1.SetUserMsgLabelRequest) (*v1.SetUserMsgLabelResponse, error) {
	var resp = &v1.SetUserMsgLabelResponse{}
	if err := s.mr.UpdateUserMsgColumn(in.MsgId, "is_label", in.IsLabel); err != nil {
		return resp, status.Error(codes.Code(code.SetMsgErrSetUserMsgLabelFailed.Code()), err.Error())
	}
	resp.MsgId = in.MsgId
	return resp, nil
}

func (s *Handler) SetGroupMsgLabel(ctx context.Context, in *v1.SetGroupMsgLabelRequest) (*v1.SetGroupMsgLabelResponse, error) {
	var resp = &v1.SetGroupMsgLabelResponse{}
	if err := s.mr.UpdateGroupMsgColumn(in.MsgId, "is_label", in.IsLabel); err != nil {
		return resp, status.Error(codes.Code(code.SetMsgErrSetGroupMsgLabelFailed.Code()), err.Error())
	}
	resp.MsgId = in.MsgId
	return resp, nil
}

func (s *Handler) GetUserMsgLabelByDialogId(ctx context.Context, in *v1.GetUserMsgLabelByDialogIdRequest) (*v1.GetUserMsgLabelByDialogIdResponse, error) {
	var resp = &v1.GetUserMsgLabelByDialogIdResponse{}

	msgs, err := s.mr.GetUserMsgLabelByDialogId(in.DialogId)
	if err != nil {
		return resp, status.Error(codes.Code(code.GetMsgErrGetUserMsgLabelByDialogIdFailed.Code()), err.Error())
	}

	var msglist = make([]*v1.UserMessage, 0)
	for _, msg := range msgs {
		msglist = append(msglist, &v1.UserMessage{
			Id:         uint32(msg.ID),
			Content:    msg.Content,
			Type:       uint32(msg.Type),
			ReplyId:    uint64(msg.ReplyId),
			SenderId:   msg.SendID,
			ReceiverId: msg.ReceiveID,
			CreatedAt:  msg.CreatedAt,
		})
	}
	resp.MsgList = msglist
	resp.DialogId = in.DialogId
	return resp, nil
}

func (s *Handler) GetGroupMsgLabelByDialogId(ctx context.Context, in *v1.GetGroupMsgLabelByDialogIdRequest) (*v1.GetGroupMsgLabelByDialogIdResponse, error) {
	var resp = &v1.GetGroupMsgLabelByDialogIdResponse{}
	msgs, err := s.mr.GetGroupMsgLabelByDialogId(in.DialogId)
	if err != nil {
		return resp, status.Error(codes.Code(code.GetMsgErrGetGroupMsgLabelByDialogIdFailed.Code()), err.Error())
	}

	var msglist = make([]*v1.GroupMessage, 0)
	for _, msg := range msgs {
		msglist = append(msglist, &v1.GroupMessage{
			Id:        uint32(msg.ID),
			UserId:    msg.UserID,
			Content:   msg.Content,
			Type:      uint32(int32(msg.Type)),
			ReplyId:   uint32(msg.ReplyId),
			GroupId:   uint32(msg.GroupID),
			ReadCount: int32(msg.ReadCount),
		})
	}
	resp.DialogId = in.DialogId
	resp.MsgList = msglist
	return resp, nil
}

func (s *Handler) SetUserMsgsReadStatus(ctx context.Context, in *v1.SetUserMsgsReadStatusRequest) (*v1.SetUserMsgsReadStatusResponse, error) {
	var resp = &v1.SetUserMsgsReadStatusResponse{}
	//获取阅后即焚消息id
	messages, err := s.mr.GetBatchUserMsgsBurnAfterReadingMessages(in.MsgIds, in.DialogId)
	if err != nil {
		return nil, err
	}
	rids := make([]uint32, 0)
	if len(messages) > 0 {
		for _, msg := range messages {
			rids = append(rids, uint32(msg.ID))
		}
	}
	if in.OpenBurnAfterReadingTimeOut == 0 {
		in.OpenBurnAfterReadingTimeOut = 10
	}
	err = s.db.Transaction(func(tx *gorm.DB) error {
		npo := persistence.NewRepositories(tx)
		if err := npo.Mr.SetUserMsgsReadStatus(in.MsgIds, in.DialogId); err != nil {
			return err
		}
		if len(rids) > 0 {
			//起一个携程，定时器根据超时时间删除
			go func(rpo *persistence.Repositories) {
				time.Sleep(time.Duration(in.OpenBurnAfterReadingTimeOut) * time.Second)
				err := rpo.Mr.LogicalDeleteUserMessages(rids)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
			}(npo)
		}

		return nil
	})
	if err != nil {
		return resp, status.Error(codes.Code(code.SetMsgErrSetUserMsgsReadStatusFailed.Code()), err.Error())
	}

	//删除消息
	return resp, nil
}

// 修改用户消息的已读状态
func (s *Handler) SetUserMsgReadStatus(ctx context.Context, in *v1.SetUserMsgReadStatusRequest) (*v1.SetUserMsgReadStatusResponse, error) {
	var resp = &v1.SetUserMsgReadStatusResponse{}
	if err := s.mr.SetUserMsgReadStatus(in.MsgId, entity.ReadType(in.IsRead)); err != nil {
		return resp, status.Error(codes.Code(code.SetMsgErrSetUserMsgReadStatusFailed.Code()), err.Error())
	}
	return resp, nil
}

func (s *Handler) GetUnreadUserMsgs(ctx context.Context, in *v1.GetUnreadUserMsgsRequest) (*v1.GetUnreadUserMsgsResponse, error) {
	var resp = &v1.GetUnreadUserMsgsResponse{}
	msgs, err := s.mr.GetUnreadUserMsgs(in.UserId, in.DialogId)
	if err != nil {
		return resp, status.Error(codes.Code(code.GetMsgErrGetUnreadUserMsgsFailed.Code()), err.Error())
	}
	var list []*v1.UserMessage
	for _, msg := range msgs {
		list = append(list, &v1.UserMessage{
			Id:         uint32(msg.ID),
			Content:    msg.Content,
			Type:       uint32(msg.Type),
			ReplyId:    uint64(msg.ReplyId),
			SenderId:   msg.SendID,
			ReceiverId: msg.ReceiveID,
			CreatedAt:  msg.CreatedAt,
			ReadAt:     msg.ReadAt,
			IsRead:     int32(msg.IsRead),
			DialogId:   uint32(msg.DialogId),
		})
	}
	resp.UserMessages = list
	return resp, nil
}

func (s *Handler) GetUserMsgIdAfterMsgList(ctx context.Context, in *v1.GetUserMsgIdAfterMsgListRequest) (*v1.GetUserMsgIdAfterMsgListResponse, error) {
	var resp = &v1.GetUserMsgIdAfterMsgListResponse{}
	if len(in.List) > 0 {
		for _, v := range in.List {
			list, total, err := s.mr.GetUserMsgIdAfterMsgList(v.DialogId, v.MsgId)
			if err != nil {
				return nil, err
			}
			if len(list) > 0 {
				mlist := &v1.GetUserMsgIdAfterMsgResponse{}
				for _, msg := range list {
					mlist.UserMessages = append(mlist.UserMessages, &v1.UserMessage{
						Id:         uint32(msg.ID),
						Content:    msg.Content,
						Type:       uint32(msg.Type),
						ReplyId:    uint64(msg.ReplyId),
						SenderId:   msg.SendID,
						ReceiverId: msg.ReceiveID,
						CreatedAt:  msg.CreatedAt,
					})
				}
				mlist.Total = uint64(total)
				mlist.DialogId = v.DialogId
				resp.Messages = append(resp.Messages, mlist)
			}
		}
	}
	return resp, nil
}

func (s *Handler) GetGroupMsgIdAfterMsgList(ctx context.Context, in *v1.GetGroupMsgIdAfterMsgListRequest) (*v1.GetGroupMsgIdAfterMsgListResponse, error) {
	var resp = &v1.GetGroupMsgIdAfterMsgListResponse{}
	if len(in.List) > 0 {
		for _, v := range in.List {
			list, total, err := s.mr.GetGroupMsgIdAfterMsgList(v.DialogId, v.MsgId)
			if err != nil {
				return nil, err
			}
			if len(list) > 0 {
				mlist := &v1.GetGroupMsgIdAfterMsgResponse{}
				for _, msg := range list {
					mlist.GroupMessages = append(mlist.GroupMessages, &v1.GroupMessage{
						Id:        uint32(msg.ID),
						UserId:    msg.UserID,
						Content:   msg.Content,
						Type:      uint32(int32(msg.Type)),
						ReplyId:   uint32(msg.ReplyId),
						GroupId:   uint32(msg.GroupID),
						ReadCount: int32(msg.ReadCount),
					})
				}
				mlist.Total = uint64(total)
				mlist.DialogId = v.DialogId
				resp.Messages = append(resp.Messages, mlist)
			}
		}
	}
	return resp, nil
}

func (s *Handler) GetGroupMessageList(ctx context.Context, in *v1.GetGroupMsgListRequest) (*v1.GetGroupMsgListResponse, error) {
	var resp = &v1.GetGroupMsgListResponse{}
	if in.MsgId != 0 {
		list, total, err := s.mr.GetGroupMsgIdBeforeMsgList(in.DialogId, uint32(in.MsgId), int(in.PageSize))
		if err != nil {
			return nil, err
		}

		if len(list) > 0 {
			for _, msg := range list {
				resp.GroupMessages = append(resp.GroupMessages, &v1.GroupMessage{
					Id:        uint32(msg.ID),
					UserId:    msg.UserID,
					Content:   msg.Content,
					Type:      uint32(int32(msg.Type)),
					ReplyId:   uint32(msg.ReplyId),
					GroupId:   uint32(msg.GroupID),
					ReadCount: int32(msg.ReadCount),
					CreatedAt: msg.CreatedAt,
					DialogId:  uint32(msg.DialogId),
					IsLabel:   v1.MsgLabel(msg.IsLabel),
					AtUsers:   msg.AtUsers,
					AtAllUser: v1.AtAllUserType(msg.AtAllUser),
				})
			}
		}
		resp.Total = total
		return resp, nil
	}
	list, err := s.mr.GetGroupMsgList(dataTransformers.GroupMsgList{
		DialogId:   in.DialogId,
		Content:    in.Content,
		UserID:     in.UserId,
		MsgType:    entity.UserMessageType(in.Type),
		PageNumber: int(in.PageNum),
		PageSize:   int(in.PageSize),
	})
	if err != nil {
		return nil, status.Error(codes.Code(code.MsgErrGetGroupMsgListFailed.Code()), err.Error())
	}

	resp.Total = list.Total
	resp.CurrentPage = list.CurrentPage

	if len(list.GroupMessages) > 0 {
		for _, msg := range list.GroupMessages {
			resp.GroupMessages = append(resp.GroupMessages, &v1.GroupMessage{
				Id:        uint32(msg.ID),
				UserId:    msg.UserID,
				Content:   msg.Content,
				Type:      uint32(int32(msg.Type)),
				ReplyId:   uint32(msg.ReplyId),
				GroupId:   uint32(msg.GroupID),
				ReadCount: int32(msg.ReadCount),
				CreatedAt: msg.CreatedAt,
				DialogId:  uint32(msg.DialogId),
				IsLabel:   v1.MsgLabel(msg.IsLabel),
				AtUsers:   msg.AtUsers,
				AtAllUser: v1.AtAllUserType(msg.AtAllUser),
			})
		}
	}

	return resp, nil
}

func (s *Handler) GetGroupMessagesByIds(ctx context.Context, in *v1.GetGroupMessagesByIdsRequest) (*v1.GetGroupMessagesByIdsResponse, error) {
	resp := &v1.GetGroupMessagesByIdsResponse{}
	msgs, err := s.mr.GetGroupMsgsByIDs(in.MsgIds)
	if err != nil {
		return nil, status.Error(codes.Code(code.GetMsgErrGetGroupMsgByIDFailed.Code()), err.Error())
	}
	if len(msgs) > 0 {
		for _, msg := range msgs {
			resp.GroupMessages = append(resp.GroupMessages, &v1.GroupMessage{
				Id:                     uint32(msg.ID),
				UserId:                 msg.UserID,
				Content:                msg.Content,
				Type:                   uint32(int32(msg.Type)),
				ReplyId:                uint32(msg.ReplyId),
				GroupId:                uint32(msg.GroupID),
				ReadCount:              int32(msg.ReadCount),
				DialogId:               uint32(msg.DialogId),
				IsLabel:                v1.MsgLabel(msg.IsLabel),
				IsBurnAfterReadingType: v1.BurnAfterReadingType(msg.IsBurnAfterReading),
				AtUsers:                msg.AtUsers,
				AtAllUser:              v1.AtAllUserType(msg.AtAllUser),
				CreatedAt:              msg.CreatedAt,
			})
		}
	}
	return resp, nil
}

func (s *Handler) GetGroupUnreadMessages(ctx context.Context, in *v1.GetGroupUnreadMessagesRequest) (*v1.GetGroupUnreadMessagesResponse, error) {
	resp := &v1.GetGroupUnreadMessagesResponse{}
	//获取群聊消息id
	//去除不需要的消息类型
	ids, err := s.mr.GetGroupMsgIdsByDialogID(in.DialogId)
	if err != nil {
		return nil, err
	}

	//获取已读消息id
	rids, err := s.gmrr.GetGroupMsgUserReadIdsByDialogID(in.DialogId, in.UserId)
	if err != nil {
		return nil, err
	}

	//求差集
	msgIds := utils.SliceDifference(ids, rids)

	msgs, err := s.mr.GetGroupMsgsByIDs(msgIds)
	if err != nil {
		return nil, err
	}
	if len(msgs) > 0 {
		for _, msg := range msgs {
			resp.GroupMessages = append(resp.GroupMessages, &v1.GroupMessage{
				Id:                     uint32(msg.ID),
				UserId:                 msg.UserID,
				Content:                msg.Content,
				Type:                   uint32(msg.Type),
				ReplyId:                uint32(msg.ReplyId),
				GroupId:                uint32(msg.GroupID),
				ReadCount:              int32(msg.ReadCount),
				DialogId:               uint32(msg.DialogId),
				IsLabel:                v1.MsgLabel(msg.IsLabel),
				IsBurnAfterReadingType: v1.BurnAfterReadingType(msg.IsBurnAfterReading),
				AtUsers:                msg.AtUsers,
				AtAllUser:              v1.AtAllUserType(msg.AtAllUser),
				CreatedAt:              msg.CreatedAt,
			})
		}
	}
	//查询指定消息
	return resp, nil
}

func (s *Handler) GetUserMessagesByIds(ctx context.Context, in *v1.GetUserMessagesByIdsRequest) (*v1.GetUserMessagesByIdsResponse, error) {
	resp := &v1.GetUserMessagesByIdsResponse{}
	msgs, err := s.mr.GetUserMsgByIDs(in.MsgIds)
	if err != nil {
		return nil, status.Error(codes.Code(code.GetMsgErrGetUserMsgByIDFailed.Code()), err.Error())
	}
	if len(msgs) > 0 {
		for _, msg := range msgs {
			resp.UserMessages = append(resp.UserMessages, &v1.UserMessage{
				Id:                     uint32(msg.ID),
				SenderId:               msg.SendID,
				ReceiverId:             msg.ReceiveID,
				ReadAt:                 msg.ReadAt,
				IsRead:                 int32(msg.IsRead),
				Content:                msg.Content,
				Type:                   uint32(msg.Type),
				ReplyId:                uint64(msg.ReplyId),
				DialogId:               uint32(msg.DialogId),
				IsLabel:                v1.MsgLabel(msg.IsLabel),
				IsBurnAfterReadingType: v1.BurnAfterReadingType(msg.IsBurnAfterReading),
				CreatedAt:              msg.CreatedAt,
			})
		}
	}
	return resp, nil
}

func (s *Handler) SendMultiUserMessage(ctx context.Context, in *v1.SendMultiUserMsgRequest) (*v1.SendMultiUserMsgResponse, error) {
	resp := &v1.SendMultiUserMsgResponse{}
	if len(in.MsgList) > 0 {
		list := make([]entity.UserMessage, 0)
		for _, msg := range in.MsgList {
			list = append(list, entity.UserMessage{
				SendID:    msg.SenderId,
				ReceiveID: msg.ReceiverId,
				Content:   msg.Content,
				Type:      entity.UserMessageType(msg.Type),
				//ReplyId:            uint(msg.ReplyId),
				DialogId: uint(msg.DialogId),
				//IsBurnAfterReading: entity.BurnAfterReadingType(msg.IsBurnAfterReadingType),
			})
		}
		err := s.mr.InsertUserMessages(list)
		if err != nil {
			return nil, status.Error(codes.Code(code.MsgErrSendMultipleFailed.Code()), err.Error())
		}
	}
	return resp, nil
}

func (s *Handler) DeleteUserMessageByDialogId(ctx context.Context, in *v1.DeleteUserMsgByDialogIdRequest) (*v1.DeleteUserMsgByDialogIdResponse, error) {
	resp := &v1.DeleteUserMsgByDialogIdResponse{}
	fmt.Println("尝试删除")
	err := s.mr.DeleteUserMessagesByDialogID(in.DialogId)
	if err != nil {
		return resp, status.Error(codes.Aborted, fmt.Sprintf("failed to delete user msg: %v", err))
	}
	return resp, err
}

func (s *Handler) DeleteGroupMessageByDialogId(ctx context.Context, in *v1.DeleteGroupMsgByDialogIdRequest) (*v1.DeleteGroupMsgByDialogIdResponse, error) {
	resp := &v1.DeleteGroupMsgByDialogIdResponse{}
	err := s.mr.DeleteGroupMessagesByDialogID(in.DialogId)
	if err != nil {
		return resp, status.Error(codes.Aborted, fmt.Sprintf("failed to delete group msg: %v", err))
	}
	return resp, err
}

func (s *Handler) ConfirmDeleteUserMessageByDialogId(ctx context.Context, in *v1.DeleteUserMsgByDialogIdRequest) (*v1.DeleteUserMsgByDialogIdResponse, error) {
	resp := &v1.DeleteUserMsgByDialogIdResponse{}
	err := s.mr.PhysicalDeleteUserMessagesByDialogID(in.DialogId)
	if err != nil {
		return resp, status.Error(codes.Aborted, fmt.Sprintf("failed to delete user msg: %v", err))
	}
	return resp, err
}

func (s *Handler) ConfirmDeleteGroupMessageByDialogId(ctx context.Context, in *v1.DeleteGroupMsgByDialogIdRequest) (*v1.DeleteGroupMsgByDialogIdResponse, error) {
	resp := &v1.DeleteGroupMsgByDialogIdResponse{}
	err := s.mr.PhysicalDeleteGroupMessagesByDialogID(in.DialogId)
	if err != nil {
		return resp, status.Error(codes.Aborted, fmt.Sprintf("failed to delete group msg: %v", err))
	}
	return resp, err
}

func (s *Handler) DeleteUserMessageByDialogIdRollback(ctx context.Context, in *v1.DeleteUserMsgByDialogIdRequest) (*v1.DeleteUserMsgByDialogIdResponse, error) {
	resp := &v1.DeleteUserMsgByDialogIdResponse{}
	err := s.mr.UpdateUserMsgColumnByDialogId(in.DialogId, "deleted_at", 0)
	if err != nil {
		return resp, status.Error(codes.Code(code.MsgErrDeleteUserMessageFailed.Code()), err.Error())
	}
	return resp, err
}

func (s *Handler) DeleteGroupMessageByDialogIdRollback(ctx context.Context, in *v1.DeleteUserMsgByDialogIdRequest) (*v1.DeleteGroupMsgByDialogIdResponse, error) {
	resp := &v1.DeleteGroupMsgByDialogIdResponse{}
	err := s.mr.UpdateGroupMsgColumnByDialogId(in.DialogId, "deleted_at", 0)
	if err != nil {
		return resp, status.Error(codes.Code(code.MsgErrDeleteGroupMessageFailed.Code()), err.Error())
	}
	return resp, err
}

func (s *Handler) GetUserLastMessageList(ctx context.Context, request *v1.GetLastMsgListRequest) (*v1.UserMessages, error) {
	resp := &v1.UserMessages{}
	msgs, total, err := s.mr.GetUserDialogLastMsgs(request.DialogId, int(request.PageNum), int(request.PageSize))
	if err != nil {
		return nil, status.Error(codes.Code(code.GetMsgErrGetUserMsgByIDFailed.Code()), err.Error())
	}
	if len(msgs) > 0 {
		for _, msg := range msgs {
			resp.UserMessages = append(resp.UserMessages, &v1.UserMessage{
				Id:                     uint32(msg.ID),
				SenderId:               msg.SendID,
				ReceiverId:             msg.ReceiveID,
				ReadAt:                 msg.ReadAt,
				IsRead:                 int32(msg.IsRead),
				Content:                msg.Content,
				Type:                   uint32(msg.Type),
				ReplyId:                uint64(msg.ReplyId),
				DialogId:               uint32(msg.DialogId),
				IsLabel:                v1.MsgLabel(msg.IsLabel),
				IsBurnAfterReadingType: v1.BurnAfterReadingType(msg.IsBurnAfterReading),
				CreatedAt:              msg.CreatedAt,
			})
		}
	}
	resp.Total = uint64(total)
	return resp, nil
}

func (s *Handler) GetGroupLastMessageList(ctx context.Context, request *v1.GetLastMsgListRequest) (*v1.GroupMessages, error) {
	resp := &v1.GroupMessages{}
	msgs, total, err := s.mr.GetGroupDialogLastMsgs(request.DialogId, int(request.PageNum), int(request.PageSize))
	if err != nil {
		return nil, status.Error(codes.Code(code.GetMsgErrGetUserMsgByIDFailed.Code()), err.Error())
	}
	if len(msgs) > 0 {
		for _, msg := range msgs {
			resp.GroupMessages = append(resp.GroupMessages, &v1.GroupMessage{
				Id:                     uint32(msg.ID),
				UserId:                 msg.UserID,
				Content:                msg.Content,
				Type:                   uint32(msg.Type),
				ReplyId:                uint32(msg.ReplyId),
				GroupId:                uint32(msg.GroupID),
				ReadCount:              int32(msg.ReadCount),
				DialogId:               uint32(msg.DialogId),
				IsLabel:                v1.MsgLabel(msg.IsLabel),
				IsBurnAfterReadingType: v1.BurnAfterReadingType(msg.IsBurnAfterReading),
				AtUsers:                msg.AtUsers,
				AtAllUser:              v1.AtAllUserType(msg.AtAllUser),
				CreatedAt:              msg.CreatedAt,
			})
		}
	}
	resp.Total = uint64(total)
	return resp, nil
}

func (s *Handler) ReadAllUserMsg(ctx context.Context, request *v1.ReadAllUserMsgRequest) (*v1.ReadAllUserMsgResponse, error) {
	resp := &v1.ReadAllUserMsgResponse{}
	if err := s.mr.ReadAllUserMsg(request.UserId, request.DialogId); err != nil {
		return nil, status.Error(codes.Code(code.SetMsgErrSetUserMsgsReadStatusFailed.Code()), err.Error())
	}
	return resp, nil
}

func (s *Handler) DeleteUserMessageById(ctx context.Context, request *v1.DeleteUserMsgByIDRequest) (*v1.DeleteUserMsgByIDResponse, error) {
	resp := &v1.DeleteUserMsgByIDResponse{}
	if err := s.mr.PhysicalDeleteUserMessage(request.ID); err != nil {
		return resp, status.Error(codes.Code(code.MsgErrDeleteUserMessageFailed.Code()), err.Error())
	}
	return resp, nil
}
