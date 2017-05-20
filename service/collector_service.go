package service

import (
	"fmt"

	"github.com/google/go-github/github"
	"google.golang.org/grpc"

	context "context"

	collector "github.com/leewind/git-project-collection/proto/collector"
	netContext "golang.org/x/net/context"

	"golang.org/x/oauth2"
)

type Service struct {
	client *github.Client
}

func (service *Service) Register(server *grpc.Server) {
	collector.RegisterCollectorServer(server, service)
}

func (service *Service) ResetClient() {
	if service.client == nil {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: "xxxxxx_your_token_xxxxxx"},
		)
		tc := oauth2.NewClient(ctx, ts)

		service.client = github.NewClient(tc)
	}
}

func (service *Service) GetStarredRepositories(ctx netContext.Context, in *collector.Empty) (*collector.Repositories, error) {
	service.ResetClient()

	option := &github.ActivityListStarredOptions{
		ListOptions: github.ListOptions{PerPage: 30},
	}

	var result []*collector.Repository
	var allStarredRepositories []*github.StarredRepository
	for {
		repos, resp, err := service.client.Activity.ListStarred(ctx, "", option)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		allStarredRepositories = append(allStarredRepositories, repos...)

		if resp.NextPage == 0 {
			break
		}

		option.Page = resp.NextPage
	}

	for i := 0; i < len(allStarredRepositories); i++ {
		repository := allStarredRepositories[i].Repository

		result = append(result,
			&collector.Repository{
				Id: int32(repository.GetID()), Name: repository.GetFullName(), Url: repository.GetURL(),
				Html: repository.GetHTMLURL(), Desc: repository.GetDescription(), StarCount: int32(repository.GetStargazersCount()),
				Git: repository.GetGitURL(), Clone: repository.GetCloneURL(),
			})
	}

	return &collector.Repositories{Repos: result}, nil
}
