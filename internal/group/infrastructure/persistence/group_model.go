package persistence

//import (
//	ptime "github.com/cossim/coss-server/pkg/utils/time"
//	"gorm.io/gorm"
//)
//
//type GroupModel struct {
//	BaseModel
//	Type            uint   `gorm:"default:0;comment:群聊类型(0=私密群, 1=公开群)" json:"type"`
//	Status          uint   `gorm:"comment:群聊状态(0=正常状态, 1=锁定状态, 2=删除状态)" json:"status"`
//	MaxMembersLimit int    `gorm:"comment:群聊人数限制" json:"max_members_limit"`
//	CreatorID       string `gorm:"type:varchar(64);comment:创建者id" json:"creator_id"`
//	Name            string `gorm:"comment:群聊名称" json:"name"`
//	Avatar          string `gorm:"default:'';comment:头像（群）" json:"avatar"`
//}
//
//type BaseModel struct {
//	ID        uint  `gorm:"primaryKey;autoIncrement;" json:"id"`
//	CreatedAt int64 `gorm:"autoCreateTime;comment:创建时间" json:"created_at"`
//	UpdatedAt int64 `gorm:"autoUpdateTime;comment:更新时间" json:"updated_at"`
//	DeletedAt int64 `gorm:"default:0;comment:删除时间" json:"deleted_at"`
//}
//
//func (bm *BaseModel) BeforeCreate(tx *gorm.DB) error {
//	now := ptime.Now()
//	bm.CreatedAt = now
//	bm.UpdatedAt = now
//	return nil
//}
//
//func (bm *BaseModel) BeforeUpdate(tx *gorm.DB) error {
//	bm.UpdatedAt = ptime.Now()
//	return nil
//}
