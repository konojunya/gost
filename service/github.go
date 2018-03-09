package service

import (
	"log"

	"github.com/google/go-github/github"
	"github.com/konojunya/gost/auth"
)

// CreateGist Gistを作成する
func CreateGist(gist *github.Gist) string {
	resGist, _, err := client.Gists.Create(auth.GetContext(), gist)
	if err != nil {
		log.Fatal(err)
	}

	return resGist.GetHTMLURL()
}
