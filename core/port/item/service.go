package item

import (
	model "pajarwp11/go-experiment/core/models"
	itemModel "pajarwp11/go-experiment/core/models/item"
)

type ServiceContract interface {
	InsertItem(req *itemModel.InsertItem) *model.DefaultResponse
	GetItemData(id string) *model.DefaultResponse
	UpdateItem(params *itemModel.InsertItem, id string) *model.DefaultResponse
	DeleteItem(id string) *model.DefaultResponse
}
