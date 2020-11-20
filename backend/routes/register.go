package routes

import (
	"projetoapi/controllers"

	"github.com/gin-gonic/gin"
)

// @Summary Adiciona um Worker
// @Description Adiciona um worker na base de dados
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Param worker body model.Worker true "Add worker"
// @Router /admin/users [post]
// @Success 201 {object} model.Worker
// @Failure 400 "Bad request"
// @Failure 404 "Not found"
func Register(c *gin.Context) {
	controllers.Register(c)
}