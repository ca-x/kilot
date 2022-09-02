package kratos

import (
	"github.com/urfave/cli/v2"
)

func Cmd() []*cli.Command {
	return []*cli.Command{
		{
			Name:  "mongo",
			Usage: "do mongodb code generation for kratos.",
			Subcommands: []*cli.Command{
				{
					Name:    "generate",
					Aliases: []string{"g"},
					Usage:   "start code generation.please pass at least one arg as model name(multi model names are supported).",
					Flags:   kratosOptions(),
					Action:  processGenerate},
			},
		},
	}
}

func processGenerate(cCtx *cli.Context) error {
	return nil
}
