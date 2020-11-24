package services

import (

	"net/http"
	"projetoapi/model"
	"strconv"

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

func Find(workers [] model.Worker, id uint) bool {
    for _, worker := range workers {
        if worker.ID == id {
            return true
        }
    }
    return false
}

func WorkerHasAccessToZone (c *gin.Context, zone *model.Zone) bool{
	var claims = GetClaims(c)

	if claims == nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Something went bad!"})
		return false 
	}

	var workers [] model.Worker

	Db.Model(&zone).Related(&workers,  "Workers")
	
	if Find(workers, claims.Id) {
		return true
	} else {
		return false
	}
}

func FindZone(c *gin.Context) *model.Zone{
	var zone model.Zone

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid parameters!"})
		return nil
	}

	uintID := uint(id)

	Db.Where("id = ?", uintID).First(&zone)

	if zone.ID != uintID {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Didn't find this zone!"})
		return nil
	}
	
	return &zone
}