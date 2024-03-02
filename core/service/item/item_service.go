package item

import (
	model "pajarwp11/go-experiment/core/models"
	itemModel "pajarwp11/go-experiment/core/models/item"
	port "pajarwp11/go-experiment/core/port/item"
)

type itemService struct {
	repo port.RepoContract
}

func New(repo port.RepoContract) *itemService {
	return &itemService{
		repo: repo,
	}
}

func (i *itemService) InsertItem(req *itemModel.InsertItem) *model.DefaultResponse {
	resp := new(model.DefaultResponse)
	err := i.repo.InsertItem(req)
	if err != nil {
		resp.Status.Code = "5000500"
		resp.Status.Message = "insert item error:" + err.Error()
		return resp
	}

	resp.Data = req
	resp.Status.Code = "2010500"
	resp.Status.Message = "insert item success"

	return resp
}

func (i *itemService) GetItemData(id string) *model.DefaultResponse {
	var err error
	var data itemModel.Item
	resp := new(model.DefaultResponse)
	data, err = i.repo.GetItemByID(id)
	if err != nil {
		resp.Status.Code = "5000600"
		resp.Status.Message = "get item data error:" + err.Error()
		return resp
	}
	resp.Data = data
	resp.Status.Code = "2000600"
	resp.Status.Message = "get item data success"
	return resp
}

func (i *itemService) UpdateItem(params *itemModel.InsertItem, id string) *model.DefaultResponse {
	resp := new(model.DefaultResponse)
	err := i.repo.UpdateItem(id, params)
	if err != nil {
		resp.Status.Code = "5000700"
		resp.Status.Message = "update item error:" + err.Error()
		return resp
	}
	resp.Data = params
	resp.Status.Code = "200700"
	resp.Status.Message = "update item success"
	return resp
}

func (i *itemService) DeleteItem(id string) *model.DefaultResponse {
	resp := new(model.DefaultResponse)
	err := i.repo.DeleteItem(id)
	if err != nil {
		resp.Status.Code = "5000800"
		resp.Status.Message = "delete item error:" + err.Error()
		return resp
	}
	resp.Status.Code = "200800"
	resp.Status.Message = "delete item success"
	return resp
}
