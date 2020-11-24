package routes

import (
	"projetoapi/controllers"

	"github.com/gin-gonic/gin"
)

// @Summary Recupera as zonas
// @Description Exibe a lista, sem todos os campos, de todas as zonas
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Success 200 {array} model.Zone
// @Router /zones [get]
// @Router /admin/zones [get]
// @Failure 404 "Not found"
func GetZones(c *gin.Context) {
	controllers.GetZones(c)
}

// @Summary Adicionar uma zona
// @Description Cria uma avaliação sobre a utilização da aplicação
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Param zone body model.Zone true "Add zone"
// @Router /admin/zones [post]
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
// @Router /admin/zones/{id} [delete]
// @Success 200 {object} model.Zone
// @Failure 404 "Not found"
func DeleteZone(c *gin.Context) {
	controllers.DeleteZone(c)
}



func GetZone(c *gin.Context) {
	controllers.GetZone(c)
}

func AddPerson(c *gin.Context) {
	controllers.AddPerson(c)
}

func RemovePerson(c *gin.Context) {
	controllers.RemovePerson(c)
}