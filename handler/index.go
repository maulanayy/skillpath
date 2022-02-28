package handler

import (
	"net/http"
	"skillpath/app"
	"skillpath/model"

	"github.com/labstack/echo"
)

type Handler struct {
	Usecase    app.Usecase
	Middleware *app.Middleware
	Config     *model.Config
}

func NewHandler(UC app.Usecase, cfg *model.Config) *Handler {
	MW := app.NewMiddleware(UC, cfg)

	return &Handler{
		Usecase:    UC,
		Middleware: MW,
		Config:     cfg,
	}
}

func badRequest(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, model.ResponseJSON{
		Message: err.Error(),
	})
}

func okResponse(c echo.Context, msg string, data interface{}) error {
	return c.JSON(http.StatusOK, model.ResponseJSON{
		Message: msg,
		Data:    data,
	})
}
