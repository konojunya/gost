package auth

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/go-github/github"

	"github.com/konojunya/gost/utils"

	"golang.org/x/oauth2"

	"github.com/konojunya/gost/model"
)

var (
	oauth2Config   *oauth2.Config
	configFilePath = os.Getenv("HOME") + "/.gost"
	ctx            = context.Background()
)

func GetOAuthConfig() *oauth2.Config {
	return oauth2Config
}

func SetOAuthConfig(config model.Config) {
	oauth2Config = &oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  config.AuthURL,
			TokenURL: config.TokenURL,
		},
		RedirectURL: config.RedirectURL,
		Scopes:      []string{"gist"},
	}
}

func GetToken() (*oauth2.Token, error) {
	raw, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}
	var config *oauth2.Token
	err = json.Unmarshal(raw, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func CreateTokenFile(token *oauth2.Token) error {
	json, err := json.Marshal(token)
	if err != nil {
		log.Fatal(err)
	}
	return ioutil.WriteFile(configFilePath, json, os.ModePerm)
}

func IsAuthedUser() bool {
	return utils.Exists(configFilePath)
}

func GetClient() *github.Client {
	token, err := GetToken()
	if err != nil {
		log.Fatal(err)
	}
	ts := oauth2.StaticTokenSource(token)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}

func GetContext() context.Context {
	return ctx
}
