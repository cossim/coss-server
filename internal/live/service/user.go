package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cossim/coss-server/internal/live/api/dto"
	"github.com/cossim/coss-server/internal/live/api/model"
	msggrpcv1 "github.com/cossim/coss-server/internal/msg/api/grpc/v1"
	pushgrpcv1 "github.com/cossim/coss-server/internal/push/api/grpc/v1"
	relationgrpcv1 "github.com/cossim/coss-server/internal/relation/api/grpc/v1"
	usergrpcv1 "github.com/cossim/coss-server/internal/user/api/grpc/v1"
	"github.com/cossim/coss-server/pkg/code"
	"github.com/cossim/coss-server/pkg/constants"
	"github.com/cossim/coss-server/pkg/utils"
	pkgtime "github.com/cossim/coss-server/pkg/utils/time"
	any2 "github.com/golang/protobuf/ptypes/any"
	"github.com/google/uuid"
	"github.com/livekit/protocol/auth"
	"github.com/livekit/protocol/livekit"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"strings"
	"time"
)

func (s *Service) CreateUserCall(ctx context.Context, senderID string, req *dto.UserCallRequest) (*dto.UserCallResponse, error) {
	recipientID := req.UserID
	inCall, err := s.isEitherUserInCall(ctx, senderID, recipientID)
	if err != nil {
		s.logger.Error("获取通话状态失败", zap.Error(err))
		return nil, code.LiveErrCreateCallFailed
	}
	if inCall {
		return nil, code.LiveErrAlreadyInCall
	}

	// Check sender's relationship with recipient
	if err := s.checkUserRelation(ctx, senderID, recipientID); err != nil {
		return nil, err
	}

	// Check recipient's relationship with sender
	if err := s.checkUserRelation(ctx, recipientID, senderID); err != nil {
		return nil, err
	}

	roomName := uuid.New().String()
	_, err = s.roomService.CreateRoom(ctx, &livekit.CreateRoomRequest{
		Name:            roomName,
		EmptyTimeout:    uint32(s.liveTimeout.Seconds()), // 空闲时间
		MaxParticipants: 2,
	})
	if err != nil {
		s.logger.Error("创建通话失败", zap.Error(err))
		return nil, err
	}

	userRoom, err := (&model.Room{
		Room: roomName,
	}).ToJSONString()
	if err != nil {
		return nil, err
	}
	if err = s.redisClient.SetKey(liveUserPrefix+senderID, userRoom, s.liveTimeout); err != nil {
		s.logger.Error("保存用户通话记录失败", zap.Error(err), zap.String("uid", senderID))
		return nil, err
	}
	if err = s.redisClient.SetKey(liveUserPrefix+recipientID, userRoom, s.liveTimeout); err != nil {
		s.logger.Error("保存用户通话记录失败", zap.Error(err), zap.String("uid", recipientID))
		return nil, err
	}

	roomInfo, err := (&model.RoomInfo{
		Room:            roomName,
		SenderID:        senderID,
		Type:            model.UserRoomType,
		Option:          req.Option,
		MaxParticipants: 2,
		Participants: map[string]*model.ActiveParticipant{
			senderID: {
				Connected: false,
			},
			recipientID: {
				Connected: false,
			},
		},
	}).ToJSONString()
	if err != nil {
		return nil, err
	}
	if err = s.redisClient.SetKey(liveRoomPrefix+roomName, roomInfo, s.liveTimeout); err != nil {
		s.logger.Error("保存房间信息失败", zap.Error(err), zap.String("room", roomName))
		return nil, err
	}

	fmt.Println("s.liveTimeout+10 => ", s.liveTimeout)
	data := map[string]interface{}{
		"url":          s.livekitServer,
		"sender_id":    senderID,
		"recipient_id": recipientID,
		"option":       req.Option,
	}

	bytes, err := utils.StructToBytes(data)
	if err != nil {
		return nil, err
	}

	msg := &pushgrpcv1.WsMsg{
		Uid:   recipientID,
		Event: pushgrpcv1.WSEventType_UserCallReqEvent,
		Data:  &any2.Any{Value: bytes}}

	toBytes, err := utils.StructToBytes(msg)
	if err != nil {
		return nil, err
	}
	_, err = s.pushService.Push(ctx, &pushgrpcv1.PushRequest{
		Type: pushgrpcv1.Type_Ws,
		Data: toBytes,
	})
	if err != nil {
		s.logger.Error("发送消息失败", zap.Error(err))
	}
	return &dto.UserCallResponse{
		Url: s.livekitServer,
	}, nil
}

