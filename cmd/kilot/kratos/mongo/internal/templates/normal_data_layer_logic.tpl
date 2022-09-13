package {{.DataPkg}}

import (
	"context"
	"github.com/Kamva/mgm/operator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"{{.BizPkg}}"
	"time"
)

const {{.ModelNameLowCase}}CollectionName = "{{.ModelName}}"

type {{.ModelNameLowCase}}DataRepo struct {

	// Todo: add your data and loggers here
	coll *mongo.Collection
}

func ({{.ModelIdentifier}} {{.ModelNameLowCase}}DataRepo) CreateOne{{.ModelName}}(ctx context.Context, data *{{.BizPkg}}.{{.ModelName}}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	if data.ID.IsZero() {
    		data.ID = primitive.NewObjectID()
    }
	return {{.ModelIdentifier}}.coll.InsertOne(ctx, data, opts...)
}

func ({{.ModelIdentifier}} {{.ModelNameLowCase}}DataRepo) CreateMany{{.ModelName}}(ctx context.Context, data []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	return {{.ModelIdentifier}}.coll.InsertMany(ctx, data, opts...)

}

func ({{.ModelIdentifier}} {{.ModelNameLowCase}}DataRepo) FindOne{{.ModelName}}(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) (*{{.BizPkg}}.{{.ModelName}}, error) {
	result := &{{.BizPkg}}.{{.ModelName}}{}
	if err := {{.ModelIdentifier}}.coll.FindOne(ctx, filter, opts...).Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}

func ({{.ModelIdentifier}} {{.ModelNameLowCase}}DataRepo) FindOne{{.ModelName}}ById(ctx context.Context, id string, opts ...*options.FindOneOptions) (*{{.BizPkg}}.{{.ModelName}}, error) {
	recordId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return {{.ModelIdentifier}}.FindOne{{.ModelName}}(ctx, bson.M{"_id": recordId}, opts...)
}

func ({{.ModelIdentifier}} {{.ModelNameLowCase}}DataRepo) Find{{.ModelName}}(ctx context.Context, filter interface{}, opts ...*options.FindOptions) ([]*{{.BizPkg}}.{{.ModelName}}, error) {
	result := make([]*{{.BizPkg}}.{{.ModelName}}, 0)
	cursor, err := {{.ModelIdentifier}}.coll.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func ({{.ModelIdentifier}} {{.ModelNameLowCase}}DataRepo) Estimated{{.ModelName}}Count(ctx context.Context, opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	return {{.ModelIdentifier}}.coll.EstimatedDocumentCount(ctx, opts...)
}

func ({{.ModelIdentifier}} {{.ModelNameLowCase}}DataRepo) GetMatch{{.ModelName}}Count(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return {{.ModelIdentifier}}.coll.CountDocuments(ctx, filter, opts...)
}

func ({{.ModelIdentifier}} {{.ModelNameLowCase}}DataRepo) UpdateOne{{.ModelName}}(ctx context.Context, filter interface{}, data interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult {
	return {{.ModelIdentifier}}.coll.FindOneAndUpdate(ctx, filter, data, opts...)
}

func ({{.ModelIdentifier}} {{.ModelNameLowCase}}DataRepo) UpdateOne{{.ModelName}}ById(ctx context.Context, id string, data interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult {
	recordId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil
	}
	return {{.ModelIdentifier}}.UpdateOne{{.ModelName}}(ctx, bson.M{"_id": recordId}, data, opts...)
}

func ({{.ModelIdentifier}} {{.ModelNameLowCase}}DataRepo) Update{{.ModelName}}(ctx context.Context, filter interface{}, data interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return {{.ModelIdentifier}}.coll.UpdateMany(ctx, filter, data, opts...)
}

func ({{.ModelIdentifier}} {{.ModelNameLowCase}}DataRepo) DeleteOne{{.ModelName}}(ctx context.Context, filter interface{}, opts ...*options.FindOneAndDeleteOptions) *mongo.SingleResult {
	return {{.ModelIdentifier}}.coll.FindOneAndDelete(ctx, filter, opts...)
}

func ({{.ModelIdentifier}} {{.ModelNameLowCase}}DataRepo) DeleteOne{{.ModelName}}ById(ctx context.Context, id string, opts ...*options.FindOneAndDeleteOptions) *mongo.SingleResult {
	recordId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil
	}
	return {{.ModelIdentifier}}.DeleteOne{{.ModelName}}(ctx, bson.M{"_id": recordId}, opts...)
}

func ({{.ModelIdentifier}} {{.ModelNameLowCase}}DataRepo) Delete{{.ModelName}}(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return {{.ModelIdentifier}}.coll.DeleteMany(ctx, filter, opts...)
}

{{if .SoftDeleteFeature}}
func ({{.ModelIdentifier}} {{.ModelNameLowCase}}DataRepo) SoftDeleteOne{{.ModelName}}(ctx context.Context, filter interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult {
	updateData := bson.M{operator.Set: bson.M{"is_deleted": true,
		"deleted_at": primitive.NewDateTimeFromTime(time.Now()),
	}}
	return {{.ModelIdentifier}}.UpdateOne{{.ModelName}}(ctx, filter, updateData, opts...)
}

func ({{.ModelIdentifier}} {{.ModelNameLowCase}}DataRepo) SoftDeleteOne{{.ModelName}}ById(ctx context.Context, id string, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult {
	recordId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil
	}
	return {{.ModelIdentifier}}.SoftDeleteOne{{.ModelName}}(ctx, bson.M{"_id": recordId}, opts...)
}

func ({{.ModelIdentifier}} {{.ModelNameLowCase}}DataRepo) SoftDelete{{.ModelName}}(ctx context.Context, filter interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return {{.ModelIdentifier}}.Update{{.ModelName}}(ctx, filter,
		bson.M{operator.Set: bson.M{"is_deleted": true,
			"deleted_at": primitive.NewDateTimeFromTime(time.Now()),
		}}, opts...)
}
{{- end}}

func New{{.ModelName}}DataRepo(db *mongo.Database) {{.BizPkg}}.{{.ModelName}}Provider {
	return &{{.ModelNameLowCase}}DataRepo{
		coll: db.Collection({{.ModelNameLowCase}}CollectionName),
	}
}
