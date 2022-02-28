package route

import (
	"net/http"
	"skillpath/handler"

	"github.com/labstack/echo"
)

func RegisterRoute(e *echo.Echo, h *handler.Handler) {
	// e.Use(h.Middleware.MiddlewareLogging, h.Middleware.Recover)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	v1 := e.Group("/v1")

	v1.POST("/login", h.Login)
	v1.POST("/register", h.Register)

}
