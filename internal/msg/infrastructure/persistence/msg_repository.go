package persistence

import (
	"context"
	"fmt"
	"github.com/cossim/coss-server/internal/msg/api/grpc/dataTransformers"
	"github.com/cossim/coss-server/internal/msg/domain/entity"
	"github.com/cossim/coss-server/internal/msg/domain/repository"
	"github.com/cossim/coss-server/pkg/code"
	"github.com/cossim/coss-server/pkg/utils/time"
	"gorm.io/gorm"
)

var _ repository.MsgRepository = &MsgRepo{}

type MsgRepo struct {
	db *gorm.DB
}

func (g *MsgRepo) Find(ctx context.Context, query *entity.UserMsgQuery) (*entity.UserMsgQueryResult, error) {
	var messages []*entity.UserMessage
	result := &entity.UserMsgQueryResult{}

	db := g.db.Model(&entity.UserMessage{})

	// 根据查询条件过滤消息 ID
	if len(query.MsgIds) > 0 {
		db = db.Where("id IN (?)", query.MsgIds)
	}

	// 根据对话 ID 过滤消息
	if len(query.DialogIds) > 0 {
		db = db.Where("dialog_id IN (?)", query.DialogIds)
	}

	at := time.Now()
	if query.EndAt <= 0 {
		query.EndAt = at
	}

	// 根据时间范围过滤消息
	if query.EndAt > 0 {
		if query.EndAt > at {
			return nil, code.MyCustomErrorCode.CustomMessage("endAt must be less than or equal to the current time")
		}
		db = db.Where("created_at BETWEEN ? AND ?", query.StartAt, query.EndAt)
	}

	if query.Content != "" {
		db = db.Where("content LIKE ?", "%"+query.Content+"%")
	}

	if query.SendID != "" {
		db = db.Where("send_id = ?", query.SendID)
	}

	if entity.IsValidMessageType(query.MsgType) {
		db = db.Where("msg_type = ?", query.MsgType)
	}

	// 查询总消息数
	var totalCount int64
	if err := db.Count(&totalCount).Error; err != nil {
		return nil, err
	}

	result.TotalCount = totalCount

	// 查询剩余消息数
	if query.PageSize > 0 && totalCount > query.PageNum*query.PageSize {
		result.Remaining = totalCount - query.PageNum*query.PageSize
	}

	// 查询总页码
	if query.PageSize > 0 {
		result.TotalPages = (totalCount + query.PageSize - 1) / query.PageSize
	}

	// 查询当前页码
	result.CurrentPage = query.PageNum

	// 分页
	if query.PageNum > 0 {
		offset := (query.PageNum - 1) * query.PageSize
		db = db.Offset(int(offset))
	}

	if query.PageSize > 0 {
		db = db.Limit(int(query.PageSize))
	}

	if query.Sort == "" {
		query.Sort = "desc"
		db = db.Order(fmt.Sprintf("created_at %s", query.Sort))
	}

	// 执行查询
	err := db.Find(&messages).Error
	if err != nil {
		return nil, err
	}

	result.Messages = messages
	result.ReturnedCount = int64(len(result.Messages))

	return result, nil
}

func (g *MsgRepo) GetGroupUnreadMsgList(dialogID uint32, msgIds []uint32) ([]*entity.GroupMessage, error) {
	var msgList []*entity.GroupMessage
	query := g.db.Model(&entity.GroupMessage{}).Where("dialog_id = ? AND deleted_at = 0", dialogID)

	if len(msgIds) > 0 {
		query = query.Where("id NOT IN (?)", msgIds)
	}

	err := query.Find(&msgList).Error
	if err != nil {
		return nil, err
	}
	return msgList, nil
}

func (g *MsgRepo) ReadAllUserMsg(uid string, dialogId uint32) error {
	return g.db.Model(&entity.UserMessage{}).
		Where("receive_id = ? AND dialog_id = ? AND deleted_at = 0", uid, dialogId).
		Updates(map[string]interface{}{
			"is_read": entity.IsRead,
			"read_at": time.Now(),
		}).Error
}

func NewMsgRepo(db *gorm.DB) *MsgRepo {
	return &MsgRepo{db: db}
}

func (g *MsgRepo) InsertUserMessage(senderId string, receiverId string, msg string, msgType entity.UserMessageType, subType entity.UserMessageSubType, replyId uint, dialogId uint, isBurnAfterReading entity.BurnAfterReadingType) (*entity.UserMessage, error) {
	content := &entity.UserMessage{
		SendID:             senderId,
		ReceiveID:          receiverId,
		Content:            msg,
		Type:               msgType,
		SubType:            subType,
		ReplyId:            replyId,
		DialogId:           dialogId,
		IsBurnAfterReading: isBurnAfterReading,
	}
	if err := g.db.Save(content).Error; err != nil {
		return nil, err
	}
	return content, nil
}

