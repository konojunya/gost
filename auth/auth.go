package auth

import (
	"context"
	"log"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var ctx = context.Background()

func getToken() *string {
	token := os.Getenv("GOST_GITHUB_TOKEN")
	if token == "" {
		log.Fatal("Please define github token as your environment variable.\nGOST_GITHUB_TOKEN: xxxxx")
	}
	return &token
}

// GetClient githubを扱うclient
func GetClient() *github.Client {
	token := getToken()

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: *token})
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}

// GetContext contextを取得する
func GetContext() context.Context {
	return ctx
}
