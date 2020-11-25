package routes

import (
	"projetoapi/controllers"

	"github.com/gin-gonic/gin"
)

// @Summary Recupera os users
// @Description Exibe a lista de workers e de zonas
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Success 200 {array} model.Worker
// @Router /admin/ [get]
// @Router /admin/association [get]
// @Failure 404 "Not found"
func AssociateUsersZones(c *gin.Context) {
	controllers.AssociateUsersZones(c)
}


func DesassociateUsersZones(c *gin.Context) {
	controllers.DesassociateUsersZones(c)
}