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
		log.Fatal(err)
	}
	gistCollection.Gists = gists
	return gistCollection
}
