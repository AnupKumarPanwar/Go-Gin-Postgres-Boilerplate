package controllers

import (
	"github.com/anupkumarpanwar/Go-Gin-Postgres-Boilerplate/database"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type HealthController struct {
}

func (h HealthController) Status(c *gin.Context) {
	db := database.GetDB()

	err := db.DB().Ping()

	if err != nil {
		logrus.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, "pong")
}
