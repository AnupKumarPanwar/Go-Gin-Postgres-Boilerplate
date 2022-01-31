package router

import (
	"Go-Gin-Postgres-Boilerplate/config"
	"Go-Gin-Postgres-Boilerplate/controllers"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Init() {
	serverConfig := config.GetConfig().Server
	log := logrus.New()
	log.SetReportCaller(true)

	router := gin.New()
	gin.SetMode(gin.DebugMode)

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)
	router.GET("/ping", health.Status)

	router.Run(serverConfig.Host + ":" + serverConfig.Port)
}
