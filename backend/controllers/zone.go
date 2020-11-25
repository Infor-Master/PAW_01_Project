package controllers

import (
	"net/http"
	"projetoapi/model"
	"projetoapi/services"

	"github.com/gin-gonic/gin"
)

func GetZones(c *gin.Context) {
	var zones []model.Zone
	services.Db.Find(&zones)

	if len(zones) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Didn't find any zone!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": zones})
}

func GetWorkerZones(c *gin.Context) {

	var worker model.Worker
	var zones []model.Zone

	var claims = services.GetClaims(c)

	if claims == nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Something went bad!"})
		return
	}

	services.Db.First(&worker, "id = ?", claims.Id)
	services.Db.Model(&worker).Related(&zones, "zones")

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": zones})
}

func GetZone(c *gin.Context) { 

	zone := services.FindZone(c)

	if services.WorkerHasAccessToZone(c, zone) {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": zone})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Access Denied!"})
	}

}

func AddZone(c *gin.Context) {
	var zone model.Zone

	if err := c.ShouldBindJSON(&zone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Check syntax!"})
		return
	}
	services.Db.Save(&zone)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Create successful!", "resourceId": zone.Name})
}

func DeleteZone(c *gin.Context) {
	var zone model.Zone

	id := c.Param("id")

	services.Db.Where("id = ?", id).Find(&zone, id)

	if zone.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	}

	services.Db.Model(&zone).Association("workers").Clear()
	services.Db.Delete(&zone)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete succeeded!"})
}

func AddPerson(c *gin.Context) {

	zone := services.FindZone(c)

	if zone.PplCount <= zone.Limits {
		zone.PplCount++

		services.Db.Save(&zone)

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Added Person"})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Limit exceeded!"})

}

func RemovePerson(c *gin.Context) {

	zone := services.FindZone(c)

	if zone.PplCount > 0 {
		zone.PplCount--

		services.Db.Save(&zone)

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Removed Person"})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Can't remove people if count is 0!"})

}
