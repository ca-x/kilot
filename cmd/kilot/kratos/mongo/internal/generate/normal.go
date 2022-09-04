package generate

import (
	"github.com/czyt/kilot/cmd/kilot/kratos/internal/templateContext"
	"github.com/czyt/kilot/cmd/kilot/kratos/mongo/internal/templates"
	"github.com/czyt/kilot/cmd/kilot/pkg/templator"
	"log"
)

var (
	normalBizCoder  = templator.New("normalBizCoder", templates.NormalBizLayerLogicTpl, true)
	normalDataCoder = templator.New("normalDataCoder", templates.NormalDataLayerLogicTpl, true)
)

const (
	normalCodeGenerationCompletedPrompt = `
 You should add your code for initDB in somewhere else.
similar to the code below.generated code needs a *Database
object to work.

    const uri = "mongodb://user:pass@sample.host:27017/?maxPoolSize=20&w=majority"
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
    database := client.Database(<DATABASE NAME>)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()`
)

func NormalCodeWithCtx(ctx templateContext.MongoContext) error {
	if err := doCodeGenerationWith(ctx, normalDataCoder, normalBizCoder); err != nil {
		return err
	}
	log.Println(normalCodeGenerationCompletedPrompt)
	return nil
}
