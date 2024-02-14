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
