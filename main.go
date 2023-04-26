package main

import (
	"dev_community_server/components/appctx"
	"dev_community_server/configs"
	"dev_community_server/middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	appConfigs := configs.NewAppConfigs()
	appCtx := appctx.NewAppContext(appConfigs.GetMongoDbConnection())

	router := gin.Default()
	router.Use(middlewares.Recover(appCtx))
	router.Use()

	router.GET("/", func(context *gin.Context) {
		context.String(200, "Ping!!!")
	})

	err := router.Run(fmt.Sprintf(":%v", appConfigs.Port))
	if err != nil {
		log.Fatal(err)
	}
}