func (g *MsgRepo) InsertGroupMessage(uid string, groupId uint, msg string, msgType entity.UserMessageType, replyId uint, dialogId uint, isBurnAfterReading entity.BurnAfterReadingType, atUsers []string, atAllUser entity.AtAllUserType) (*entity.GroupMessage, error) {
	content := &entity.GroupMessage{
		UserID:             uid,
		GroupID:            groupId,
		Content:            msg,
		Type:               msgType,
		ReplyId:            replyId,
		DialogId:           dialogId,
		IsBurnAfterReading: isBurnAfterReading,
		AtUsers:            atUsers,
		AtAllUser:          atAllUser,
	}
	if err := g.db.Save(content).Error; err != nil {
		return nil, err
	}
	return content, nil
}

func (g *MsgRepo) GetUserMsgList(dialogId uint32, sendId string, content string, msgType entity.UserMessageType, pageNumber, pageSize int) ([]entity.UserMessage, int32, int32) {
	var results []entity.UserMessage

	query := g.db.Model(&entity.UserMessage{})
	query = query.Where("dialog_id = ? ", dialogId).Order("created_at DESC")

	var count int64
	err := query.Count(&count).Error
	if err != nil {
		return nil, 0, 0
	}
	if content != "" {
		query = query.Where("content LIKE ?", "%"+content+"%")
	}
	if sendId != "" {
		query = query.Where("send_id = ?", sendId)
	}
	if entity.IsValidMessageType(msgType) {
		query = query.Where("msg_type = ?", msgType)
	}

	offset := (pageNumber - 1) * pageSize
	query = query.Offset(offset).Limit(pageSize).Find(&results)

	return results, int32(count), int32(pageNumber)
}

func (g *MsgRepo) GetLastMsgsForUserWithFriends(userID string, friendIDs []string) ([]*entity.UserMessage, error) {
	var userMessages []*entity.UserMessage

	result := g.db.
		Where("(send_id = ? AND receive_id IN (?)) OR (send_id IN (?) AND receive_id = ?)",
			userID, friendIDs, friendIDs, userID).
		Group("receive_id").
		Order("created_at DESC").
		Find(&userMessages)

	if result.Error != nil {
		return nil, result.Error
	}

	return userMessages, nil
}

func (g *MsgRepo) GetLastMsgsForGroupsWithIDs(groupIDs []uint) ([]*entity.GroupMessage, error) {
	var groupMessages []*entity.GroupMessage

	result := g.db.
		Where("group_id IN (?)", groupIDs).
		Group("group_id").
		Order("created_at DESC").
		Find(&groupMessages)

	if result.Error != nil {
		return nil, result.Error
	}

	return groupMessages, nil
}

func (g *MsgRepo) UpdateUserMessage(msg entity.UserMessage) error {
	err := g.db.Model(&msg).Updates(msg).Error
	return err
}

func (g *MsgRepo) UpdateGroupMessage(msg entity.GroupMessage) error {
	err := g.db.Model(&msg).Updates(msg).Error
	return err
}

func (g *MsgRepo) LogicalDeleteUserMessage(msgId uint32) error {
	err := g.db.Model(&entity.UserMessage{}).Where("id = ?", msgId).Update("deleted_at", time.Now()).Error
	return err
}

func (g *MsgRepo) LogicalDeleteGroupMessage(msgId uint32) error {
	err := g.db.Model(&entity.GroupMessage{}).Where("id = ?", msgId).Update("deleted_at", time.Now()).Error
	return err
}

func (g *MsgRepo) PhysicalDeleteUserMessage(msgId uint32) error {
	err := g.db.Delete(&entity.UserMessage{}, msgId).Error
	return err
}

func (g *MsgRepo) PhysicalDeleteGroupMessage(msgId uint32) error {
	err := g.db.Delete(&entity.GroupMessage{}, msgId).Error
	return err
}

func (g *MsgRepo) GetUserMsgByID(msgId uint32) (*entity.UserMessage, error) {
	msg := &entity.UserMessage{}
	if err := g.db.Model(&entity.UserMessage{}).Where("id = ? AND deleted_at = 0", msgId).First(msg).Error; err != nil {
		return nil, err
	}
	return msg, nil
}

