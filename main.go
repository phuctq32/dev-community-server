package main

import (
	"dev_community_server/components/appctx"
	"dev_community_server/configs"
)

func main() {
	appConfigs := configs.NewAppConfigs()
	appctx.NewAppContext(appConfigs.GetMongoDbConnection())
}
