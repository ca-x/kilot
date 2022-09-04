package generate

import (
	"github.com/czyt/kilot/cmd/kilot/kratos/internal/templateContext"
	"github.com/czyt/kilot/cmd/kilot/kratos/mongo/internal/templates"
	"github.com/czyt/kilot/cmd/kilot/pkg/templator"
)

var (
	normalBizCoder  = templator.New("normalBizCoder", templates.NormalBizLayerLogicTpl, true)
	normalDataCoder = templator.New("normalDataCoder", templates.NormalDataLayerLogicTpl, true)
)

func NormalCodeWithCtx(ctx templateContext.MongoContext) error {
	return doCodeGenerationWith(ctx, normalDataCoder, normalBizCoder)
}
