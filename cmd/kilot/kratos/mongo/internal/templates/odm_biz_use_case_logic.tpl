package {{.BizPkg}}

import (
	"context"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type {{.UseCaseName}}UseCase struct {
    log    *log.Helper
}


func New{{.UseCaseName}}UseCase(logger log.Logger) {{.UseCaseName}}UseCase {
	return &{{.UseCaseName}}UseCase{
	    log: log.NewHelper(log.With(logger, "module", "biz/{{.UseCaseName}}UseCase")),
	}
}

