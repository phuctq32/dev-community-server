package main

import (
	"dev_community_server/components/appctx"
	"dev_community_server/configs"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	appConfigs := configs.NewAppConfigs()
	appctx.NewAppContext(appConfigs.GetMongoDbConnection())

	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.String(200, "Ping!!!")
	})

	err := router.Run(fmt.Sprintf(":%v", appConfigs.Port))
	if err != nil {
		log.Fatal(err)
	}
}
