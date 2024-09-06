package main

import (
	"alucard/internal/config"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	cfg, err_cli := config.LoadConfig("internal/cli/config.yml")

	if err_cli != nil {
		log.Fatal(err_cli)
		return
	}

	app := &cli.App{
		Name:      cfg.AppName,
		UsageText: cfg.Description,
		Authors: []*cli.Author{
			{
				Name:  cfg.Author.Name,
				Email: cfg.Author.Email,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
