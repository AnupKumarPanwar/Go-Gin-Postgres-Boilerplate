package router

import (
	"github.com/gin-gonic/gin"
	"rightswipe/config"
	"rightswipe/controllers"
)

func Init() {
	serverConfig := config.GetConfig().Server
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)
	router.GET("/ping", health.Status)

	router.Run(serverConfig.Host + ":" + serverConfig.Port)
}
