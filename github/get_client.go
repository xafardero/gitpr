package github

import (
	"context"

	"github.com/google/go-github/v28/github"
	"golang.org/x/oauth2"
)

//Client get client
func Client() *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "token"},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)

}
