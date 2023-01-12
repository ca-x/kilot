package main

import (
	"github.com/czyt/kilot/cmd/kilot/kratos"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"sort"
)

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "print tool version.(#^.^#)",
	}
	app := &cli.App{
		Authors: []*cli.Author{
			{
				Name:  "czyt",
				Email: "czyt@w.cn",
			},
		},
		Name:                 "kilot",
		Version:              "0.1.11",
		EnableBashCompletion: true,
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