func (s *Service) UserJoinRoom(ctx context.Context, uid, driverId string) (*dto.UserJoinResponse, error) {
	s.logger.Debug("用户开始加入房间", zap.String("uid", uid), zap.String("driverId", driverId))
	room, err := s.getRedisUserRoom(ctx, uid)
	if err != nil {
		s.logger.Error("获取房间信息失败", zap.Error(err))
		return nil, err
	}

	if err = s.checkUserRoom(ctx, room, uid); err != nil {
		return nil, err
	}

	room.NumParticipants++
	if room.NumParticipants > room.MaxParticipants {
		return nil, code.LiveErrMaxParticipantsExceeded
	}

	room.Participants[uid] = &model.ActiveParticipant{
		Connected: true,
	}
	roomStr, err := room.ToJSONString()
	if err != nil {
		s.logger.Error("更新用户房间记录失败", zap.Error(err))
		return nil, code.LiveErrJoinCallFailed
	}

	if room.NumParticipants == 2 {
		for k := range room.Participants {
			if err = s.redisClient.PersistKey(liveUserPrefix + k); err != nil {
				s.logger.Error("更新用户房间记录失败", zap.Error(err))
				return nil, err
			}
		}
		if err = s.redisClient.SetKey(liveRoomPrefix+room.Room, roomStr, 0); err != nil {
			s.logger.Error("更新房间信息失败", zap.Error(err), zap.String("room", room.Room))
			return nil, err
		}
	} else {
		if err = s.redisClient.UpdateKey(liveRoomPrefix+room.Room, roomStr, -1); err != nil {
			s.logger.Error("更新房间信息失败", zap.Error(err), zap.String("room", room.Room))
			return nil, err
		}
	}

	user, err := s.userService.UserInfo(ctx, &usergrpcv1.UserInfoRequest{UserId: uid})
	if err != nil {
		return nil, err
	}

	token, err := s.GetAdminJoinToken(ctx, room.Room, user.NickName, uid)
	if err != nil {
		return nil, code.LiveErrJoinCallFailed
	}

	data := map[string]interface{}{
		"room":         room.Room,
		"sender_id":    room.SenderID,
		"recipient_id": uid,
	}
	bytes, err := utils.StructToBytes(data)
	if err != nil {
		return nil, err
	}
	for k := range room.Participants {
		if k == uid {
			continue
		}

		msg := &pushgrpcv1.WsMsg{Uid: k, Event: pushgrpcv1.WSEventType_UserCallAcceptEvent, Data: &any2.Any{Value: bytes}}

		toBytes, err := utils.StructToBytes(msg)
		if err != nil {
			return nil, err
		}
		//if err = s.publishServiceMessage(ctx, msg); err != nil {
		//	s.logger.Error("推送用户通话接受事件失败", zap.Error(err))
		//}
		_, err = s.pushService.Push(ctx, &pushgrpcv1.PushRequest{
			Type: pushgrpcv1.Type_Ws,
			Data: toBytes,
		})
		if err != nil {
			s.logger.Error("发送消息失败", zap.Error(err))
		}
	}

	s.logger.Info("用户加入房间", zap.String("uid", uid), zap.String("room", roomStr), zap.String("livekit", s.livekitServer))

	return &dto.UserJoinResponse{
		Url:   s.livekitServer,
		Token: token,
	}, nil
}

