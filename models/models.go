package models

import (
	"database/sql"
	"time"
)

type Post struct {
	Id          int    `json:"id"`
	Title       string `json:"name"`
	SubTitle    string
	Description string
	Body        string
	DatePosted  Time
	DateUpdated Time
}
