package user

import (
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"

	model "pajarwp11/go-experiment/core/models/user"
)

var (
	tableUser = "user"
)

type (
	Repository struct {
		DB *gorm.DB
	}
)

func New(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (repo *Repository) InsertUser(user model.UserModel) error {
	err := repo.DB.Table(tableUser).Create(user).Error
	if err != nil {
		log.Error("error insert user: " + err.Error())
	}
	return err
}
