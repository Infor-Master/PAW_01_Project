package controllers

import (
	"net/http"
	"projetoapi/model"
	"projetoapi/services"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var workers []model.Worker

	services.Db.Find(&workers)

	if len(workers) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Didn't find any zone!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": workers})
}

func DeleteUser(c *gin.Context) {
	var worker model.Worker

	id := c.Param("id")

	services.Db.First(&worker, id)

	if worker.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	}

	services.Db.Model(&worker).Association("zones").Delete(&worker)
	services.Db.Delete(&worker)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete succeeded!"})
}
