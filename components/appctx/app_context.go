package appctx

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppContext interface {
	GetDbConnection() *mongo.Database
	GetValidator() *validator.Validate
}

type appContext struct {
	db        *mongo.Database
	validator *validator.Validate
}

func NewAppContext(db *mongo.Database, validator *validator.Validate) *appContext {
	return &appContext{db: db, validator: validator}
}

func (appCtx *appContext) GetDbConnection() *mongo.Database {
	return appCtx.db
}

func (appCtx *appContext) GetValidator() *validator.Validate {
	return appCtx.validator
}
