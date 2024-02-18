package user

import (
	"errors"

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

func (repo *Repository) InsertUser(user *model.UserModel) error {
	err := repo.DB.Table(tableUser).Create(user).Error
	if err != nil {
		log.Error("error insert user: " + err.Error())
	}
	return err
}

func (repo *Repository) GetUserList(params *model.UserListRequest) (list []model.UserListData, total int64, err error) {
	page := params.Page
	limit := params.Limit
	result := repo.DB.
		Select(`
			id, 
			name, 
			email,
			phone, 
			created_at,
			updated_at  
		`).
		Table(`user`)

	result.Count(&total)

	res, err := result.Limit(limit).Offset((page - 1) * limit).Order(`id DESC`).Rows()
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}
	}

	for res.Next() {
		var data = model.UserListData{}
		repo.DB.ScanRows(res, &data)
		list = append(list, data)
	}

	return
}

func (repo *Repository) GetUserByID(id int) (user model.UserListData, err error) {
	result := repo.DB.
		Select(`
			id, 
			name, 
			email,
			phone, 
			created_at,
			updated_at  
		`).
		Table(`user`).
		Where("id=?", id).
		First(&user)

	return user, result.Error
}

func (repo *Repository) UpdateUser(params *model.UpdateUser, id int) error {
	result := repo.DB.Table("user").Where("id=?", id).Updates(params)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return errors.New("user id not exist")
	}

	return nil
}

func (repo *Repository) DeleteUser(id int) error {
	return repo.DB.Table("user").Where("id=?", id).Delete(model.UserModel{}).Error
}
