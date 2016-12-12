package models

import (
	"time"
)

type Post struct {
	Id          int    `json:"id"`
	Title       string `json:"name"`
	SubTitle    string
	Description string
	Body        string
	DatePosted  time.Time
	DateUpdated time.Time
}
