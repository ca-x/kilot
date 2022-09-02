package kratos

import "github.com/urfave/cli/v2"

func kratosOptions() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:  "use-odm",
			Value: true,
			Usage: "this flag set whether to use mongo odm (use mgm).",
		},
		&cli.BoolFlag{
			Name:    "format",
			Aliases: []string{"f"},
			Value:   true,
			Usage:   "this flag set whether to format generated code before write to file.",
		},
		&cli.StringFlag{
			Name:  "model-suffix",
			Value: "",
			Usage: "this flag set model name suffix.default is empty.",
		},
	}
}
