package main

import (
	"log"
	"os"

	"github.com/hinshun/genplugin/action"
	"gopkg.in/urfave/cli.v2"
)

func main() {
	app := &cli.App{
		Name:   "genplugin",
		Usage:  "generates docker plugins",
		Action: action.GeneratePlugin,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
