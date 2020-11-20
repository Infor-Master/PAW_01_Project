package controllers

import (
	"net/http"
	"projetoapi/model"
	"projetoapi/services"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var worker model.Worker

	if err := c.ShouldBindJSON(&worker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Check syntax!"})
		return
	}
	services.Db.Save(&worker)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Create successful!", "resourceId": worker.Name})
}