func (s *Service) GetUserRoom(ctx context.Context, uid string) (*dto.UserShowResponse, error) {
	room, err := s.getRedisUserRoom(ctx, uid)
	if err != nil {
		s.logger.Error("获取用户房间信息失败", zap.Error(err), zap.String("uid", uid))
		return nil, err
	}

	if _, ok := room.Participants[uid]; !ok {
		return nil, code.Forbidden
	}

	livekitRoom, err := s.getLivekitRoom(ctx, room.Room)
	if err != nil {
		if err := s.deleteUserRoom(ctx, room); err != nil {
			s.logger.Error("删除群聊房间信息失败", zap.Error(err))
			return nil, code.LiveErrGetCallInfoFailed
		}
	}

	redisRoom, err := s.getRedisRoom(ctx, liveUserPrefix+uid)
	if err != nil {
		return nil, err
	}

	resp := &dto.UserShowResponse{
		StartAt:     livekitRoom.CreationTime,
		Duration:    0,
		Room:        livekitRoom.Name,
		Type:        string(redisRoom.Type),
		Participant: make([]*dto.ParticipantInfo, 0),
	}
	res, err := s.roomService.ListParticipants(ctx, &livekit.ListParticipantsRequest{
		Room: room.Room,
	})
	if err != nil {
		s.logger.Error("获取通话信息失败", zap.Error(err))
		return nil, code.LiveErrGetCallInfoFailed
	}

	for _, p := range res.Participants {
		resp.Participant = append(resp.Participant, &dto.ParticipantInfo{
			Sid:         p.Sid,
			Identity:    p.Identity,
			State:       dto.ParticipantState(p.State),
			JoinedAt:    p.JoinedAt,
			Name:        p.Name,
			IsPublisher: p.IsPublisher,
		})
	}

	return resp, nil
}

