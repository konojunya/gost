package service

import (
	"github.com/google/go-github/github"
	"github.com/konojunya/gost/auth"
)

// CreateGist Gistを作成する
func CreateGist(gist *github.Gist) (string, *github.Response, error) {
	client := auth.GetClient()
	gist, res, err := client.Gists.Create(auth.GetContext(), gist)
	if err != nil {
		return "", res, err
	}
	return gist.GetHTMLURL(), res, nil
}
