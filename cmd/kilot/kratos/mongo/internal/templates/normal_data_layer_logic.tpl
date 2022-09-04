package data

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"biz"
)

const {{.ModelNameLowCase}}CollectionName = "{{.ModelNameSnackCase}}"

type {{.ModelNameLowCase}}DataRepo struct {

	// Todo: add your data and loggers here
	coll *mongo.Collection
}

func ({{.ModelIdentifier}} {{.ModelNameLowCase}}) Create{{.ModelName}}(ctx context.Context, data *biz.{{.ModelName}}, opts ...*options.InsertOneOptions) error {
	_, err := d.coll.InsertOne(ctx, data, opts...)
	if err != nil {
		return err
	}
	return nil
}

func ({{.ModelIdentifier}} {{.ModelNameLowCase}}) FindOne{{.ModelName}}(ctx context.Context, id string, opts ...*options.FindOneOptions) (result *biz.{{.ModelName}}, err error) {
	recordId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": recordId}
	err = d.coll.FindOne(ctx, filter, opts...).Decode(result)
	if err != nil {
		return nil, err
	}
	return
}

func ({{.ModelIdentifier}} {{.ModelNameLowCase}}) FindAll{{.ModelName}}s(ctx context.Context, opts ...*options.FindOptions) (result []*biz.{{.ModelName}}, err error) {
	filter := bson.M{}
	cursor, err := d.coll.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &result)
	if err != nil {
		return nil, err
	}
	return
}

func ({{.ModelIdentifier}} {{.ModelNameLowCase}}) Update{{.ModelName}}(ctx context.Context, id string, data *biz.{{.ModelName}}, opts ...*options.FindOneAndUpdateOptions) error {
	recordId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": recordId}
	if updateResult := d.coll.FindOneAndUpdate(ctx, filter, data, opts...); updateResult.Err() != nil {
		return updateResult.Err()
	}
	return nil
}

func ({{.ModelIdentifier}} {{.ModelNameLowCase}}) Delete{{.ModelName}}(ctx context.Context, id string, opts ...*options.FindOneAndDeleteOptions) error {
	recordId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": recordId}
	if deleteResult := d.coll.FindOneAndDelete(ctx, filter, opts...); deleteResult.Err() != nil {
		return deleteResult.Err()
	}
	return nil
}

func New{{.ModelName}}DataRepo(db *mongo.Database) biz.{{.ModelName}}Provider {
	return &{{.ModelNameLowCase}}DataRepo{
		coll: db.Collection({{.ModelNameLowCase}}CollectionName),
	}
}
