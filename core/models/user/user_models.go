package user

import (
	"time"
)

type UserModel struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"-" gorm:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RegisterUser struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserListRequest struct {
	Limit int `query:"limit"  validate:"required"`
	Page  int `query:"page"  validate:"required"`
}

type UserListData struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateUser struct {
	ID        int       `param:"id" json:"id" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"required"`
	Phone     string    `json:"phone" validate:"required"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetUserData struct {
	ID int `param:"id" validate:"required"`
}

type DeleteUser struct {
	ID int `param:"id" validate:"required"`
}
