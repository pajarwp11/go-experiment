package user

import (
	model "pajarwp11/go-experiment/core/models"
	userModel "pajarwp11/go-experiment/core/models/user"
)

type ServiceContract interface {
	RegisterUser(req *userModel.RegisterUser) (resp *model.DefaultResponse, err error)
}
