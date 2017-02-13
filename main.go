package main

import (
	"fmt"
	"github.com/google/go-github/github"
	"github.com/highlanderkev/api/handlers"
	"github.com/labstack/echo"
	"gopkg.in/zabawaba99/firego.v1"
	"log"
	"os"
)

func main() {
	fmt.Printf("Hello world\n")
	ip := os.Getenv("IP")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	fmt.Println("PORT:", port, os.Getenv("DATABASE_URL"))

	f := firego.New("https://highlanderkev-github-io.firebaseio.com/", nil)
	client := github.NewClient(nil)

	// db := pg.Connect(&pg.Options{
	// 	Addr:     os.Getenv("DATABASE_URL"),
	// 	User:     os.Getenv("DATABASE_USER"),
	// 	Password: os.Getenv("DATABASE_PASSWORD"),
	// 	Database: os.Getenv("DATABASE")
	// })
	// migrateDB(db)

	router := echo.New()
	router.GET("/", handlers.GetStatus())
	router.GET("/status", handlers.GetStatus())
	router.GET("/version", handlers.GetVersion())
	router.GET("/repos", handlers.GetRepos(client))
	router.GET("/posts", handlers.GetPosts(f))
	//router.POST("/post", handlers.SavePost(db))
	// router.DELETE("/posts/:id", handlers.DeletePost(db))

	router.Start(ip + ":" + port)
}
