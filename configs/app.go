package configs

import (
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type AppConfig interface {
	GetMongoDbConnection() *mongo.Database
}

type appConfigs struct {
	DbConfigs `mapstructure:",squash"`
}

func NewAppConfigs() (configs *appConfigs) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&configs); err != nil {
		log.Fatal(err)
	}

	return
}
