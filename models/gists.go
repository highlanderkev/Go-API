package models

import (
	"github.com/google/go-github/github"
	"log"
)

type GistCollection struct {
	Gists []*github.Gist `json:"gists"`
}

func GetGists(client *github.Client) GistCollection {
	var gistCollection GistCollection
	gists, _, err := client.Gists.List("highlanderkev", nil)
	if err != nil {
		if _, ok := err.(*github.RateLimitError); ok {
			log.Fatal("Github Rate Limit Hit.")
		}
		log.Fatal(err)
	}
	gistCollection.Gists = gists
	return gistCollection
}
