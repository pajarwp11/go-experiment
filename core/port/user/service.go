package user

import (
	model "pajarwp11/go-experiment/core/models"
	userModel "pajarwp11/go-experiment/core/models/user"
)

type ServiceContract interface {
	RegisterUser(req *userModel.RegisterUser) *model.DefaultResponse
	GetUserList(params *userModel.UserListRequest) (response *model.DefaultResponse, total int64)
}