func (s *Service) UserRejectRoom(ctx context.Context, uid string, driverId string) (interface{}, error) {
	room, err := s.getRedisUserRoom(ctx, uid)
	if err != nil {
		return nil, err
	}

	if _, ok := room.Participants[uid]; !ok {
		return nil, code.Forbidden
	}

	rel, err := s.relationUserService.GetUserRelation(ctx, &relationgrpcv1.GetUserRelationRequest{UserId: uid, FriendId: room.SenderID})
	if err != nil {
		return nil, err
	}

	if rel.Status != relationgrpcv1.RelationStatus_RELATION_NORMAL {
		return nil, code.RelationUserErrFriendRelationNotFound
	}

	var senderID = room.SenderID
	var receiverId string
	for k, _ := range room.Participants {
		if k != senderID {
			receiverId = k
			break
		}
	}

	var msgType dto.UserMessageType
	if room.Option.AudioEnabled {
		msgType = dto.MessageTypeVoiceCall
	}
	if room.Option.VideoEnabled {
		msgType = dto.MessageTypeVideoCall
	}

	isBurnAfterReading := rel.OpenBurnAfterReading
	OpenBurnAfterReadingTimeOut := rel.OpenBurnAfterReadingTimeOut
	message, err := s.msgService.SendUserMessage(ctx, &msggrpcv1.SendUserMsgRequest{
		DialogId:               rel.DialogId,
		SenderId:               senderID,
		ReceiverId:             receiverId,
		Content:                "已拒绝",
		Type:                   msgType.Int32(),
		IsBurnAfterReadingType: msggrpcv1.BurnAfterReadingType(isBurnAfterReading),
	})
	if err != nil {
		s.logger.Error("发送消息失败", zap.Error(err))
		return nil, code.LiveErrLeaveCallFailed
	}

	info, err := s.userService.UserInfo(ctx, &usergrpcv1.UserInfoRequest{
		UserId: senderID,
	})
	if err != nil {
		s.logger.Error("获取用户信息失败", zap.Error(err))
		return nil, code.LiveErrLeaveCallFailed
	}
	data := &constants.WsUserMsg{
		SenderId:                senderID,
		Content:                 "已拒绝",
		MsgType:                 msgType.Uint(),
		MsgId:                   message.MsgId,
		ReceiverId:              receiverId,
		SendAt:                  pkgtime.Now(),
		DialogId:                rel.DialogId,
		IsBurnAfterReading:      constants.BurnAfterReadingType(isBurnAfterReading),
		BurnAfterReadingTimeOut: OpenBurnAfterReadingTimeOut,
		SenderInfo: constants.SenderInfo{
			Avatar: info.Avatar,
			Name:   info.NickName,
			UserId: senderID,
		},
	}
	bytes, err := utils.StructToBytes(data)
	if err != nil {
		return nil, err
	}
	var msgs []*pushgrpcv1.WsMsg
	for k, _ := range room.Participants {
		msgContent := &pushgrpcv1.WsMsg{
			Uid:      k,
			Event:    pushgrpcv1.WSEventType_SendUserMessageEvent,
			DriverId: driverId,
			Data:     &any2.Any{Value: bytes},
		}
		msgs = append(msgs, msgContent)

		//if err := s.publishServiceMessage(ctx, msgContent); err != nil {
		//	s.logger.Error("发送消息失败", zap.Error(err))
		//}
	}
	toBytes, err := utils.StructToBytes(msgs)
	if err != nil {
		return nil, err
	}
	_, err = s.pushService.Push(ctx, &pushgrpcv1.PushRequest{
		Type: pushgrpcv1.Type_Ws_Batch_User,
		Data: toBytes,
	})
	if err != nil {
		s.logger.Error("发送消息失败", zap.Error(err))
	}

	if err := s.deleteUserRoom(ctx, room); err != nil {
		return nil, err
	}
	data2 := map[string]interface{}{
		"url":          s.livekitServer,
		"sender_id":    senderID,
		"recipient_id": uid,
	}
	structToBytes, err := utils.StructToBytes(data2)
	if err != nil {
		return nil, err
	}
	msg := &pushgrpcv1.WsMsg{Uid: senderID, Event: pushgrpcv1.WSEventType_UserCallRejectEvent, Data: &any2.Any{Value: structToBytes}}
	toBytes3, err := utils.StructToBytes(msg)
	if err != nil {
		return nil, err
	}
	_, err = s.pushService.Push(ctx, &pushgrpcv1.PushRequest{
		Type: pushgrpcv1.Type_Ws,
		Data: toBytes3,
	})
	if err != nil {
		s.logger.Error("发送消息失败", zap.Error(err))
	}
	//if err = s.publishServiceMessage(ctx, msg); err != nil {
	//	s.logger.Error("发送消息失败", zap.Error(err))
	//}

	s.logger.Info("UserRejectRoom", zap.String("room", room.Room), zap.String("SenderID", senderID), zap.String("RecipientID", uid))

	return nil, nil
}

