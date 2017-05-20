package main

import (
	context "context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/google/go-github/github"
	pb "github.com/leewind/git-project-collection/proto/collector"
	ncontext "golang.org/x/net/context"
	"golang.org/x/oauth2"
)

const (
	port = ":50051"
)

func getStarredRepositories() ([]*pb.Repository, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "48b2c01475c4fb29648d429994461a577321978a"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	option := &github.ActivityListStarredOptions{
		ListOptions: github.ListOptions{PerPage: 30},
	}

	var result []*pb.Repository
	var allStarredRepositories []*github.StarredRepository
	for {
		repos, resp, err := client.Activity.ListStarred(ctx, "", option)
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

	fmt.Println(len(allStarredRepositories))
	for i := 0; i < len(allStarredRepositories); i++ {
		repository := allStarredRepositories[i].Repository

		fmt.Printf("id: %d, name: %s, url: %s, html: %s, desc: %s, star count: %d, git: %s, clone: %s \n",
			repository.GetID(), repository.GetFullName(), repository.GetURL(), repository.GetHTMLURL(),
			repository.GetDescription(), repository.GetStargazersCount(), repository.GetGitURL(), repository.GetCloneURL())

		result = append(result,
			&pb.Repository{Id: int32(repository.GetID()), Name: repository.GetFullName(), Url: repository.GetURL(),
				Html: repository.GetHTMLURL(), Desc: repository.GetDescription(), StarCount: int32(repository.GetStargazersCount()),
				Git: repository.GetGitURL(), Clone: repository.GetCloneURL()})

	}

	return result, nil
}

type server struct{}

func (s *server) GetStarredRepositories(ctx ncontext.Context, in *pb.Empty) (*pb.Repositories, error) {
	repos, err := getStarredRepositories()
	if err != nil {
		return nil, err
	}

	return &pb.Repositories{Repos: repos}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCollectorServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
