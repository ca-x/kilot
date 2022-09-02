package main

import (
	"github.com/czyt/kilot/cmd/kilot/kratos"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"sort"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:        "kratos",
				Aliases:     []string{"k"},
				Usage:       "generate code for go-kratos framework",
				Subcommands: kratos.Cmd(),
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
