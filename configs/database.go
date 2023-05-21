package configs

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type MongoDBConfig interface {
	GetConnection() *mongo.Database
}

type mongoDBConfig struct {
	MongoUsername string `mapstructure:"MONGO_USERNAME"`
	MongoPassword string `mapstructure:"MONGO_PASSWORD"`
	MongoDbName   string `mapstructure:"MONGO_DB_NAME"`
	db            *mongo.Database
}

func (config *mongoDBConfig) GetConnection() *mongo.Database {
	if config.db != nil {
		return config.db
	}

	mongoUri := fmt.Sprintf(
		"mongodb+srv://%v:%v@cluster0.g528okd.mongodb.net/?retryWrites=true&w=majority",
		config.MongoUsername,
		config.MongoPassword,
	)

	opts := options.Client().ApplyURI(mongoUri)
	ctx, _ := context.WithTimeout(context.TODO(), time.Second*10)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB Connected")

	config.db = client.Database(config.MongoDbName)

	return config.db
}
