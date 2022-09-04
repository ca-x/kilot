package biz

import (
	"context"
	"github.com/Kamva/mgm"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type {{.ModelName}}Provider interface {
	Create{{.ModelName}}(ctx context.Context, data *{{.ModelName}}, opts ...*options.InsertOneOptions) error
	FindOne{{.ModelName}}(ctx context.Context, id string, opts ...*options.FindOneOptions) (result *{{.ModelName}}, err error)
	FindAll{{.ModelName}}(ctx context.Context, opts ...*options.FindOptions) (result []*{{.ModelName}}, err error)
	Update{{.ModelName}}(ctx context.Context, id string, data *{{.ModelName}}, opts ...*options.FindOneAndUpdateOptions) error
	Delete{{.ModelName}}(ctx context.Context, id string, opts ...*options.FindOneAndDeleteOptions) error
}

type {{.ModelName}} struct {
	mgm.DefaultModel `bson:",inline"`
	// TODO: add field for your own business
}
