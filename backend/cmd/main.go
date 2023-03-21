package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/fabl3ss/banking_system/di"
	"github.com/fabl3ss/banking_system/pkg/cmd"
)

func main() {
	fxOptions := di.ProvideModules()

	app := &cli.App{
		Name: "backend",
		Commands: []*cli.Command{
			cmd.NewMigrationsCommand(fxOptions),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