func (s *Service) UserLeaveRoom(ctx context.Context, uid, driverId string) (interface{}, error) {
	room, err := s.getRedisUserRoom(ctx, uid)
	if err != nil {
		if strings.Contains(err.Error(), "redis: nil") {
			fmt.Println("UserLeaveRoom 1 err => ", err)
			return nil, code.LiveErrCallNotFound
		}
		fmt.Println("UserLeaveRoom 2 err => ", err)
		return nil, err
	}

	if _, ok := room.Participants[uid]; !ok {
		return nil, code.Forbidden
	}

	livekitRoom, err := s.getLivekitRoom(ctx, room.Room)
	if err != nil {
		if room != nil && errors.Is(err, code.LiveErrCallNotFound) {
			if err := s.deleteUserRoom(ctx, room); err != nil {
				s.logger.Error("删除群聊房间信息失败", zap.Error(err))
			}
		}
		return nil, fmt.Errorf("failed to get Livekit room: %w", err)
	}

	formatDuration := func(duration time.Duration) string {
		minutes := int(duration.Minutes())
		seconds := int(duration.Seconds()) % 60
		return fmt.Sprintf("%02d:%02d", minutes, seconds)
	}

	creationTime := time.Unix(livekitRoom.CreationTime, 0)
	duration := time.Since(creationTime)
	content := fmt.Sprintf("通话时长：%s", formatDuration(duration))

	var (
		senderID                    string
		receiverId                  string
		did                         uint32
		isBurnAfterReading          relationgrpcv1.OpenBurnAfterReadingType
		OpenBurnAfterReadingTimeOut int64
	)

	// Find receiver ID
	for k := range room.Participants {
		if k == room.SenderID {
			senderID = k
		} else {
			receiverId = k
		}
	}

	s.logger.Info("UserLeaveRoom", zap.String("room", room.Room), zap.String("挂断用户", uid), zap.String("被挂断用户", receiverId), zap.String("通话时长", content), zap.String("创建者", senderID))
	if err = s.deleteUserRoom(ctx, room); err != nil {
		return nil, err
	}

	rel1, err := s.relationUserService.GetUserRelation(ctx, &relationgrpcv1.GetUserRelationRequest{UserId: senderID, FriendId: receiverId})
	if err != nil {
		return nil, err
	}
	if rel1.Status != relationgrpcv1.RelationStatus_RELATION_NORMAL {
		return nil, code.RelationUserErrFriendRelationNotFound
	}

	// 获取对话ID、烧毁模式和烧毁超时时间
	did = rel1.DialogId
	isBurnAfterReading = rel1.OpenBurnAfterReading
	OpenBurnAfterReadingTimeOut = rel1.OpenBurnAfterReadingTimeOut

	rel2, err := s.relationUserService.GetUserRelation(ctx, &relationgrpcv1.GetUserRelationRequest{UserId: receiverId, FriendId: senderID})
	if err != nil {
		return nil, err
	}
	if rel2.Status != relationgrpcv1.RelationStatus_RELATION_NORMAL {
		return nil, code.RelationUserErrFriendRelationNotFound
	}

	var msgType dto.UserMessageType

	if room.Option.AudioEnabled {
		msgType = dto.MessageTypeVoiceCall
	}

	if room.Option.VideoEnabled {
		msgType = dto.MessageTypeVideoCall
	}

	message, err := s.msgService.SendUserMessage(ctx, &msggrpcv1.SendUserMsgRequest{
		DialogId:               did,
		SenderId:               senderID,
		ReceiverId:             receiverId,
		Content:                content,
		Type:                   msgType.Int32(),
		IsBurnAfterReadingType: msggrpcv1.BurnAfterReadingType(isBurnAfterReading),
	})
	if err != nil {
		s.logger.Error("发送消息失败", zap.Error(err))
		return nil, code.LiveErrLeaveCallFailed
	}

	info, err := s.userService.UserInfo(ctx, &usergrpcv1.UserInfoRequest{
		UserId: senderID,
	})
	if err != nil {
		s.logger.Error("获取用户信息失败", zap.Error(err))
		return nil, code.LiveErrLeaveCallFailed
	}
	data := &constants.WsUserMsg{
		SenderId:                senderID,
		Content:                 content,
		MsgType:                 msgType.Uint(),
		MsgId:                   message.MsgId,
		ReceiverId:              receiverId,
		SendAt:                  pkgtime.Now(),
		DialogId:                did,
		IsBurnAfterReading:      constants.BurnAfterReadingType(isBurnAfterReading),
		BurnAfterReadingTimeOut: OpenBurnAfterReadingTimeOut,
		SenderInfo: constants.SenderInfo{
			Avatar: info.Avatar,
			Name:   info.NickName,
			UserId: senderID,
		},
	}

	bytes, err := utils.StructToBytes(data)
	if err != nil {
		return nil, err
	}

	for k := range room.Participants {
		msgContent := &pushgrpcv1.WsMsg{
			Uid:      k,
			Event:    pushgrpcv1.WSEventType_SendUserMessageEvent,
			DriverId: driverId,
			Data:     &any2.Any{Value: bytes},
		}
		toBytes, err := utils.StructToBytes(msgContent)
		if err != nil {
			return nil, err
		}

		_, err = s.pushService.Push(ctx, &pushgrpcv1.PushRequest{
			Type: pushgrpcv1.Type_Ws,
			Data: toBytes,
		})
		if err != nil {
			s.logger.Error("发送消息失败", zap.Error(err))
		}
		//if err := s.publishServiceMessage(ctx, msgContent); err != nil {
		//	s.logger.Error("发送消息失败", zap.Error(err))
		//}

		// 发送用户结束通话事件消息
		data2 := map[string]interface{}{
			"room":         room.Room,
			"sender_id":    senderID,
			"recipient_id": receiverId,
		}
		structToBytes, err := utils.StructToBytes(data2)
		if err != nil {
			return nil, err
		}
		endCallMsg := &pushgrpcv1.WsMsg{
			Uid:   k,
			Event: pushgrpcv1.WSEventType_UserCallEndEvent,
			Data:  &any2.Any{Value: structToBytes},
		}
		toBytes2, err := utils.StructToBytes(endCallMsg)
		if err != nil {
			return nil, err
		}

		_, err = s.pushService.Push(ctx, &pushgrpcv1.PushRequest{
			Type: pushgrpcv1.Type_Ws,
			Data: toBytes2,
		})
		if err != nil {
			s.logger.Error("发送消息失败", zap.Error(err))
		}

	}

	return nil, nil
}

