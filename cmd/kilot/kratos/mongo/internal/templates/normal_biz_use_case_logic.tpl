package {{.BizPkg}}

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type {{.UseCaseName}}UseCase struct {
    log    *log.Helper
}


func New{{.UseCaseName}}UseCase(logger log.Logger) {{.UseCaseName}}UseCase {
	return &{{.UseCaseName}}UseCase{
	    log: log.NewHelper(log.With(logger, "module", "biz/{{.UseCaseName}}UseCase")),
	}
}