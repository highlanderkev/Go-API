package handlers

import (
	"database/sql"
	"github.com/highlanderkev/api/models"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
)

type H map[string]interface{}

func GetPosts(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetPosts(db))
	}
}

func PutPost(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var post models.Post
		c.Bind(&post)
		id, err := models.PutPost(db, post.Title)
		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
		} else {
			log.Fatalf("Error putting post: %q", err)
			return err
		}
	}
}

func DeletePost(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		_, err := models.DeletePost(db, id)
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
		} else {
			log.Fatalf("Error deleting post: %q", err)
			return err
		}
	}
}
