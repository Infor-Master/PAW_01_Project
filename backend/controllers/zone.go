package controllers

import (
	"fmt"
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

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid parameters!"})
		return
	}

	uintID := uint(id)

	services.Db.First(&worker, "id = ?", uintID)
	services.Db.Model(&worker).Related(&zones, "zones")

	//db.First(&language, "id = ?", 111)
	//db.Model(&language).Related(&users,  "Users")
	// SELECT * FROM "users" INNER JOIN "user_languages" ON
	//"user_languages"."user_id" = "users"."id"
	//WHERE  ("user_languages"."language_id" IN ('111'))

	if worker.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Didn't find this zone!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": zones})
}

func GetZone(c *gin.Context) {

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

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": zone})
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
	fmt.Print(id)
	services.Db.First(&zone, id)

	if zone.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	}

	services.Db.Delete(&zone)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete succeeded!"})
}

// nao esta feita ainda
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

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": zone})

}
