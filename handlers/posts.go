package handlers

import (
	"github.com/highlanderkev/api/models"
	"github.com/labstack/echo"
	"gopkg.in/zabawaba99/firego.v1"
	"net/http"
)

func GetPosts(f *firego.Firebase) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetPosts(f))
	}
}
