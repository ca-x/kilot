package generate

import (
	"fmt"
	"github.com/czyt/kilot/cmd/kilot/kratos/internal/templateContext"
	"github.com/czyt/kilot/cmd/kilot/kratos/mongo/internal/templates"
	"github.com/czyt/kilot/cmd/kilot/pkg/formater"
	"github.com/czyt/kilot/cmd/kilot/pkg/templator"
)

var (
	odmBizCoder  = templator.New("odmBizCoder", templates.OdmBizTpl, true)
	odmDataCoder = templator.New("odmDataCoder", templates.OdmDataTpl, true)
)

func OdmCodeWithCtx(ctx templateContext.MongoContext) error {
	if !ctx.FormatCode {
		odmBizCoder.ChangeCodeFormatOption(false)
		odmDataCoder.ChangeCodeFormatOption(false)
	}
	// get all model name
	for _, model := range ctx.ModelNames.Value() {
		modelName := fmt.Sprintf("%s%s%s", ctx.ModelPrefix, model, ctx.ModelSuffix)
		ctxVal := map[string]interface{}{
			"ModelName":        formater.From(modelName).ToCamel(),
			"ModelNameLowCase": formater.From(modelName).Untitle(),
		}
		fmt.Println(fmt.Sprintf("start to generate  model:%s biz code.", modelName))
		bizCodeBuff, err := odmBizCoder.Execute(ctxVal)
		if err != nil {
			continue
		}
		fmt.Println(bizCodeBuff.String())
		dataCodeBuff, err := odmDataCoder.Execute(ctxVal)
		if err != nil {
			continue
		}
		fmt.Println(dataCodeBuff.String())
	}
	return nil
}
