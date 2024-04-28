package relation

import (
	"context"
)

type UserQuery struct {
	UserId   string
	FriendId []string
	Status   *UserRelationStatus
	Force    bool
}

type UserRepository interface {
	Get(ctx context.Context, userId, friendId string) (*UserRelation, error)
	Create(ctx context.Context, ur *UserRelation) (*UserRelation, error)
	Update(ctx context.Context, ur *UserRelation) (*UserRelation, error)
	Delete(ctx context.Context, userId, friendId string) error

	Find(ctx context.Context, query *UserQuery) ([]*UserRelation, error)

	// EstablishFriendship 建立双向好友关系
	// 如果创建成功，将建立两个用户之间的好友关系，其中一个是发送者到接收者的关系，
	// 另一个是接收者到发送者的关系。
	EstablishFriendship(ctx context.Context, dialogID uint32, senderID, receiverID string) error

	// RestoreFriendship 单向删除后重新建立好友关系
	RestoreFriendship(ctx context.Context, dialogID uint32, userId, friendId string) error

	// Blacklist 获取黑名单
	Blacklist(ctx context.Context, userId string) (*Blacklist, error)

	// FriendRequestList 获取好友请求列表
	FriendRequestList(ctx context.Context, userId string) ([]*Friend, error)

	UpdateStatus(ctx context.Context, id uint32, status UserRelationStatus) error

	// UpdateFields 根据 ID 更新 UserRelation 对象的多个字段
	UpdateFields(ctx context.Context, id uint32, fields map[string]interface{}) error

	// SetUserFriendSilentNotification 设置好友静默通知
	SetUserFriendSilentNotification(ctx context.Context, uid, friendId string, silentNotification bool) error
	// SetUserOpenBurnAfterReading 设置好友阅后即焚
	SetUserOpenBurnAfterReading(ctx context.Context, uid, friendId string, openBurnAfterReading bool, burnAfterReadingTimeOut int64) error

	// SetFriendRemark 设置好友备注
	SetFriendRemark(ctx context.Context, userId, friendId string, remark string) error
}
