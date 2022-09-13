package {{.BizPkg}}

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type {{.ModelName}} struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	// TODO: : add field for your own business
	{{if .SoftDeleteFeature}}
	IsDeleted bool      `bson:"is_deleted" json:"is_deleted"`
	DeletedAt time.Time `bson:"deleted_at" json:"deleted_at"`
	{{end}}
    UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
    CreatedAt time.Time `bson:"created_at" json:"created_at"`

}

type {{.ModelName}}Provider interface {
	CreateOne{{.ModelName}}(ctx context.Context, data *{{.ModelName}}, opts ...*options.InsertOneOptions) (result *mongo.InsertOneResult, err error)
	CreateMany{{.ModelName}}(ctx context.Context, data []interface{}, opts ...*options.InsertManyOptions) (result *mongo.InsertManyResult, err error)
	FindOne{{.ModelName}}(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) (result *{{.ModelName}}, err error)
	FindOne{{.ModelName}}ById(ctx context.Context, id string, opts ...*options.FindOneOptions) (result *{{.ModelName}}, err error)
	Find{{.ModelName}}(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (result []*{{.ModelName}}, err error)
	Estimated{{.ModelName}}Count(ctx context.Context, opts ...*options.EstimatedDocumentCountOptions) (int64, error)
	GetMatch{{.ModelName}}Count(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error)
	UpdateOne{{.ModelName}}(ctx context.Context, filter interface{}, data interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult
	UpdateOne{{.ModelName}}ById(ctx context.Context, id string, data interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult
	Update{{.ModelName}}(ctx context.Context, filter interface{}, data interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteOne{{.ModelName}}(ctx context.Context, filter interface{}, opts ...*options.FindOneAndDeleteOptions) *mongo.SingleResult
	DeleteOne{{.ModelName}}ById(ctx context.Context, id string, opts ...*options.FindOneAndDeleteOptions) *mongo.SingleResult
	Delete{{.ModelName}}(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	{{if .SoftDeleteFeature}}
	SoftDeleteOne{{.ModelName}}(ctx context.Context, filter interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult
	SoftDeleteOne{{.ModelName}}ById(ctx context.Context, id string, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult
	SoftDelete{{.ModelName}}(ctx context.Context, filter interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	{{- end}}
}
