package templateContext

import "github.com/urfave/cli/v2"

type MongoContext struct {
	FormatCode       bool
	UseOdm           bool
	WithSoftDelete   bool
	ModelSuffix      string
	ModelPrefix      string
	ModelOutputDir   string
	BizLayerCodeDir  string
	DataLayerCodeDir string
	UseCase          string
	ModelNames       cli.StringSlice
}
