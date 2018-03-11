package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"

	"github.com/google/go-github/github"
	"github.com/konojunya/go-frame"

	"github.com/konojunya/gost/service"
	"github.com/konojunya/gost/utils"
)

var (
	f     = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	usage = "\n\nUsage:\n$ gost /path/to/file [options]:\n-m: description\n-private: private gist"
)

type Result struct {
	GistURL     string `frame:"Gist's URL"`
	Description string
	Public      bool   `frame:"Is the published gist?"`
	FilePath    string `frame:"Local file path"`
}

func main() {
	description := f.String("m", "", "Gist Description")
	private := f.Bool("private", false, "Gist created")
	f.Parse(os.Args[1:])
	for 0 < f.NArg() {
		f.Parse(f.Args()[1:])
	}

	if len(os.Args) == 1 {
		fmt.Println("Please input upload filepath." + usage)
		return
	}

	filepath := os.Args[1]
	if !utils.Exists(filepath) {
		fmt.Println("file not found... :(")
		return
	}

	rep := regexp.MustCompile(`/`)
	re := rep.Split(filepath, -1)
	filename := re[len(re)-1]

	body := utils.GetFile(filepath)

	gist := &github.Gist{
		Description: description,
		Public:      inverted(private),
		Files: map[github.GistFilename]github.GistFile{
			github.GistFilename(filename): github.GistFile{
				Content: github.String(body),
			},
		},
	}

	gistURL := service.CreateGist(gist)

	result := Result{
		GistURL:     gistURL,
		Description: *description,
		Public:      *inverted(private),
		FilePath:    filepath,
	}

	frame.Print(result)
}

func inverted(b *bool) *bool {
	i := !*b
	return &i
}
