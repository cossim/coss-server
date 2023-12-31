package repository

import "github.com/cossim/coss-server/service/relation/domain/entity"

type GroupRelationRepository interface {
	InsertUserGroup(ur *entity.UserGroup) (*entity.UserGroup, error)
	GetUserGroupIDs(gid uint) ([]string, error)
}