// checkUserRelation checks the relationship status between two users
func (s *Service) checkUserRelation(ctx context.Context, userID, friendID string) error {
	rel, err := s.relationUserService.GetUserRelation(ctx, &relationgrpcv1.GetUserRelationRequest{
		UserId:   userID,
		FriendId: friendID,
	})
	if err != nil {
		s.logger.Error("Failed to get user relation", zap.Error(err))
		return err
	}

	if rel.Status != relationgrpcv1.RelationStatus_RELATION_NORMAL {
		return code.RelationUserErrFriendRelationNotFound
	}

	return nil
}

func (s *Service) checkUserRoom(ctx context.Context, roomInfo *model.RoomInfo, uid string) error {
	if _, ok := roomInfo.Participants[uid]; !ok {
		return code.Forbidden
	}

	if _, ok := roomInfo.Participants[uid]; ok && roomInfo.Participants[uid].Connected {
		return code.LiveErrAlreadyInCall
	}

	if roomInfo.NumParticipants+1 > roomInfo.MaxParticipants {
		return code.LiveErrMaxParticipantsExceeded
	}

	room, err := s.getLivekitRoom(ctx, roomInfo.Room)
	if err != nil {
		return err
	}

	if room.NumParticipants+1 > room.MaxParticipants {
		return code.LiveErrMaxParticipantsExceeded
	}

	return nil
}

func (s *Service) joinRoom(userName string, room *livekit.Room) (string, error) {
	token := s.roomService.CreateToken().SetIdentity(userName).AddGrant(&auth.VideoGrant{
		Room:     room.Name,
		RoomJoin: true,
	})

	return token.ToJWT()
}

func (s *Service) getLivekitRoom(ctx context.Context, room string) (*livekit.Room, error) {
	rooms, err := s.roomService.ListRooms(ctx, &livekit.ListRoomsRequest{Names: []string{room}})
	if err != nil {
		s.logger.Error("获取房间信息失败", zap.Error(err))
		return nil, code.LiveErrGetCallInfoFailed
	}

	if len(rooms.Rooms) == 0 {
		return nil, code.LiveErrCallNotFound
	}

	return rooms.Rooms[0], nil
}

