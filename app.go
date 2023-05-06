package main

import (
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	"dev_community_server/configs"
	"dev_community_server/middlewares"
	"dev_community_server/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Bootstrap() {
	// Load Configs
	appConfigs := configs.NewAppConfigs()
	appCtx := appctx.NewAppContext(appConfigs, common.NewValidator())

	// Gin setup
	router := gin.Default()
	router.Use(middlewares.Recover(appCtx))
	routes.SetupRoutes(appCtx, router)

	// Start
	err := router.Run(fmt.Sprintf(":%v", *appConfigs.GetPort()))
	if err != nil {
		panic(err)
	}
}
