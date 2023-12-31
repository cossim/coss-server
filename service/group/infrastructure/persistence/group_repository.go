package persistence

import (
	"github.com/cossim/coss-server/service/group/domain/entity"
	"gorm.io/gorm"
)

type GroupRepo struct {
	db *gorm.DB
}

func NewGroupRepo(db *gorm.DB) *GroupRepo {
	return &GroupRepo{db: db}
}

func (repo *GroupRepo) GetGroupInfoByGid(gid uint) (*entity.Group, error) {
	var group entity.Group
	result := repo.db.First(&group, gid)
	if result.Error != nil {
		return nil, result.Error
	}
	return &group, nil
}

func (repo *GroupRepo) GetBatchGetGroupInfoByIDs(groupIds []uint) ([]*entity.Group, error) {
	var groups []*entity.Group
	result := repo.db.Find(&groups, groupIds)
	if result.Error != nil {
		return nil, result.Error
	}
	return groups, nil
}

func (repo *GroupRepo) UpdateGroup(group *entity.Group) (*entity.Group, error) {
	result := repo.db.Save(group)
	if result.Error != nil {
		return nil, result.Error
	}
	return group, nil
}

func (repo *GroupRepo) InsertGroup(group *entity.Group) (*entity.Group, error) {
	result := repo.db.Create(group)
	if result.Error != nil {
		return nil, result.Error
	}
	return group, nil
}

func (repo *GroupRepo) DeleteGroup(gid uint) error {
	result := repo.db.Model(&entity.Group{}).Where("id = ?", gid).Update("deleted_at", gorm.Expr("NOW()"))
	return result.Error
}
