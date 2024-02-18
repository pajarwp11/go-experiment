package user

import (
	"time"

	model "pajarwp11/go-experiment/core/models"
	userModel "pajarwp11/go-experiment/core/models/user"
	port "pajarwp11/go-experiment/core/port/user"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo port.RepoContract
}

func New(repo port.RepoContract) *userService {
	return &userService{
		repo: repo,
	}
}

func (u *userService) RegisterUser(req *userModel.RegisterUser) *model.DefaultResponse {
	userData := new(userModel.UserModel)
	resp := new(model.DefaultResponse)
	userData.Name = req.Name
	userData.Email = req.Email
	userData.Phone = req.Phone
	userData.CreatedAt = time.Now()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		resp.Status.Code = "5000000"
		resp.Status.Message = "register user error: error generate hash password"
		return resp
	}

	userData.Password = string(hashedPassword)

	err = u.repo.InsertUser(userData)
	if err != nil {
		resp.Status.Code = "5000001"
		resp.Status.Message = "register user error:" + err.Error()
		return resp
	}

	resp.Data = userData
	resp.Status.Code = "2010000"
	resp.Status.Message = "user register success"

	return resp
}

func (u *userService) GetUserList(params *userModel.UserListRequest) (*model.DefaultResponse, int64) {
	var err error
	var list []userModel.UserListData
	var total int64
	resp := new(model.DefaultResponse)
	list, total, err = u.repo.GetUserList(params)
	if err != nil {
		resp.Status.Code = "5000100"
		resp.Status.Message = "get user list error:" + err.Error()
		return resp, total
	}
	resp.Data = list
	resp.Status.Code = "2000100"
	resp.Status.Message = "get user list success"
	return resp, total
}

func (u *userService) GetUserData(id int) *model.DefaultResponse {
	var err error
	var data userModel.UserListData
	resp := new(model.DefaultResponse)
	data, err = u.repo.GetUserByID(id)
	if err != nil {
		resp.Status.Code = "5000200"
		resp.Status.Message = "get user data error:" + err.Error()
		return resp
	}
	resp.Data = data
	resp.Status.Code = "2000200"
	resp.Status.Message = "get user data success"
	return resp
}

func (u *userService) UpdateUser(params *userModel.UpdateUser, id int) *model.DefaultResponse {
	resp := new(model.DefaultResponse)
	params.UpdatedAt = time.Now()
	err := u.repo.UpdateUser(params, id)
	if err != nil {
		resp.Status.Code = "5000300"
		resp.Status.Message = "update user error:" + err.Error()
		return resp
	}
	resp.Data = params
	resp.Status.Code = "200300"
	resp.Status.Message = "update user success"
	return resp
}

func (u *userService) DeleteUser(id int) *model.DefaultResponse {
	resp := new(model.DefaultResponse)
	err := u.repo.DeleteUser(id)
	if err != nil {
		resp.Status.Code = "5000400"
		resp.Status.Message = "delete user error:" + err.Error()
		return resp
	}
	resp.Status.Code = "200400"
	resp.Status.Message = "delete user success"
	return resp
}
