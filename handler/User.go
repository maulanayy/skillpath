package handler

import (
	"skillpath/model"

	"github.com/labstack/echo"
)

func (h *Handler) Login(c echo.Context) error {

	var reqData model.Login
	if err := c.Bind(&reqData); err != nil {
		return badRequest(c, err)
	}
	data, err := h.Usecase.Login(reqData)
	if err != nil {
		return badRequest(c, err)
	}
	return okResponse(c, "login success", data)
}

func (h *Handler) Register(c echo.Context) error {

	var reqData model.Register
	if err := c.Bind(&reqData); err != nil {
		return badRequest(c, err)
	}

	user := model.User{
		Email:    reqData.Email,
		Password: reqData.Password,
		Info: model.Info{
			Name: reqData.Name,
		},
	}
	err := h.Usecase.Create(user)
	if err != nil {
		return badRequest(c, err)
	}
	return okResponse(c, "create success", nil)
}
