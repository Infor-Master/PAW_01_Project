package routes

import (
	"projetoapi/controllers"

	"github.com/gin-gonic/gin"
)

// @Summary Recupera as zonas
// @Description Exibe a lista de todas as zonas
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Success 200 {array} model.Zone
// @Router /api/zones/all [get]
// @Failure 404 "Not found"
func GetZones(c *gin.Context) {
	controllers.GetZones(c)
}

// @Summary Recupera as zonas de um worker
// @Description Exibe a lista de todas as zonas de um worker
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Success 200 {array} model.Zone
// @Router /api/zones/worker [get]
// @Failure 404 "Not found"
func GetWorkerZones(c *gin.Context) {
	controllers.GetWorkerZones(c)
}

// @Summary Adicionar uma zona
// @Description Adiciona uma zona
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Param zone body model.Zone true "Add zone"
// @Router /api/admin/zones [post]
// @Success 201 {object} model.Zone
// @Failure 400 "Bad request"
// @Failure 404 "Not found"
func AddZone(c *gin.Context) {
	controllers.AddZone(c)
}

// @Summary Exclui uma zona pelo ID
// @Description Exclui uma zona realizada
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Param id path int true "Zone ID"
// @Router /api/admin/zones/{id} [delete]
// @Success 200 {object} model.Zone
// @Failure 404 "Not found"
func DeleteZone(c *gin.Context) {
	controllers.DeleteZone(c)
}

// @Summary Recupera uma zona
// @Description Exibe a lista de todas as zonas
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Param id path int true "Zone ID"
// @Success 200 {object} model.Zone
// @Router /api/zones/id/{id} [get]
// @Failure 404 "Not found"
func GetZone(c *gin.Context) {
	controllers.GetZone(c)
}

// @Summary Adicionar um worker
// @Description Adiciona um worker à zona
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Param worker body model.Worker true "Add worker"
// @Router /api/zones/id/{id}/add [post]
// @Success 201 {object} model.Worker
// @Failure 400 "Bad request"
// @Failure 404 "Not found"
func AddPerson(c *gin.Context) {
	controllers.AddPerson(c)
}

// @Summary Remove um worker
// @Description Remove um worker da zona
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Param id path int true "Worker ID"
// @Router /api/zones/id/{id}/remove [post]
// @Success 201 {object} model.Worker
// @Failure 400 "Bad request"
// @Failure 404 "Not found"
func RemovePerson(c *gin.Context) {
	controllers.RemovePerson(c)
}
