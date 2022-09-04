package generate

import (
	"fmt"
	"github.com/czyt/kilot/cmd/kilot/kratos/internal/templateContext"
	"github.com/czyt/kilot/cmd/kilot/kratos/mongo/internal/templates"
	"github.com/czyt/kilot/cmd/kilot/pkg/formater"
	"github.com/czyt/kilot/cmd/kilot/pkg/ioplus"
	"github.com/czyt/kilot/cmd/kilot/pkg/templator"
	"log"
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
	// check output dir
	if ctx.ModelOutputDir == "" || !ioplus.IsDirExist(ctx.ModelOutputDir) {
		ctx.ModelOutputDir = ioplus.GetWorkingDir()
	}
	// get all model name
	for _, model := range ctx.ModelNames.Value() {
		if err := genCode(odmBizCoder, model, "biz", ctx); err != nil {
			log.Printf("generate biz layer code failed:%s\n", err)
			continue
		}
		if err := genCode(odmDataCoder, model, "data", ctx); err != nil {
			log.Printf("generate data layer code failed:%s\n", err)
			continue
		}
	}
	return nil
}

func genCode(worker *templator.TemplateWorker, model string, layer string, ctx templateContext.MongoContext) error {
	ctxVal := createCtxVal(model, ctx.ModelPrefix, ctx.ModelSuffix)
	modelName := createModelName(model, ctx.ModelPrefix, ctx.ModelSuffix)

	codeBuff, err := worker.Execute(ctxVal)
	if err != nil {
		log.Printf("execute %s tpl for model %s failed:%v\n", layer, modelName, err)
		return err
	}
	log.Printf("write generated %s layer code for model %s to file.\n", layer, modelName)
	if err := writeSourceCodeFile(
		ioplus.FilePathFrom(
			ctx.ModelOutputDir,
			layer,
			fmt.Sprintf("%s.go", formater.From(modelName).Lower()),
		),
		codeBuff.Bytes(),
	); err != nil {
		log.Printf("generate %s layer code for model %s failed:%v\n", layer, modelName, err)
	}
	return nil
}

func createCtxVal(model string, modelPrefix string, modelSuffix string) map[string]interface{} {
	modelName := createModelName(model, modelPrefix, modelSuffix)
	return map[string]interface{}{
		"ModelName":        formater.From(modelName).ToCamel(),
		"ModelNameLowCase": formater.From(modelName).Untitle(),
		"ModelIdentifier":  formater.From(modelName).Untitle()[:1],
	}
}

func createModelName(model string, modelPrefix string, modelSuffix string) string {
	return fmt.Sprintf("%s%s%s", modelPrefix, model, modelSuffix)
}

func writeSourceCodeFile(fileName string, codePayload []byte) error {
	if err := ioplus.EnsureFileDir(fileName); err != nil {
		log.Printf("create Dir failed:%v", err)
		return err
	}
	if err := ioplus.SaveFile(fileName, codePayload); err != nil {
		log.Printf("write to file:%s failed:%v", fileName, err)
		return err
	}
	return nil
}
