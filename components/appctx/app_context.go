package appctx

import (
	"dev_community_server/common"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppContext interface {
	GetDbConnection() *mongo.Database
	GetValidator() common.Validator
	GetSecretKey() string
}

type appContext struct {
	db        *mongo.Database
	validator common.Validator
	secretKey string
}

func NewAppContext(db *mongo.Database, validator common.Validator, secretKey string) *appContext {
	return &appContext{db: db, validator: validator, secretKey: secretKey}
}

func (appCtx *appContext) GetDbConnection() *mongo.Database {
	return appCtx.db
}

func (appCtx *appContext) GetValidator() common.Validator {
	return appCtx.validator
}

func (appCtx *appContext) GetSecretKey() string {
	return appCtx.secretKey
}
