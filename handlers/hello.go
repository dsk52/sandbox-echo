package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello")
	}
}