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
	options := getGithubRepoOptions()
	for {
		repos, resp, err := client.Repositories.List("highlanderkev", options)
		if err != nil {
			log.Fatal(err)
		}
		repoCollection.Repos = append(repoCollection.Repos, repos...)
		if resp.NextPage == 0 {
			break
		}
		options.ListOptions.Page = resp.NextPage
	}
	//fmt.Printf("%s\n", repoCollection.Repos)
	return repoCollection
}

func getGithubRepoOptions() *github.RepositoryListOptions {
	return &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}
}
