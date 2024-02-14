package user

import (
	"net/http"
	model "pajarwp11/go-experiment/core/models"
	userModel "pajarwp11/go-experiment/core/models/user"
	port "pajarwp11/go-experiment/core/port/user"
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