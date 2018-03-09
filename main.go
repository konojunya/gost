package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "4e0f44db361bf296dfe768f1395bf9bd172b7bd3"},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	gist := &github.Gist{
		Description: github.String("hoge"),
		Public:      github.Bool(true),
	}
	_, res, err := client.Gists.Create(ctx, gist)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.StatusCode)
}
