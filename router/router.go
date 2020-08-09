package router

import (
	"github.com/anupkumarpanwar/Go-Gin-Postgres-Boilerplate/config"
	"github.com/anupkumarpanwar/Go-Gin-Postgres-Boilerplate/controllers"
	"github.com/gin-gonic/gin"
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
