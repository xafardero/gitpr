package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v28/github"
)

//CreatePRDTO is a DTO
type CreatePRDTO struct {
	Title string
	Head  string
	Base  string
	Body  string
	Owner string
	Repo  string
}

//CreatePR this create a PR (convert this into struct)
func CreatePR(cpd CreatePRDTO) {
	ctx := context.Background()

	cl := Client()

	if cpd.Title == "" {
		cpd.Title = "Default message"
	}

	npp := github.NewPullRequest{
		Title: &cpd.Title,
		Head:  &cpd.Head,
		Base:  &cpd.Base,
		Body:  &cpd.Body,
		// "Draft": false,
	}

	repos, _, err := cl.PullRequests.Create(ctx, cpd.Owner, cpd.Repo, &npp)

	if err == nil {
		fmt.Println(repos)
	}

	fmt.Println(err)
}
