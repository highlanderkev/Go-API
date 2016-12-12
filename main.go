package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/highlanderkev/api/models"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Printf("Hello world\n")
	ip := os.Getenv("IP")
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	fmt.Println("PORT:", port)

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world")
	})

	router.Run(ip + ":" + port)
}
