package models

import (
	"fmt"
	firego "gopkg.in/zabawaba99/firego.v1"
	"log"
)

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type PostCollection struct {
	Posts map[string]interface{} `json:"posts"`
}

func GetPosts(f *firego.Firebase) PostCollection {
	var postCollection PostCollection
	if err := f.Value(&postCollection.Posts); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", postCollection.Posts)
	return postCollection
}
