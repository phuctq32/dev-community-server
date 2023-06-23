package appctx

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/configs"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type AppContext interface {
	GetValidator() common.Validator
	GetAppConfig() configs.AppConfig
	GetSendGridConfig() configs.SendGridConfig
	GetCloudinaryConfig() configs.CloudinaryConfig
	GetMongoDBConnection() *mongo.Database
}

type appContext struct {
	db        *mongo.Database
	configs   configs.Config
	validator common.Validator
}

func NewAppContext(configs configs.Config, validator common.Validator) *appContext {
	return &appContext{configs: configs, validator: validator}
}

func (appCtx *appContext) GetAppConfig() configs.AppConfig {
	return appCtx.configs.GetAppConfig()
}

func (appCtx *appContext) GetValidator() common.Validator {
	return appCtx.validator
}

func (appCtx *appContext) GetSendGridConfig() configs.SendGridConfig {
	return appCtx.configs.GetSendGridConfig()
}

func (appCtx *appContext) GetCloudinaryConfig() configs.CloudinaryConfig {
	return appCtx.configs.GetCloudinaryConfig()
}

func (appCtx *appContext) GetMongoDBConnection() *mongo.Database {
	if appCtx.db != nil {
		return appCtx.db
	}

	mongoUri := fmt.Sprintf(
		"mongodb+srv://%v:%v@cluster0.g528okd.mongodb.net/?retryWrites=true&w=majority",
		*appCtx.configs.GetMongoDbConfig().GetMongoUsername(),
		*appCtx.configs.GetMongoDbConfig().GetMongoPassword(),
	)

	opts := options.Client().ApplyURI(mongoUri)
	ctx, _ := context.WithTimeout(context.TODO(), time.Second*10)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB Connected")

	appCtx.db = client.Database(*appCtx.configs.GetMongoDbConfig().GetMongoDbName())

	return appCtx.db
}
