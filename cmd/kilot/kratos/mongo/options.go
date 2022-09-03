package mongo

import (
	"github.com/czyt/kilot/cmd/kilot/kratos/internal/templateContext"
	"github.com/urfave/cli/v2"
)

const (
	userOdmFlag     = "use-odm"
	formatFlag      = "format"
	modelSuffixFlag = "model-suffix"
	modelPrefixFlag = "model-prefix"
	modelNamesFlag  = "model-names"
)

var (
	tplContext templateContext.MongoContext
)

func Options() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:        userOdmFlag,
			Value:       true,
			Usage:       "this flag set whether to use mongo odm (use mgm).",
			Destination: &tplContext.UseOdm,
		},
		&cli.BoolFlag{
			Name:        formatFlag,
			Aliases:     []string{"f"},
			Value:       true,
			Usage:       "set whether to format generated code before write to file.",
			Destination: &tplContext.FormatCode,
		},
		&cli.StringFlag{
			Name:        modelSuffixFlag,
			Value:       "",
			Usage:       "set model name suffix `Suffix` .default is empty.",
			Destination: &tplContext.ModelSuffix,
		},
		&cli.StringFlag{
			Name:        modelPrefixFlag,
			Value:       "",
			Usage:       "set model name prefix `Prefix`.default is empty.",
			Destination: &tplContext.ModelPrefix,
		},
		&cli.StringSliceFlag{
			Name:        modelNamesFlag,
			Aliases:     []string{"mn"},
			Value:       nil,
			Usage:       "set model names `ModelName` for mongo code generate.multi model names supported.",
			Required:    true,
			Destination: &tplContext.ModelNames,
		},
	}
}
