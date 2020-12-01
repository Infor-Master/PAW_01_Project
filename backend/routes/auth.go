package routes

import (
	"projetoapi/controllers"

	"github.com/gin-gonic/gin"
)

// @Summary Realizar autenticação
// @Description Autentica o utilizador e gera o token para os próximos acessos
// @Accept  json
// @Produce  json
// @Router /api/login [post]
// @Param json body string true "Do login"
// @Success 200 {object} model.Claims
// @Failure 400 "Bad request"
// @Failure 401 "Unauthorized"
func GenerateToken(c *gin.Context) {
	controllers.LoginHandler(c)
}

// @Summary Atualiza token de autenticação
// @Description Atualiza o token de autenticação do usuário
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Router /api/refresh_token [put]
// @param Authorization header string true "Token"
// @Success 200 {object} model.Claims
// @Failure 400 "Bad request"
// @Failure 401 "Unauthorized"
func RefreshToken(c *gin.Context) {
	controllers.RefreshHandler(c)
}

// @Summary Adiciona um Worker
// @Description Adiciona um worker na base de dados
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Param worker body model.Worker true "Add worker"
// @Router /api/admin/users [post]
// @Success 201 {object} model.Worker
// @Failure 400 "Bad request"
// @Failure 404 "Not found"
func Register(c *gin.Context) {
	controllers.Register(c)
}
