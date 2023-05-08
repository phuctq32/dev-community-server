package configs

import (
	"github.com/spf13/viper"
	"log"
)

type AppConfig interface {
	GetMongoDbConfig() MongoDBConfig
	GetPort() *int
	GetSecretKey() *string
	GetSendGridConfig() SendGridConfig
	GetCloudinaryConfig() CloudinaryConfig
}

type appConfigs struct {
	mongoDBConfig    `mapstructure:",squash"`
	sendgridConfig   `mapstructure:",squash"`
	cloudinaryConfig `mapstructure:",squash"`
	Port             int    `mapstructure:"PORT"`
	SecretKey        string `mapstructure:"SECRET_KEY"`
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

func (config *appConfigs) GetMongoDbConfig() MongoDBConfig {
	return &config.mongoDBConfig
}

func (config *appConfigs) GetSendGridConfig() SendGridConfig {
	return &config.sendgridConfig
}

func (config *appConfigs) GetCloudinaryConfig() CloudinaryConfig {
	return &config.cloudinaryConfig
}
