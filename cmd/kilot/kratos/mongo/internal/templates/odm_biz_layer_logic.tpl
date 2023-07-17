package {{.BizPkg}}

import (
	"context"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type {{.ModelName}}Provider interface {
    Aggregate{{.ModelName}}(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) (*mongo.Cursor, error)
    Create{{.ModelName}}(ctx context.Context, data *{{.ModelName}}, opts ...*options.InsertOneOptions) error
	CreateOne{{.ModelName}}(ctx context.Context, data *{{.ModelName}}, opts ...*options.InsertOneOptions) (result *mongo.InsertOneResult, err error)
	CreateOne{{.ModelName}}Index(ctx context.Context, index mongo.IndexModel, opts ...*options.CreateIndexesOptions) (result string, err error)
	CreateMany{{.ModelName}}(ctx context.Context, data []interface{}, opts ...*options.InsertManyOptions) (result *mongo.InsertManyResult, err error)
	CreateMany{{.ModelName}}Index(ctx context.Context, index []mongo.IndexModel, opts ...*options.CreateIndexesOptions) (result []string, err error)
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
    {{end}}
}

type {{.ModelName}} struct {
	mgm.DefaultModel `bson:",inline"`
	// TODO: add field for your own business
	{{if .SoftDeleteFeature}}
    IsDeleted bool      `bson:"is_deleted" json:"is_deleted"`
    DeletedAt time.Time `bson:"deleted_at" json:"deleted_at"`
    {{- end}}
}
