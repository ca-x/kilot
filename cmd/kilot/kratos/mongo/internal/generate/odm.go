package generate

import (
	"github.com/ca-x/kilot/cmd/kilot/kratos/internal/templateContext"
	"github.com/ca-x/kilot/cmd/kilot/kratos/mongo/internal/templates"
	"github.com/ca-x/kilot/cmd/kilot/pkg/templator"
	"log"
)

var (
	odmBizCoder     = templator.New("odmBizCoder", templates.OdmBizLayerLogicTpl, true)
	odmDataCoder    = templator.New("odmDataCoder", templates.OdmDataLayerLogicTpl, true)
	odmUseCaseCoder = templator.New("odmUseCaseCoder", templates.OdmUseCaseLogicTpl, true)
)

const (
	odmCodeGenerationCompletedPrompt = `
 You should add your code for initDB in somewhere else.
similar to the code below.

    opt := options.Client().ApplyURI(<DB ADDRESS>)
	mgm.SetDefaultConfig(nil, <DATABASE NAME>, opt)

then the mgm odm will do the connection itself.`
)

func OdmCodeWithCtx(ctx templateContext.MongoContext) error {
	if err := doCodeGenerationWith(ctx, odmDataCoder, odmBizCoder); err != nil {
		return err
	}
	log.Println(odmCodeGenerationCompletedPrompt)
	return nil
}
