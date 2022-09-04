package templateContext

import "github.com/urfave/cli/v2"

type MongoContext struct {
	FormatCode       bool
	UseOdm           bool
	ModelSuffix      string
	ModelPrefix      string
	ModelOutputDir   string
	BizLayerCodeDir  string
	DataLayerCodeDir string
	ModelNames       cli.StringSlice
}
