package data

import (
	"context"
	"github.com/Kamva/mgm"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"biz"
)

type {{.ModelNameLowCase}}DataRepo struct {
	model mgm.Model
	// Todo: add your data and loggers here
	// data *Data
}

func (d {{.ModelNameLowCase}}DataRepo) Create{{.ModelName}}(ctx context.Context, data *biz.{{.ModelName}}, opts ...*options.InsertOneOptions) error {
	return mgm.Coll(d.model).CreateWithCtx(ctx, data, opts...)
}

func (d {{.ModelNameLowCase}}DataRepo) FindOne{{.ModelName}}(ctx context.Context, id string, opts ...*options.FindOneOptions) (result *biz.{{.ModelName}}, err error) {
	recordId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": recordId}
	err = mgm.Coll(d.model).FindOne(ctx, filter, opts...).Decode(result)
	if err != nil {
		return nil, err
	}
	return
}

func (d {{.ModelNameLowCase}}DataRepo) FindAll{{.ModelName}}(ctx context.Context, opts ...*options.FindOptions) (result []*biz.{{.ModelName}}, err error) {
	filter := bson.M{}
	cursor, err := mgm.Coll(d.model).Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &result)
	if err != nil {
		return nil, err
	}
	return
}

func (d {{.ModelNameLowCase}}DataRepo) Update{{.ModelName}}(ctx context.Context, id string, data *biz.{{.ModelName}}, opts ...*options.FindOneAndUpdateOptions) error {
	recordId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": recordId}
	if updateResult := mgm.Coll(d.model).FindOneAndUpdate(ctx, filter, data, opts...); updateResult.Err() != nil {
		return updateResult.Err()
	}
	return nil
}

func (d {{.ModelNameLowCase}}DataRepo) Delete{{.ModelName}}(ctx context.Context, id string, opts ...*options.FindOneAndDeleteOptions) error {
	recordId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": recordId}
	if deleteResult := mgm.Coll(d.model).FindOneAndDelete(ctx, filter, opts...); deleteResult.Err() != nil {
		return deleteResult.Err()
	}
	return nil
}

func New{{.ModelName}}DataRepo() biz.{{.ModelName}}Provider {
	return &{{.ModelNameLowCase}}DataRepo{
		model: &biz.{{.ModelName}}{},
	}
}
