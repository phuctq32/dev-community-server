package appctx

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type AppContext interface {
	GetDbConnection() *mongo.Database
}

type appContext struct {
	db *mongo.Database
}

func NewAppContext(db *mongo.Database) *appContext {
	return &appContext{db: db}
}

func (appCtx *appContext) GetDbConnection() *mongo.Database {
	return appCtx.db
}
