/*
Package service implement the interface of collector.proto
for github starred repositories information

		createClient

		Service
			client

			NewService
			GetStarredRepositories
*/
package service

import (
	"fmt"

	"github.com/google/go-github/github"
	"google.golang.org/grpc"

	context "context"

	collector "github.com/leewind/git-project-collection/api/collector"
	netContext "golang.org/x/net/context"

	"golang.org/x/oauth2"
)

// CollectorService service对象主体
type CollectorService struct {
	client *github.Client
}

func createClient() *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "xxxxxx_your_token_xxxxxx"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return client
}

// NewService 初始化Service，将service注册到server上
func NewService(server *grpc.Server) *CollectorService {
	client := createClient()
	service := &CollectorService{client}
	collector.RegisterCollectorServer(server, service)
	return service
}

// GetStarredRepositories 复写proto代码中的接口，获取所有收藏Repositories列表
func (service *CollectorService) GetStarredRepositories(ctx netContext.Context, in *collector.Empty) (*collector.Repositories, error) {
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
