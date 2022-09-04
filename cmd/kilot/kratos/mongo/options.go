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
	modelOutputFlag = "model-output-dir"
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
			Aliases:     []string{"s"},
			Value:       "",
			Usage:       "set model name suffix `Suffix` .default is empty.",
			Destination: &tplContext.ModelSuffix,
		},
		&cli.StringFlag{
			Name:        modelPrefixFlag,
			Aliases:     []string{"p"},
			Value:       "",
			Usage:       "set model name prefix `Prefix`.default is empty.",
			Destination: &tplContext.ModelPrefix,
		},
		&cli.StringFlag{
			Name:        modelOutputFlag,
			Aliases:     []string{"d"},
			Value:       "",
			Usage:       "set model save dir `Dir`.if not set tool working dir will be used.",
			Destination: &tplContext.ModelOutputDir,
		},
		&cli.StringSliceFlag{
			Name:        modelNamesFlag,
			Aliases:     []string{"n"},
			Value:       nil,
			Usage:       "set model names `ModelName` for mongo code generate.multi model names supported.",
			Required:    true,
			Destination: &tplContext.ModelNames,
		},
	}
}
