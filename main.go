package main

import (
	"log"
	"os"

	"github.com/LintaoAmons/uuid-go/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:            "uuid",
		Description:     "A tool to generate uuid and copy it your clipboard automaticlly",
		HideHelpCommand: true,
		UsageText:       "uuid [count] -- count is set to 1 if not passed",
		Usage:           "uuid [count]",
		Action:          commands.UuidCommand.Action,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
