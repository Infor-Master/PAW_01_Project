package routes

import (
	"projetoapi/controllers"

	"github.com/gin-gonic/gin"
)

// @Summary Associa workers a zonas
// @Description Associa workers a zonas via IDs
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Success 200 {string} resourceId
// @Router /api/admin/association [post]
// @Failure 400,404 "Not Found"
func AssociateUsersZones(c *gin.Context) {
	controllers.AssociateUsersZones(c)
}

// @Summary Desassocia workers a zonas
// @Description Desassocia workers a zonas vis IDs
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Success 200 {array} model.Worker
// @Router /admin/ [get]
// @Router /admin/association [get]
// @Failure 404 "Not found"
func DesassociateUsersZones(c *gin.Context) {
	controllers.DesassociateUsersZones(c)
}
