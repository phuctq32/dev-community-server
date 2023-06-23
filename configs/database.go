package configs

type MongoDBConfig interface {
	GetMongoUsername() *string
	GetMongoPassword() *string
	GetMongoDbName() *string
}

type mongoDBConfig struct {
	MongoUsername string `mapstructure:"MONGO_USERNAME"`
	MongoPassword string `mapstructure:"MONGO_PASSWORD"`
	MongoDbName   string `mapstructure:"MONGO_DB_NAME"`
}

func (config *mongoDBConfig) GetMongoUsername() *string {
	return &config.MongoUsername
}

func (config *mongoDBConfig) GetMongoPassword() *string {
	return &config.MongoPassword
}

func (config *mongoDBConfig) GetMongoDbName() *string {
	return &config.MongoDbName
}
