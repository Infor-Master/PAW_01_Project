package controllers

import (
	"net/http"
	"projetoapi/model"
	"projetoapi/services"
	"strconv"

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

	//worker

	var claims = services.GetClaims(c)

	if claims == nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Something went bad!"})
		return
	}

	var worker model.Worker

	services.Db.Where("id = ?", claims.Id).First(&worker)

	if worker.ID != claims.Id {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Didn't find this worker!"})
		return
	}

	/*// zone

	var zone model.Zone

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid parameters!"})
		return
	}

	uintID := uint(id)

	services.Db.Where("id = ?", uintID).First(&zone)

	if zone.ID != uintID {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Didn't find this zone!"})
		return
	}

	var temp model.Worker

	services.Db.Where(&Worker{Id: worker.Id}, &Zone{Id:zone.Id}).First(&temp)

	//fmt.Print("vamos la - >     ")
	//fmt.Println(temp)

	//db.Model(&zone).Related(&worker,  "Users")

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": zone})*/

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

	services.Db.First(&zone, id)

	if zone.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	}

	services.Db.Delete(&zone)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete succeeded!"})
}

func AddPerson(c *gin.Context) {

	var zone model.Zone

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid parameters!"})
		return
	}

	uintID := uint(id)

	services.Db.Where("id = ?", uintID).First(&zone)

	if zone.ID != uintID {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Didn't find this zone!"})
		return
	}

	if zone.PplCount <= zone.Limits {
		zone.PplCount++

		services.Db.Save(&zone)

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Added Person"})
	}

	c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Limit exceeded!"})

}

func RemovePerson(c *gin.Context) {

	var zone model.Zone

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid parameters!"})
		return
	}

	uintID := uint(id)

	services.Db.Where("id = ?", uintID).First(&zone)

	if zone.ID != uintID {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Didn't find this zone!"})
		return
	}

	if zone.PplCount > 0 {
		zone.PplCount--
		services.Db.Save(&zone)

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Removed Person"})
	}

	c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Can't remove people if count is 0!"})

}
