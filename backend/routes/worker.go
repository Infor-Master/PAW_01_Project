package routes

import (
	"projetoapi/controllers"

	"github.com/gin-gonic/gin"
)

// @Summary Recupera os users
// @Description Exibe a lista, sem todos os campos, de todas as zonas
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Success 200 {array} model.Worker
// @Router /api/admin/users [get]
// @Failure 404 "Not found"
func GetUsers(c *gin.Context) {
	controllers.GetUsers(c)
}

// @Summary Exclui uma worker pelo ID
// @Description Exclui uma zona realizada
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Param id path int true "Worker ID"
// @Router /api/admin/users [delete]
// @Success 200 {object} model.Worker
// @Failure 404 "Not found"
func DeleteUser(c *gin.Context) {
	controllers.DeleteUser(c)
}
