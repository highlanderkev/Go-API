package main

import (
	"database/sql"
	"fmt"
	"github.com/highlanderkev/api/handlers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func migrateDB(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS posts(
		id INTEGER NOT NULL PRIMARY KEY,
		title VARCHAR NOT NULL
	);
	`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Error migrating database: %q", err)
	}
}

func main() {
	fmt.Printf("Hello world\n")
	ip := os.Getenv("IP")
	port := os.Getenv("PORT")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	fmt.Println("PORT:", port, os.Getenv("DATABASE_URL"))

	db, err2 := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err2 != nil {
		log.Fatalf("Error opening database: %q", err2)
	}
	migrateDB(db)

	router := echo.New()
	router.GET("/posts", handlers.GetPosts(db))
	router.PUT("/posts", handlers.PutPost(db))
	router.DELETE("/posts/:id", handlers.DeletePost(db))

	router.Start(ip + ":" + port)
}
