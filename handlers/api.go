package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

type H map[string]interface{}

func GetStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, H{
			"status": "OK",
		})
	}
}

func GetVersion() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, H{
			"version": "1",
		})
	}
}
