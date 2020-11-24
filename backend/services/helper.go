package services

import (

	"projetoapi/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)


func GetClaims(c *gin.Context) *model.Claims {

	var tokenInput, _, _ = GetAuthorizationToken(c)
	token, _ := jwt.ParseWithClaims(tokenInput, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
		//fmt.Printf("%v %v", claims.Username, claims.StandardClaims.ExpiresAt)
		//c.Set("username", claims.Username)
		//fmt.Print("claims: ")
		//fmt.Println(claims)
		return claims
	}
	return nil
}