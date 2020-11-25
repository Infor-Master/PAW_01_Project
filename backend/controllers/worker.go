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

	services.Db.Where("id = ?", id).Find(&worker, id)

	if worker.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	}

	services.Db.Model(&worker).Association("zones").Clear()
	services.Db.Delete(&worker)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete succeeded!"})
}

func FindWorker(workers [] model.Worker, id uint) bool {
    for _, worker := range workers {
        if worker.ID == id {
            return true
        }
    }
    return false
}

func WorkerHasAccessToZone (c *gin.Context, zone *model.Zone) bool{
	var claims = services.GetClaims(c)

	if claims == nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Something went bad!"})
		return false 
	}

	var workers [] model.Worker

	services.Db.Model(&zone).Related(&workers,  "Workers")
	
	if FindWorker(workers, claims.Id) {
		return true
	} else {
		return false
	}
}