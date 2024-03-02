package item

import (
	model "pajarwp11/go-experiment/core/models/item"
)

type RepoContract interface {
	GetItemByID(id string) (doc model.Item, err error)
	InsertItem(itemData *model.InsertItem) error
	UpdateItem(id string, itemData *model.InsertItem) error
	DeleteItem(id string) error
}
