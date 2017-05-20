package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const (
	port = ":50051"
)

func main() {

}

func getStarredRepositories() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "48b2c01475c4fb29648d429994461a577321978a"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	option := &github.ActivityListStarredOptions{
		ListOptions: github.ListOptions{PerPage: 30},
	}

	var allStarredRepositories []*github.StarredRepository
	for {
		repos, resp, err := client.Activity.ListStarred(ctx, "", option)
		if err != nil {
			fmt.Println(err)
		}

		allStarredRepositories = append(allStarredRepositories, repos...)

		if resp.NextPage == 0 {
			break
		}

		option.Page = resp.NextPage
	}

	fmt.Println(len(allStarredRepositories))
	for i := 0; i < len(allStarredRepositories); i++ {
		repository := allStarredRepositories[i].Repository
		fmt.Printf("id: %d, name: %s, url: %s, html: %s, desc: %s, star count: %d, git: %s, clone: %s \n",
			repository.GetID(), repository.GetFullName(), repository.GetURL(), repository.GetHTMLURL(),
			repository.GetDescription(), repository.GetStargazersCount(), repository.GetGitURL(), repository.GetCloneURL())
	}
}