func (s *Service) getRedisRoom(ctx context.Context, uid string) (*model.RoomInfo, error) {
	room, err := s.redisClient.GetKey(uid)
	if err != nil {
		return nil, err
	}
	d1 := &model.Room{}
	if err = json.Unmarshal([]byte(room.(string)), &d1); err != nil {
		return nil, err
	}

	roomInfo, err := s.redisClient.GetKey(liveRoomPrefix + d1.Room)
	if err != nil {
		return nil, err
	}
	d2 := &model.RoomInfo{}
	if err = json.Unmarshal([]byte(roomInfo.(string)), &d2); err != nil {
		return nil, err
	}

	return d2, nil
}

func (s *Service) deleteLivekitRoom(ctx context.Context, room string) error {
	_, err := s.roomService.DeleteRoom(ctx, &livekit.DeleteRoomRequest{
		Room: room,
	})
	if err != nil && !strings.Contains(err.Error(), "room not found") {
		s.logger.Error("error deleting room", zap.Error(err))
		return code.LiveErrLeaveCallFailed
	}

	return nil
}

func (s *Service) deleteUserRedisRoom(ctx context.Context, room *model.RoomInfo) error {
	for id, _ := range room.Participants {
		if err := s.redisClient.DelKey(liveUserPrefix + id); err != nil {
			s.logger.Error("删除用户房间信息", zap.Error(err))
		}
	}

	if err := s.redisClient.DelKey(liveRoomPrefix + room.Room); err != nil {
		s.logger.Error("删除房间失败", zap.Error(err), zap.String("room", room.Room))
	}

	return nil
}

func (s *Service) deleteUserRoom(ctx context.Context, room *model.RoomInfo) error {
	if err := s.deleteLivekitRoom(ctx, room.Room); err != nil {
		return err
	}

	if err := s.deleteUserRedisRoom(ctx, room); err != nil {
		return err
	}

	return nil
}

func (s *Service) deleteRedisRoomByPrefix(ctx context.Context, prefix, roomID string) error {
	pattern := prefix + roomID + ":*"
	keys, err := s.redisClient.Client.Keys(ctx, pattern).Result()
	if err != nil {
		return err
	}

	if len(keys) == 0 {
		return nil
	}

	_, err = s.redisClient.Client.Del(ctx, keys...).Result()
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) deleteGroupCallByRoomID(ctx context.Context, roomID string) error {
	pattern := liveGroupPrefix + roomID + ":*"
	keys, err := s.redisClient.Client.Keys(ctx, pattern).Result()
	if err != nil {
		return err
	}

	if len(keys) == 0 {
		return nil
	}

	_, err = s.redisClient.Client.Del(ctx, keys...).Result()
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) isEitherUserInCall(ctx context.Context, userID1, userID2 string) (bool, error) {
	is1, err := s.isUserInCall(ctx, userID1)
	if err != nil {
		return false, err
	}
	if is1 {
		return true, nil
	}

	is2, err := s.isUserInCall(ctx, userID2)
	if err != nil {
		return false, err
	}
	return is2, nil
}

func (s *Service) isUserInCall(ctx context.Context, id string) (bool, error) {
	pattern := liveUserPrefix + id
	count, err := s.redisClient.Client.Exists(ctx, pattern).Result()
	if err != nil {
		s.logger.Error("isUserInCall err", zap.Error(err))
		// 检查错误并且错误消息中不包含 "does not exist" 或 "expired"
		if !strings.Contains(err.Error(), "does not exist") && !strings.Contains(err.Error(), "expired") && errors.Is(err, redis.Nil) {
			return false, err
		}
		return true, nil
	}
	return count > 0, nil
}

func (s *Service) getRedisUserRoom(ctx context.Context, uid string) (*model.RoomInfo, error) {
	room, err := s.redisClient.GetKey(liveUserPrefix + uid)
	if err != nil {
		return nil, code.LiveErrCallNotFound.Reason(err)
	}
	d1 := &model.Room{}
	if err = json.Unmarshal([]byte(room.(string)), &d1); err != nil {
		return nil, err
	}

	roomInfo, err := s.redisClient.GetKey(liveRoomPrefix + d1.Room)
	if err != nil {
		if strings.Contains(err.Error(), "redis: nil") {
			if err := s.redisClient.DelKey(liveUserPrefix + uid); err != nil {
				s.logger.Error("获取用户通话信息失败后删除房间失败", zap.Error(err))
			}
			return nil, code.LiveErrCallNotFound
		}
		return nil, err
	}
	d2 := &model.RoomInfo{}
	if err = json.Unmarshal([]byte(roomInfo.(string)), &d2); err != nil {
		return nil, err
	}

	return d2, nil
}