func (g *MsgRepo) GetGroupMsgByID(msgId uint32) (*entity.GroupMessage, error) {
	msg := &entity.GroupMessage{}
	if err := g.db.Model(&entity.GroupMessage{}).Where("id = ? AND deleted_at = 0", msgId).First(msg).Error; err != nil {
		return nil, err
	}
	return msg, nil
}

func (g *MsgRepo) UpdateUserMsgColumn(msgId uint32, column string, value interface{}) error {
	return g.db.Model(&entity.UserMessage{}).Where("id = ?", msgId).Update(column, value).Error
}

func (g *MsgRepo) UpdateGroupMsgColumn(msgId uint32, column string, value interface{}) error {
	return g.db.Model(&entity.GroupMessage{}).Where("id = ?", msgId).Update(column, value).Error
}

func (g *MsgRepo) GetUserMsgLabelByDialogId(dialogId uint32) ([]*entity.UserMessage, error) {
	var userMessages []*entity.UserMessage
	if err := g.db.Model(&entity.UserMessage{}).Where("dialog_id = ? AND is_label = ? AND deleted_at = 0", dialogId, entity.IsLabel).Find(&userMessages).Error; err != nil {
		return nil, err
	}
	return userMessages, nil
}

func (g *MsgRepo) GetGroupMsgLabelByDialogId(dialogId uint32) ([]*entity.GroupMessage, error) {
	var groupMessages []*entity.GroupMessage
	if err := g.db.Model(&entity.GroupMessage{}).Where("dialog_id = ? AND is_label = ? AND deleted_at = 0", dialogId, entity.IsLabel).Find(&groupMessages).Error; err != nil {
		return nil, err
	}
	return groupMessages, nil
}

func (g *MsgRepo) SetUserMsgsReadStatus(msgIds []uint32, dialogId uint32) error {
	return g.db.Model(&entity.UserMessage{}).
		Where("id IN (?) AND dialog_id = ? AND deleted_at = 0", msgIds, dialogId).
		Updates(map[string]interface{}{
			"is_read": entity.IsRead,
			"read_at": time.Now(),
		}).Error
}

func (g *MsgRepo) SetUserMsgReadStatus(msgId uint32, isRead entity.ReadType) error {
	dd := g.db.Model(&entity.UserMessage{}).Where("id = ? AND deleted_at = 0", msgId)
	if isRead == entity.IsRead {
		return dd.Updates(map[string]interface{}{
			"is_read": entity.IsRead,
			"read_at": time.Now(),
		}).Error
	} else {
		return dd.Updates(map[string]interface{}{
			"is_read": entity.IsRead,
			"read_at": 0,
		}).Error
	}
}

func (g *MsgRepo) GetUnreadUserMsgs(uid string, dialogId uint32) ([]*entity.UserMessage, error) {
	var userMessages []*entity.UserMessage
	if err := g.db.Model(&entity.UserMessage{}).Where("dialog_id = ? AND is_read = ? AND receive_id = ? AND deleted_at = 0",
		dialogId,
		entity.NotRead,
		uid).
		Where("type NOT IN (?)", []entity.UserMessageType{entity.MessageTypeLabel, entity.MessageTypeDelete}).
		Find(&userMessages).Error; err != nil {
		return nil, err
	}
	return userMessages, nil
}

func (g *MsgRepo) GetBatchUserMsgsBurnAfterReadingMessages(msgIds []uint32, dialogID uint32) ([]*entity.UserMessage, error) {
	var userMessages []*entity.UserMessage
	err := g.db.Model(&entity.UserMessage{}).
		Where("dialog_id = ? AND id IN (?) AND is_burn_after_reading = ?", dialogID, msgIds, entity.IsBurnAfterReading).
		Find(&userMessages).Error
	if err != nil {
		return nil, err
	}
	return userMessages, nil
}

func (g *MsgRepo) PhysicalDeleteUserMessages(msgIds []uint32) error {
	return g.db.Delete(&entity.UserMessage{}, msgIds).Error
}

func (g *MsgRepo) PhysicalDeleteGroupMessages(msgIds []uint32) error {
	return g.db.Delete(&entity.GroupMessage{}, msgIds).Error
}

func (g *MsgRepo) LogicalDeleteUserMessages(msgIds []uint32) error {
	return g.db.Model(&entity.UserMessage{}).
		Where("id IN (?)", msgIds).
		Update("deleted_at", time.Now()).Error
}

func (g *MsgRepo) LogicalDeleteGroupMessages(msgIds []uint32) error {
	return g.db.Model(&entity.GroupMessage{}).
		Where("id IN (?)", msgIds).
		Update("deleted_at", time.Now()).Error
}

