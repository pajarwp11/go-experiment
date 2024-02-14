package user

import (
	"errors"
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

func (u *userService) RegisterUser(req *userModel.RegisterUser) (resp *model.DefaultResponse, err error) {
	var userData *userModel.UserModel
	userData.Name = req.Name
	userData.Email = req.Email
	userData.Phone = req.Phone
	userData.CreatedAt = time.Now()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return resp, errors.New("register user error: error generate hash password")
	}

	userData.Password = string(hashedPassword)

	err = u.repo.InsertUser(userData)
	if err != nil {
		return resp, errors.New("register user error:" + err.Error())
	}

	return resp, nil
}
