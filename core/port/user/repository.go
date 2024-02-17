package user

import (
	model "pajarwp11/go-experiment/core/models/user"
)

type RepoContract interface {
	InsertUser(user *model.UserModel) error
	GetUserList(params *model.UserListRequest) (list []model.UserListData, total int64, err error)
}