func (g *MsgRepo) GetUserMsgIdAfterMsgList(dialogId uint32, msgIds uint32) ([]*entity.UserMessage, int64, error) {
	var userMessages []*entity.UserMessage
	var total int64
	query := g.db.Model(&entity.UserMessage{}).
		Where("dialog_id = ? AND id > ? AND deleted_at = 0", dialogId, msgIds).
		Order("id desc")
	err := g.db.Model(&entity.UserMessage{}).
		Where("dialog_id = ?  AND deleted_at = 0", dialogId).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Find(&userMessages).Error
	if err != nil {
		return nil, 0, err
	}
	return userMessages, total, err
}

func (g *MsgRepo) GetGroupMsgIdAfterMsgList(dialogId uint32, msgIds uint32) ([]*entity.GroupMessage, int64, error) {
	var groupMessages []*entity.GroupMessage
	var total int64
	query := g.db.Model(&entity.GroupMessage{}).
		Where("dialog_id = ? AND id > ? AND deleted_at = 0", dialogId, msgIds).
		Order("id desc")
	err := g.db.Model(&entity.GroupMessage{}).
		Where("dialog_id = ? AND deleted_at = 0", dialogId).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Find(&groupMessages).Error
	if err != nil {
		return nil, 0, err
	}
	return groupMessages, total, err
}

func (g *MsgRepo) GetGroupMsgList(list dataTransformers.GroupMsgList) (*dataTransformers.GroupMsgListResponse, error) {
	response := &dataTransformers.GroupMsgListResponse{}

	query := g.db.Model(&entity.GroupMessage{}).
		Where("dialog_id = ? AND deleted_at = 0", list.DialogId)

	var total int64
	err := query.Count(&total).Error
	if err != nil {
		return response, err
	}
	if list.UserID != "" {
		query = query.Where("user_id = ?", list.UserID)
	}

	if list.Content != "" {
		query = query.Where("content LIKE ?", "%"+list.Content+"%")
	}

	if list.MsgType != 0 {
		query = query.Where("msg_type = ?", list.MsgType)
	}

	var groupMessages []entity.GroupMessage
	err = query.Order("id DESC").
		Limit(list.PageSize).
		Offset(list.PageSize * (list.PageNumber - 1)).
		Find(&groupMessages).
		Error
	if err != nil {
		return response, err
	}

	response.GroupMessages = groupMessages
	response.Total = int32(total)
	response.CurrentPage = int32(list.PageNumber)

	return response, nil
}

func (g *MsgRepo) GetGroupMsgsByIDs(msgIds []uint32) ([]*entity.GroupMessage, error) {
	var groupMessages []*entity.GroupMessage
	err := g.db.Model(&entity.GroupMessage{}).
		Where("id IN (?) AND deleted_at = 0", msgIds).
		Find(&groupMessages).Error

	return groupMessages, err
}

func (g *MsgRepo) GetGroupMsgIdsByDialogID(dialogID uint32) ([]uint32, error) {
	var msgIds []uint32
	err := g.db.Model(&entity.GroupMessage{}).
		Where("dialog_id = ? AND deleted_at = 0", dialogID).
		Select("id").
		Where("type NOT IN (?)", []entity.UserMessageType{entity.MessageTypeLabel, entity.MessageTypeDelete}).
		Find(&msgIds).Error
	return msgIds, err
}

func (g *MsgRepo) GetUserMsgByIDs(msgIds []uint32) ([]*entity.UserMessage, error) {
	var userMessages []*entity.UserMessage
	err := g.db.Model(&entity.UserMessage{}).
		Where("id IN (?) AND deleted_at = 0", msgIds).
		Find(&userMessages).Error
	return userMessages, err
}

func (g *MsgRepo) InsertUserMessages(message []entity.UserMessage) error {
	return g.db.Create(&message).Error
}

func (g *MsgRepo) DeleteUserMessagesByDialogID(dialogId uint32) error {
	return g.db.Model(&entity.UserMessage{}).Where("dialog_id = ?", dialogId).Update("deleted_at", time.Now()).Error
}

func (g *MsgRepo) DeleteGroupMessagesByDialogID(dialogId uint32) error {
	return g.db.Model(&entity.GroupMessage{}).Where("dialog_id = ?", dialogId).Update("deleted_at", time.Now()).Error
}

func (g *MsgRepo) UpdateUserMsgColumnByDialogId(dialogId uint32, column string, value interface{}) error {
	return g.db.Model(&entity.UserMessage{}).Where("dialog_id = ?", dialogId).Update(column, value).Error
}

