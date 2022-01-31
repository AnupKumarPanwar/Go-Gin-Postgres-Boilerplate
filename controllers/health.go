package controllers

import (
	"Go-Gin-Postgres-Boilerplate/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type HealthController struct {
}

func (h HealthController) Status(c *gin.Context) {

	db, err := database.GetDB().DB()

	if err != nil {
		logrus.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = db.Ping()

	if err != nil {
		logrus.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, "pong")
}
