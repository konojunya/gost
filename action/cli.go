package action

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/konojunya/go-frame"

	"github.com/konojunya/gost/utils"

	"github.com/google/go-github/github"
	"github.com/konojunya/gost/service"

	"github.com/konojunya/gost/auth"
	"github.com/konojunya/gost/model"
	"github.com/konojunya/gost/server"
	"github.com/urfave/cli"
)

func loadConfig(file string) (*model.Config, error) {
	var config model.Config

	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(b, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// Login github login with oauth
func Login(c *cli.Context) {
	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}
	auth.SetOAuthConfig(*config)
	server.Listen()
}

// CreateGist upload gist your github
func CreateGist(c *cli.Context) {
	if !auth.IsAuthedUser() {
		fmt.Println("you are not authorize yet :)\n$ gost login")
		return
	}

	file := c.Args().Get(0)
	if len(file) == 0 {
		fmt.Println("Please input filepath\n\nUsage:\n$ gost create /path/to/file")
		os.Exit(0)
	}
	found := utils.Exists(file)
	if !found {
		log.Fatal(fmt.Errorf("%v is not found", file))
	}

	body := utils.GetFile(file)
	description := c.String("message")
	public := !c.Bool("private")
	gist := &github.Gist{
		Description: &description,
		Public:      &public,
		Files: map[github.GistFilename]github.GistFile{
			github.GistFilename(file): github.GistFile{
				Content: github.String(body),
			},
		},
	}

	gistURL, res, err := service.CreateGist(gist)
	if res.StatusCode == http.StatusUnauthorized {
		fmt.Println("Your token is expired\nPlease login.\n\n$ gost login")
		os.Exit(0)
	}
	if err != nil {
		log.Fatal(err)
	}
	result := model.Result{
		GistURL:     gistURL,
		Description: description,
		Public:      public,
		FilePath:    file,
	}

	frame.Print(result)

}
