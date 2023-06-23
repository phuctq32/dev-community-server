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
	"time"
)

func Bootstrap() {
	// Load Configs
	appConfigs := configs.NewConfigs()
	appCtx := appctx.NewAppContext(appConfigs, common.NewValidator())

	// Gin setup
	router := gin.Default()

	// Using cors
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	// Using custom recover middleware
	router.Use(middlewares.Recover(appCtx))

	// Routes setup
	routes.SetupRoutes(appCtx, router)

	// Start
	err := router.Run(fmt.Sprintf(":%v", *appConfigs.GetAppConfig().GetPort()))
	if err != nil {
		panic(err)
	}
}
