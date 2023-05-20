package main

import (
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	"dev_community_server/configs"
	"dev_community_server/middlewares"
	"dev_community_server/routes"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Bootstrap() {
	// Load Configs
	appConfigs := configs.NewAppConfigs()
	appCtx := appctx.NewAppContext(appConfigs, common.NewValidator())

	// Gin setup
	router := gin.Default()

	// Using cors
	router.Use(cors.Default())

	// Using custom recover middleware
	router.Use(middlewares.Recover(appCtx))

	// Routes setup
	routes.SetupRoutes(appCtx, router)

	// Start
	err := router.Run(fmt.Sprintf(":%v", *appConfigs.GetPort()))
	if err != nil {
		panic(err)
	}
}
