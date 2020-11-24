package controllers

/*func bool verifyIfWorkerHasZone(c *gin.Context){
	var zone = GetZone(c)
	var worker = GetWorker(c)
}*/


/*func Worker GetWorker(c *gin.Context){
	var worker model.Worker;

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid parameters!"})
		return
	}

	uintID := uint(id)

	services.Db.Where("id = ?", uintID).First(&worker)

	if worker.ID != uintID {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Didn't find this worker!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": worker})
	return worker

}*/
}*/
