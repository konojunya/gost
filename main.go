package main

import (
	"os"

	"github.com/konojunya/gost/cli"
)

func main() {
	app := cli.Getapp()
	app.Run(os.Args)
}
