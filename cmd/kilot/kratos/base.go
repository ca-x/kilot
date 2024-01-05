package kratos

import (
	"github.com/ca-x/kilot/cmd/kilot/kratos/mongo"
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
					Flags:   mongo.Options(),
					Action:  mongo.CodeGenerator,
				},
			},
		},
	}
}
