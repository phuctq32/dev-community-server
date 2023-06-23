package configs

import (
	"github.com/spf13/viper"
	"log"
)

type Config interface {
	GetAppConfig() AppConfig
	GetMongoDbConfig() MongoDBConfig
	GetSendGridConfig() SendGridConfig
	GetCloudinaryConfig() CloudinaryConfig
}

type AppConfig interface {
	GetPort() *int
	GetSecretKey() *string
}

type appConfig struct {
	Port      int    `mapstructure:"PORT"`
	SecretKey string `mapstructure:"SECRET_KEY"`
}

func (config *appConfig) GetPort() *int {
	return &config.Port
}

func (config *appConfig) GetSecretKey() *string {
	return &config.SecretKey
}

type configs struct {
	appConfig        `mapstructure:",squash"`
	mongoDBConfig    `mapstructure:",squash"`
	sendgridConfig   `mapstructure:",squash"`
	cloudinaryConfig `mapstructure:",squash"`
}

func NewConfigs() Config {
	var cfs *configs
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&cfs); err != nil {
		log.Fatal(err)
	}

	return cfs
}

func (config *configs) GetAppConfig() AppConfig {
	return &config.appConfig
}

func (config *configs) GetMongoDbConfig() MongoDBConfig {
	return &config.mongoDBConfig
}

func (config *configs) GetSendGridConfig() SendGridConfig {
	return &config.sendgridConfig
}

func (config *configs) GetCloudinaryConfig() CloudinaryConfig {
	return &config.cloudinaryConfig
}
