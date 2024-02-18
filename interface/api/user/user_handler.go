package user

import (
	"net/http"
	model "pajarwp11/go-experiment/core/models"
	userModel "pajarwp11/go-experiment/core/models/user"
	port "pajarwp11/go-experiment/core/port/user"
	"pajarwp11/go-experiment/utils/meta"
	"pajarwp11/go-experiment/utils/validator"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	svc port.ServiceContract
}

func New(userService port.ServiceContract) *Handler {
	return &Handler{
		svc: userService,
	}
}

func (h *Handler) RegisterUser(c echo.Context) error {
	res := new(model.DefaultResponse)
	req := new(userModel.RegisterUser)

	if err := c.Bind(&req); err != nil {
		res.Status.Message = err.Error()
		res.Status.Code = "4220000"
		return c.JSON(http.StatusUnprocessableEntity, res)
	}

	errorValidation := validator.ValidateReq(req)
	if errorValidation != nil {
		res.Status.Message = "register user validation error"
		res.Status.Code = "4220001"
		res.Errors = errorValidation
		return c.JSON(http.StatusUnprocessableEntity, res)
	}

	res = h.svc.RegisterUser(req)

	httpCode := res.Status.Code[0:3]
	httpCodeInt, _ := strconv.Atoi(httpCode)
	return c.JSON(httpCodeInt, res)
}

func (h *Handler) GetUserList(c echo.Context) error {
	res := new(model.DefaultResponse)
	req := new(userModel.UserListRequest)

	if err := c.Bind(req); err != nil {
		res.Status.Message = err.Error()
		res.Status.Code = "4220100"
		return c.JSON(http.StatusUnprocessableEntity, res)
	}

	errorValidation := validator.ValidateReq(req)
	if errorValidation != nil {
		res.Status.Message = "get user list validation error"
		res.Status.Code = "4220101"
		res.Errors = errorValidation
		return c.JSON(http.StatusUnprocessableEntity, res)
	}

	if req.Limit <= 0 {
		req.Limit = 10
	}
	if req.Page <= 0 {
		req.Page = 1
	}

	res, total := h.svc.GetUserList(req)

	from := (req.Page * req.Limit) - req.Limit
	res.Meta = meta.CreateMeta(req.Page, req.Limit, from, int(total))

	httpCode := res.Status.Code[0:3]
	httpCodeInt, _ := strconv.Atoi(httpCode)
	return c.JSON(httpCodeInt, res)
}

func (h *Handler) GetUserData(c echo.Context) error {
	res := new(model.DefaultResponse)
	req := new(userModel.GetUserData)

	if err := c.Bind(req); err != nil {
		res.Status.Message = err.Error()
		res.Status.Code = "4220200"
		return c.JSON(http.StatusUnprocessableEntity, res)
	}

	errorValidation := validator.ValidateReq(req)
	if errorValidation != nil {
		res.Status.Message = "get user data validation error"
		res.Status.Code = "4220201"
		res.Errors = errorValidation
		return c.JSON(http.StatusUnprocessableEntity, res)
	}

	res = h.svc.GetUserData(req.ID)

	httpCode := res.Status.Code[0:3]
	httpCodeInt, _ := strconv.Atoi(httpCode)
	return c.JSON(httpCodeInt, res)
}

func (h *Handler) UpdateUser(c echo.Context) error {
	res := new(model.DefaultResponse)
	req := new(userModel.UpdateUser)

	if err := c.Bind(req); err != nil {
		res.Status.Message = err.Error()
		res.Status.Code = "4220300"
		return c.JSON(http.StatusUnprocessableEntity, res)
	}

	errorValidation := validator.ValidateReq(req)
	if errorValidation != nil {
		res.Status.Message = "update user validation error"
		res.Status.Code = "4220301"
		res.Errors = errorValidation
		return c.JSON(http.StatusUnprocessableEntity, res)
	}

	res = h.svc.UpdateUser(req, req.ID)

	httpCode := res.Status.Code[0:3]
	httpCodeInt, _ := strconv.Atoi(httpCode)
	return c.JSON(httpCodeInt, res)
}

func (h *Handler) DeleteUser(c echo.Context) error {
	res := new(model.DefaultResponse)
	req := new(userModel.GetUserData)

	if err := c.Bind(req); err != nil {
		res.Status.Message = err.Error()
		res.Status.Code = "4220400"
		return c.JSON(http.StatusUnprocessableEntity, res)
	}

	errorValidation := validator.ValidateReq(req)
	if errorValidation != nil {
		res.Status.Message = "delete user validation error"
		res.Status.Code = "4220401"
		res.Errors = errorValidation
		return c.JSON(http.StatusUnprocessableEntity, res)
	}

	res = h.svc.DeleteUser(req.ID)

	httpCode := res.Status.Code[0:3]
	httpCodeInt, _ := strconv.Atoi(httpCode)
	return c.JSON(httpCodeInt, res)
}
