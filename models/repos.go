package models

import (
	"fmt"
	"github.com/google/go-github/github"
	"log"
)

type RepoCollection struct {
	Repos []*github.Repository `json:"repos"`
}

func GetRepos(client *github.Client) RepoCollection {
	var repoCollection RepoCollection
	repos, _, err := client.Repositories.List("highlanderkev", nil)
	if err != nil {
		log.Fatal(err)
	}
	repoCollection.Repos = repos
	fmt.Printf("%s\n", repoCollection.Repos)
	return repoCollection
}
