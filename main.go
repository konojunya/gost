package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	flag.Parse()
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		log.Fatal("Unauthorized: No token present")
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	body, err := ioutil.ReadFile("./test/index.html")
	if err != nil {
		log.Fatal(err)
	}

	gist, _, err := client.Gists.Create(ctx, &github.Gist{
		Description: github.String("hoge"),
		Public:      github.Bool(true),
		Files: map[github.GistFilename]github.GistFile{
			"index.html": github.GistFile{
				Content: github.String(string(body)),
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*gist.HTMLURL)
}
