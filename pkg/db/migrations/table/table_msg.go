package table

import (
	"github.com/cossim/coss-server/service/msg/domain/entity"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func (d InitDatabase) AddTableUserMsg() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202401031200",
		Migrate: func(tx *gorm.DB) error {
			// 执行迁移操作，例如创建表
			return tx.AutoMigrate(&entity.UserMessage{})
		},
	}
}

func (d InitDatabase) AddTableGroupMsg() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202401031200",
		Migrate: func(tx *gorm.DB) error {
			// 执行迁移操作，例如创建表
			return tx.AutoMigrate(&entity.GroupMessage{})
		},
	}
}
func (d InitDatabase) AddTableDialog() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202401031201",
		//Migrate: func(tx *gorm.DB) error {
		//	// 执行迁移操作，例如创建表
		//	return tx.AutoMigrate(&entity.Dialog{})
		//},
	}
}
func (d InitDatabase) AddTableDialogUser() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202401031200",
		//Migrate: func(tx *gorm.DB) error {
		//	// 执行迁移操作，例如创建表
		//	return tx.AutoMigrate(&entity.DialogUser{})
		//},
	}
}
