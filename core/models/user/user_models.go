package user

import (
	"time"
)

type UserModel struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

type RegisterUser struct {
	Name     string `json:"name" bson:"name" validate:"required"`
	Email    string `json:"email" bson:"email" validate:"required"`
	Phone    string `json:"phone" bson:"phone" validate:"required"`
	Password string `json:"password" bson:"password" validate:"required"`
}
