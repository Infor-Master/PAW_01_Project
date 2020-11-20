package routes

import (
	"projetoapi/controllers"

	"github.com/gin-gonic/gin"
)


// @Summary Adicionar pessoa a zona
// @Description Adiciona uma pessoa a um local específico
// @Accept  json
// @Produce  json
// @Router /zones/:id/add [post]
// @Param evaluation body model.Worker true "Do register"
// @Success 200 {object} model.Claims
// @Failure 400 "Bad request"
// @Failure 401 "Unauthorized"
func addUser(c *gin.Context) {
	controllers.addPerson(c)
}