// GetRedisObjectWithCondition retrieves an object from Redis cache based on a condition and unmarshals it into the provided output interface.
func (s *Service) getRedisRoomWithPrefix(ctx context.Context, prefix, key string, out interface{}) (string, error) {
	pattern := prefix + key
	s.logger.Info("scanKeysWithPrefixAndContent", zap.String("pattern", pattern))
	keys, err := s.scanKeysWithPrefixAndContent(ctx, pattern)
	if err != nil {
		s.logger.Error("scanKeysWithPrefixAndContent", zap.Error(err))
		return "", err
	}
	if len(keys) > 1 {
		return "", code.LiveErrGetCallInfoFailed
	}
	if len(keys) == 0 {
		return "", code.LiveErrCallNotFound
	}
	room, err := s.redisClient.GetKey(keys[0])
	if err != nil {
		return "", code.LiveErrCallNotFound.Reason(err)
	}
	if err = json.Unmarshal([]byte(room.(string)), &out); err != nil {
		return "", err
	}
	return keys[0], nil
}

func (s *Service) scanKeysWithPrefixAndContent(ctx context.Context, pattern string) ([]string, error) {
	var cursor uint64 = 0
	var keys []string
	for {
		result, nextCursor, err := s.redisClient.Client.Scan(ctx, cursor, pattern, 10).Result()
		if err != nil {
			return nil, err
		}
		keys = append(keys, result...)
		if nextCursor == 0 {
			break
		}
		cursor = nextCursor
	}
	return keys, nil
}

func (s *Service) deleteRedisLiveUser(ctx context.Context, userID string) error {
	pattern := liveUserPrefix + "*:" + userID + ":*"
	keys, err := s.redisClient.Client.Keys(ctx, pattern).Result()
	if err != nil {
		return err
	}
	if len(keys) == 0 {
		return nil
	}
	_, err = s.redisClient.Client.Del(ctx, keys...).Result()
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetUserJoinToken(ctx context.Context, room, userName, userID string) (string, error) {
	at := auth.NewAccessToken(s.liveApiKey, s.liveApiSecret)
	grant := &auth.VideoGrant{
		RoomJoin: true,
		Room:     room,
	}
	at.AddGrant(grant).SetName(userName).SetIdentity(userID).SetValidFor(s.liveTimeout)
	jwt, err := at.ToJWT()
	if err != nil {
		s.logger.Error("Failed to generate JWT", zap.Error(err))
		return "", err
	}

	return jwt, nil
}

func (s *Service) GetAdminJoinToken(ctx context.Context, room, userName, userID string) (string, error) {
	at := auth.NewAccessToken(s.liveApiKey, s.liveApiSecret)
	grant := &auth.VideoGrant{
		RoomCreate: true,
		RoomList:   true,
		RoomAdmin:  true,
		RoomJoin:   true,
		Room:       room,
	}
	at.AddGrant(grant).SetName(userName).SetIdentity(userID).SetValidFor(s.liveTimeout)
	jwt, err := at.ToJWT()
	if err != nil {
		s.logger.Error("Failed to generate JWT", zap.Error(err))
		return "", err
	}

	return jwt, nil
}

//func (s *Service) publishServiceMessage(ctx context.Context, msg pushgrpcv1.WsMsg) error {
//	return s.mqClient.PublishServiceMessage(msg_queue.LiveUserService, msg_queue.MsgService, msg_queue.Service_Exchange, msg_queue.SendMessage, msg)
//}
