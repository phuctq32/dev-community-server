package configs

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type DbConfigs struct {
	MongoUsername string `mapstructure:"MONGO_USERNAME"`
	MongoPassword string `mapstructure:"MONGO_PASSWORD"`
	MongoDbName   string `mapstructure:"MONGO_DB_NAME"`
}

func (configs *appConfigs) GetMongoDbConnection() *mongo.Database {
	mongoUri := fmt.Sprintf(
		"mongodb+srv://%v:%v@cluster0.g528okd.mongodb.net/?retryWrites=true&w=majority",
		configs.MongoUsername,
		configs.MongoPassword,
	)

	opts := options.Client().ApplyURI(mongoUri)
	ctx, _ := context.WithTimeout(context.TODO(), time.Second*10)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB Connected")

	return client.Database(configs.MongoDbName)
}
