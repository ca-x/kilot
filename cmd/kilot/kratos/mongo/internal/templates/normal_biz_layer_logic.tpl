package {{.BizPkg}}

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type {{.ModelName}} struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	// TODO: : add field for your own business
	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}

type {{.ModelName}}Provider interface {
	Create{{.ModelName}}(ctx context.Context, data *{{.ModelName}}, opts ...*options.InsertOneOptions) error
	FindOne{{.ModelName}}(ctx context.Context, id string, opts ...*options.FindOneOptions) (result *{{.ModelName}}, err error)
	FindAll{{.ModelName}}s(ctx context.Context, opts ...*options.FindOptions) (result []*{{.ModelName}}, err error)
	Update{{.ModelName}}(ctx context.Context, id string, data *{{.ModelName}}, opts ...*options.FindOneAndUpdateOptions) error
	Delete{{.ModelName}}(ctx context.Context, id string, opts ...*options.FindOneAndDeleteOptions) error
}
