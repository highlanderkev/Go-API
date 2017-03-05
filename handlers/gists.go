package handlers

import (
	"github.com/google/go-github/github"
	"github.com/highlanderkev/api/models"
	"github.com/labstack/echo"
	"net/http"
)

func GetGists(client *github.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetGists(client))
	}
}