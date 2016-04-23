package models



type Post struct {
    Id int
    Title string
    SubTitle string
    Description string
    Body string
    DatePosted time.Time
    DateUpdated time.Time
}