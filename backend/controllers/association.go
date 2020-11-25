package controllers

import (
	"net/http"
	"projetoapi/model"
	"projetoapi/services"

	"github.com/gin-gonic/gin"
)

func AssociateUsersZones(c *gin.Context) {

	type JDATA struct {
		IDWORKER int `json:"idWorker" binding:"required"`
		IDZONE int `json:"idZone" binding:"required"`
	}

	var json JDATA

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Check syntax!"})
		return
	}

	var zone model.Zone
	var worker model.Worker

	services.Db.Where("id = ?", json.IDWORKER).Find(&worker, json.IDWORKER)
	services.Db.Where("id = ?", json.IDZONE).Find(&zone, json.IDZONE)

	services.Db.Model(&worker).Association("zones").Append(&zone)
	
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Create successful!","resourceId": json})
}

func DesassociateUsersZones(c *gin.Context) {

	type JDATA struct {
		IDWORKER int `json:"idWorker" binding:"required"`
		IDZONE int `json:"idZone" binding:"required"`
	}

	var json JDATA

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Check syntax!"})
		return
	}

	var zone model.Zone
	var worker model.Worker

	services.Db.Where("id = ?", json.IDWORKER).Find(&worker, json.IDWORKER)
	services.Db.Where("id = ?", json.IDZONE).Find(&zone, json.IDZONE)

	services.Db.Model(&worker).Association("zones").Delete(&zone)
	
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Delete successful!","resourceId": json})
}