func (g *MsgRepo) UpdateGroupMsgColumnByDialogId(dialogId uint32, column string, value interface{}) error {
	return g.db.Model(&entity.GroupMessage{}).Where("dialog_id = ?", dialogId).Update(column, value).Error
}

func (g *MsgRepo) PhysicalDeleteUserMessagesByDialogID(dialogId uint32) error {
	return g.db.Where("dialog_id = ?", dialogId).Delete(&entity.UserMessage{}).Error

}

func (g *MsgRepo) PhysicalDeleteGroupMessagesByDialogID(dialogId uint32) error {
	return g.db.Where("dialog_id = ?", dialogId).Delete(&entity.GroupMessage{}).Error
}

func (g *MsgRepo) GetUserDialogLastMsgs(dialogId uint32, pageNumber, pageSize int) ([]entity.UserMessage, int64, error) {
	var userMessages []entity.UserMessage
	var total int64
	query := g.db.Model(&entity.UserMessage{}).
		Where("dialog_id = ? AND is_burn_after_reading = ? AND deleted_at = 0", dialogId, entity.NotBurnAfterReading).
		Order("id DESC")

	err := g.db.Model(&entity.UserMessage{}).
		Where("dialog_id = ?  AND deleted_at = 0", dialogId).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Limit(pageSize).
		Offset(pageSize * (pageNumber - 1)).
		Find(&userMessages).Error
	if err != nil {
		return nil, 0, err
	}
	return userMessages, total, err
}

func (g *MsgRepo) GetGroupDialogLastMsgs(dialogId uint32, pageNumber, pageSize int) ([]entity.GroupMessage, int64, error) {
	var groupMessages []entity.GroupMessage
	var total int64
	query := g.db.Model(&entity.GroupMessage{}).
		Where("dialog_id = ? AND deleted_at = 0", dialogId).
		Order("id DESC")

	err := g.db.Model(&entity.GroupMessage{}).
		Where("dialog_id = ? AND deleted_at = 0", dialogId).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Limit(pageSize).
		Offset(pageSize * (pageNumber - 1)).
		Find(&groupMessages).Error
	if err != nil {
		return nil, 0, err
	}
	return groupMessages, total, err
}

func (g *MsgRepo) GetLastUserMsgsByDialogIDs(dialogIds []uint) ([]*entity.UserMessage, error) {
	userMessages := make([]*entity.UserMessage, 0)

	for _, dialogId := range dialogIds {
		var lastMsg entity.UserMessage
		g.db.Where("dialog_id = ? AND deleted_at = 0", dialogId).Order("created_at DESC").First(&lastMsg)
		if lastMsg.ID != 0 {
			userMessages = append(userMessages, &lastMsg)
		}
	}
	return userMessages, nil
}

func (g *MsgRepo) GetLastGroupMsgsByDialogIDs(dialogIds []uint) ([]*entity.GroupMessage, error) {
	groupMessages := make([]*entity.GroupMessage, 0)
	for _, dialogId := range dialogIds {
		var lastMsg entity.GroupMessage
		g.db.Where("dialog_id = ?  AND deleted_at = 0", dialogId).Order("created_at DESC").First(&lastMsg)
		if lastMsg.ID != 0 {
			groupMessages = append(groupMessages, &lastMsg)
		}
	}
	return groupMessages, nil
}

func (g *MsgRepo) GetUserMsgIdBeforeMsgList(dialogId uint32, msgId uint32, pageSize int) ([]*entity.UserMessage, int32, error) {
	var userMessages []*entity.UserMessage
	var total int64
	err := g.db.Model(&entity.UserMessage{}).Where("dialog_id = ? AND deleted_at = 0", dialogId).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = g.db.Model(&entity.UserMessage{}).
		Where("dialog_id = ? AND id < ? AND deleted_at = 0", dialogId, msgId).Order("id DESC").
		Limit(pageSize).
		Find(&userMessages).Error
	if err != nil {
		return nil, 0, err
	}
	return userMessages, int32(total), err
}

func (g *MsgRepo) GetGroupMsgIdBeforeMsgList(dialogId uint32, msgId uint32, pageSize int) ([]*entity.GroupMessage, int32, error) {
	var groupMessages []*entity.GroupMessage
	var total int64
	err := g.db.Model(&entity.GroupMessage{}).Where("dialog_id = ? AND deleted_at = 0", dialogId).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = g.db.Model(&entity.GroupMessage{}).
		Where("dialog_id = ? AND id < ? AND deleted_at = 0", dialogId, msgId).Order("id DESC").
		Limit(pageSize).
		Find(&groupMessages).Error
	if err != nil {
		return nil, 0, err
	}
	return groupMessages, int32(total), err
}
