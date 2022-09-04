package generate

import (
	"github.com/czyt/kilot/cmd/kilot/kratos/internal/templateContext"
	"github.com/czyt/kilot/cmd/kilot/kratos/mongo/internal/templates"
	"github.com/czyt/kilot/cmd/kilot/pkg/templator"
)

var (
	odmBizCoder  = templator.New("odmBizCoder", templates.OdmBizLayerLogicTpl, true)
	odmDataCoder = templator.New("odmDataCoder", templates.OdmDataLayerLogicTpl, true)
)

func OdmCodeWithCtx(ctx templateContext.MongoContext) error {
	return doCodeGenerationWith(ctx, odmDataCoder, odmBizCoder)
}
