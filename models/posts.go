package models

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

type Post struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
}

type PostCollection struct {
    Posts []Post `json:"posts"`
}

func GetPosts(db *sql.DB) PostCollection {
    rows, err := db.Query("SELECT id, title FROM posts")
    if err != nil {
        log.Fatalf("error: %q", err)
        // panic(err)
    }
    
    defer rows.Close()
    result := PostCollection{}
    for rows.Next() {
        post := Post{}
        err2 := rows.Scan(&post.ID, &post.Title)
        if err2 != nil {
            log.Fatalf("error: %q", err2)
            // panic(err2)
        }
        result.Posts = append(result.Posts, post)
    }
    return result
}

func PutPost(db *sql.DB, title string) (int64, error) {
    sql := "INSERT INTO posts(title) VALUES(?)"
    
    stmt, err := db.Prepare(sql)
    if err != nil {
        log.Fatalf("error: %q", err)
        panic(err)
    }
    
    defer stmt.Close()
    
    result, err2 := stmt.Exec(title)
    if err2 != nil {
        log.Fatalf("error: %q", err2)
        panic(err2)
    }
    
    return result.LastInsertId()
}

func DeletePost(db *sql.DB, id int) (int64, error) {
    sql := "DELETE FROM posts WHERE id = ?"
    
    stmt, err := db.Prepare(sql)
    if err != nil {
        log.Fatalf("error: %q", err)
        panic(err)
    }
    
    result, err2 := stmt.Exec(id)
    if err2 != nil {
        log.Fatalf("error: %q", err2)
        panic(err2)
    }
    
    return result.RowsAffected()
}

// 	SubTitle    string
// 	Description string
// 	Body        string
// 	DatePosted  time.Time
// 	DateUpdated time.Time