package controllers

import (
	"net/http"
	"projetoapi/model"
	"projetoapi/services"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var creds model.Worker
	var usr model.Worker

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}
	services.OpenDatabase()

	services.Db.Find(&usr, "username = ?", creds.Username)
	if usr.Username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid User!"})
		return
	}

	if(!services.ComparePasswords(usr.Password,[]byte(creds.Password))){
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid User!"})
		return
	}

	token := services.GenerateTokenJWT(usr)

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Access denied!"})
			return
		}
		defer services.Db.Close()
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Success!", "token": token})

}

func RefreshHandler(c *gin.Context) {

	user := model.Worker{
		Username: c.GetHeader("username"),
	}

	token := services.GenerateTokenJWT(user)

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Acesso não autorizado"})
		return
	}

	defer services.Db.Close()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusNoContent, "message": "Token atualizado com sucesso!", "token": token})
}

func Register(c *gin.Context) {
	var worker model.Worker

	if err := c.ShouldBindJSON(&worker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Check syntax!"})
		return
	}

	hash := services.HashAndSalt([]byte(worker.Password))

	worker.Password = hash

	services.Db.Save(&worker)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Create successful!", "resourceId": worker.Name})
}
