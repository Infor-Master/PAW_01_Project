package controllers

/*
func GetWorkerZones(c *gin.Context) {

	var worker model.Worker
	var zone model.Zone

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid parameters!"})
		return
	}

	uintID := uint(id)

	services.Db.Where("id = ?", uintID).First(&worker)

	if worker.ID != uintID {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Didn't find this zone!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": worker.ZoneID})
}
*/
