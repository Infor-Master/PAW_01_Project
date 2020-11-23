package routes

import (
	"projetoapi/controllers"

	"github.com/gin-gonic/gin"
)

func GetWorkerZoneByID(c *gin.Context) {
	controllers.GetWorkerZoneByID(c)
}
