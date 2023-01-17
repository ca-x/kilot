package generate

import (
	"fmt"
	"github.com/czyt/kilot/cmd/kilot/kratos/internal/templateContext"
	"github.com/czyt/kilot/cmd/kilot/pkg/formater"
	"github.com/czyt/kilot/cmd/kilot/pkg/ioplus"
	"github.com/czyt/kilot/cmd/kilot/pkg/templator"
	"log"
)

func doCodeGenerationWith(ctx templateContext.MongoContext, dataWorker, bizWorker *templator.TemplateWorker) error {
	if !ctx.FormatCode {
		dataWorker.ChangeCodeFormatOption(false)
		bizWorker.ChangeCodeFormatOption(false)
	}
	// check output dir
	if ctx.ModelOutputDir == "" || !ioplus.IsDirExist(ctx.ModelOutputDir) {
		ctx.ModelOutputDir = ioplus.GetWorkingDir()
	}

	// get all model name
	for _, model := range ctx.ModelNames.Value() {
		if err := genCode(bizWorker, model, ctx.BizLayerCodeDir, ctx); err != nil {
			log.Printf("generate biz layer code failed:%s\n", err)
			continue
		}
		if err := genCode(dataWorker, model, ctx.DataLayerCodeDir, ctx); err != nil {
			log.Printf("generate data layer code failed:%s\n", err)
			continue
		}
	}

	if ctx.UseCase != "" {
		// generate UseCase code
		useCaseDataContext := make(map[string]any, 4)
		useCaseDataContext["UseCaseName"] = formater.From(ctx.UseCase).ToCamel()
		useCaseDataContext["UseCaseIdentifier"] = formater.From(ctx.UseCase).Untitle()[:1]
		worker := normalUseCaseCoder
		if ctx.UseOdm {
			worker = odmUseCaseCoder
		}
		if err := genCode(worker, ctx.UseCase, ctx.BizLayerCodeDir, ctx); err != nil {
			log.Printf("generate useCase layer code failed:%s\n", err)
			return err
		}

	}
	return nil
}

func genCode(worker *templator.TemplateWorker, model string, layerDir string, ctx templateContext.MongoContext) error {
	ctxVal := createCtxVal(model, ctx)
	modelName := createModelName(model, ctx.ModelPrefix, ctx.ModelSuffix)

	codeBuff, err := worker.Execute(ctxVal)
	if err != nil {
		log.Printf("execute %s tpl for pkg %s failed:%v\n", layerDir, modelName, err)
		return err
	}
	log.Printf("write generated %s layer code for pkg %s to file.\n", layerDir, modelName)
	if err := writeSourceCodeFile(
		ioplus.FilePathFrom(
			ctx.ModelOutputDir,
			layerDir,
			fmt.Sprintf("%s.go", formater.From(modelName).Lower()),
		),
		codeBuff.Bytes(),
	); err != nil {
		log.Printf("generate %s layer code for Pkg %s failed:%v\n", layerDir, modelName, err)
	}
	return nil
}

func createCtxVal(model string, ctx templateContext.MongoContext) map[string]interface{} {
	modelName := createModelName(model, ctx.ModelPrefix, ctx.ModelSuffix)
	return map[string]interface{}{
		"SoftDeleteFeature":  ctx.WithSoftDelete,
		"ModelName":          formater.From(modelName).ToCamel(),
		"ModelNameSnackCase": formater.From(modelName).ToSnake(),
		"ModelNameLowCase":   formater.From(modelName).Untitle(),
		"ModelIdentifier":    formater.From(modelName).Untitle()[:1],
		"BizPkg":             formater.From(ctx.BizLayerCodeDir).Lower(),
		"DataPkg":            formater.From(ctx.DataLayerCodeDir).Lower(),
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
