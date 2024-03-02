package item

import (
	"net/http"
	model "pajarwp11/go-experiment/core/models"
	itemModel "pajarwp11/go-experiment/core/models/item"
	port "pajarwp11/go-experiment/core/port/item"
	"pajarwp11/go-experiment/utils/validator"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	svc port.ServiceContract
}

func New(itemService port.ServiceContract) *Handler {
	return &Handler{
		svc: itemService,
	}
}

func (h *Handler) InsertItem(c echo.Context) error {
	res := new(model.DefaultResponse)
	req := new(itemModel.InsertItem)

	if err := c.Bind(&req); err != nil {
		res.Status.Message = err.Error()
		res.Status.Code = "4220500"
		return c.JSON(http.StatusUnprocessableEntity, res)
	}

	errorValidation := validator.ValidateReq(req)
	if errorValidation != nil {
		res.Status.Message = "insert item validation error"
		res.Status.Code = "4220501"
		res.Errors = errorValidation
		return c.JSON(http.StatusUnprocessableEntity, res)
	}

	res = h.svc.InsertItem(req)

	httpCode := res.Status.Code[0:3]
	httpCodeInt, _ := strconv.Atoi(httpCode)
	return c.JSON(httpCodeInt, res)
}

func (h *Handler) GetItemData(c echo.Context) error {
	res := new(model.DefaultResponse)
	req := new(itemModel.GetItemData)

	if err := c.Bind(req); err != nil {
		res.Status.Message = err.Error()
		res.Status.Code = "4220600"
		return c.JSON(http.StatusUnprocessableEntity, res)
	}

	errorValidation := validator.ValidateReq(req)
	if errorValidation != nil {
		res.Status.Message = "get item data validation error"
		res.Status.Code = "4220601"
		res.Errors = errorValidation
		return c.JSON(http.StatusUnprocessableEntity, res)
	}

	res = h.svc.GetItemData(req.ID)

	httpCode := res.Status.Code[0:3]
	httpCodeInt, _ := strconv.Atoi(httpCode)
	return c.JSON(httpCodeInt, res)
}

func (h *Handler) UpdateItem(c echo.Context) error {
	res := new(model.DefaultResponse)
	req := new(itemModel.UpdateItem)

	if err := c.Bind(req); err != nil {
		res.Status.Message = err.Error()
		res.Status.Code = "4220700"
		return c.JSON(http.StatusUnprocessableEntity, res)
	}

	errorValidation := validator.ValidateReq(req)
	if errorValidation != nil {
		res.Status.Message = "update item validation error"
		res.Status.Code = "4220701"
		res.Errors = errorValidation
		return c.JSON(http.StatusUnprocessableEntity, res)
	}

	res = h.svc.UpdateItem(req)

	httpCode := res.Status.Code[0:3]
	httpCodeInt, _ := strconv.Atoi(httpCode)
	return c.JSON(httpCodeInt, res)
}

func (h *Handler) DeleteItem(c echo.Context) error {
	res := new(model.DefaultResponse)
	req := new(itemModel.GetItemData)

	if err := c.Bind(req); err != nil {
		res.Status.Message = err.Error()
		res.Status.Code = "4220800"
		return c.JSON(http.StatusUnprocessableEntity, res)
	}

	errorValidation := validator.ValidateReq(req)
	if errorValidation != nil {
		res.Status.Message = "delete item validation error"
		res.Status.Code = "4220801"
		res.Errors = errorValidation
		return c.JSON(http.StatusUnprocessableEntity, res)
	}

	res = h.svc.DeleteItem(req.ID)

	httpCode := res.Status.Code[0:3]
	httpCodeInt, _ := strconv.Atoi(httpCode)
	return c.JSON(httpCodeInt, res)
}
