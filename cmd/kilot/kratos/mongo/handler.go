package mongo

import (
	"errors"
	"github.com/czyt/kilot/cmd/kilot/kratos/mongo/internal/generate"
	"github.com/urfave/cli/v2"
)

func CodeGenerator(*cli.Context) error {
	// check model names
	if len(tplContext.ModelNames.Value()) == 0 {
		return errors.New("please provide at least one model name use `-mn` option")
	}
	if tplContext.UseOdm {
		return generate.OdmCodeWithCtx(tplContext)
	}
	return generate.NormalCodeWithCtx(tplContext)
}
