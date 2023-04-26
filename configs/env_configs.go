package configs

import (
	"github.com/spf13/viper"
	"log"
)

type DbConfigs struct {
	MongoUsername string `mapstructure:"MONGO_USERNAME"`
	MongoPassword string `mapstructure:"MONGO_PASSWORD"`
	MongoDB       string `mapstructure:"MONGO_DB_NAME"`
}

type envConfigs struct {
	DbConfigs `mapstructure:",squash"`
}

var EnvConfigs *envConfigs

func InitEnvConfigs() {
	var configs *envConfigs
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&configs); err != nil {
		log.Fatal(err)
	}

	EnvConfigs = configs
}
