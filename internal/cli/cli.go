package main

import (
	"alucard/internal/config"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

type CLIConfig struct {
	AppName     string `yaml:"app_name"`
	Description string `yaml:"description"`
	Author      Author `yaml:"author"`
}

type Author struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}

func main() {
	cfg, err_cli := config.LoadConfig[CLIConfig]("internal/cli/config.yml")

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
