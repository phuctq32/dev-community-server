package configs

import (
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type AppConfig interface {
	GetMongoDbConnection() *mongo.Database
	GetPort() *int
	GetSecretKey() *string
}

type appConfigs struct {
	DbConfigs `mapstructure:",squash"`
	Port      int    `mapstructure:"PORT"`
	SecretKey string `mapstructure:"SECRET_KEY"`
}

func NewAppConfigs() AppConfig {
	var configs *appConfigs
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&configs); err != nil {
		log.Fatal(err)
	}

	return configs
}

func (config *appConfigs) GetPort() *int {
	return &config.Port
}

func (config *appConfigs) GetSecretKey() *string {
	return &config.SecretKey
}
