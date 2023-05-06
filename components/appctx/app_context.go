package appctx

import (
	"dev_community_server/common"
	"dev_community_server/configs"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppContext interface {
	GetMongoDbConnection() *mongo.Database
	GetValidator() common.Validator
	GetSecretKey() *string
	GetSendGridConfigs() configs.SendGridConfigs
}

type appContext struct {
	configs   configs.AppConfig
	validator common.Validator
}

func NewAppContext(configs configs.AppConfig, validator common.Validator) *appContext {
	return &appContext{configs: configs, validator: validator}
}

func (appCtx *appContext) GetMongoDbConnection() *mongo.Database {
	return appCtx.configs.GetMongoDbConnection()
}

func (appCtx *appContext) GetValidator() common.Validator {
	return appCtx.validator
}

func (appCtx *appContext) GetSecretKey() *string {
	return appCtx.configs.GetSecretKey()
}

func (appCtx *appContext) GetSendGridConfigs() configs.SendGridConfigs {
	return appCtx.configs
}
