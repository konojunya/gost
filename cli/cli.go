package cli

import (
	"github.com/konojunya/gost/action"
	"github.com/urfave/cli"
)

// Getapp cliアプリケーションの本体を返却する
func Getapp() *cli.App {
	app := cli.NewApp()
	config(app)
	app.Commands = getCommands()

	return app
}

func config(app *cli.App) {
	app.Name = "gost"
	app.Usage = "gost is a command line tool written in Go (Golang), easily uploading code to GitHub's Gist."
	app.Version = "1.0.0"
}

func getCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "login",
			Usage:  "GitHub login with OAuth",
			Action: action.Login,
		},
		{
			Name:      "create",
			Usage:     "Create Gist with your file",
			UsageText: "gost create <path/to/file> [options...]",
			Action:    action.CreateGist,
			Flags:     getFlags(),
		},
	}
}

func getFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "message, m",
			Usage: "Gist Description",
		},
		cli.BoolFlag{
			Name:  "private, p",
			Usage: "Gist's public",
		},
	}
}
