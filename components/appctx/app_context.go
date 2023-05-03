package appctx

import (
	"dev_community_server/common"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppContext interface {
	GetDbConnection() *mongo.Database
	GetValidator() common.Validator
}

type appContext struct {
	db        *mongo.Database
	validator common.Validator
}

func NewAppContext(db *mongo.Database, validator common.Validator) *appContext {
	return &appContext{db: db, validator: validator}
}

func (appCtx *appContext) GetDbConnection() *mongo.Database {
	return appCtx.db
}

func (appCtx *appContext) GetValidator() common.Validator {
	return appCtx.validator
}
