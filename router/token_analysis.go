package router

import (
	"dy/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func TokenAnalysisRoute(c *gin.Context) {
	tokenString := c.Query("token")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetJWTconfig()), nil
	}, jwt.WithJSONNumber())

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "unknown forbidden",
		})
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["UserId"], claims["exp"])
		c.JSON(http.StatusOK, gin.H{
			"userId": claims["UserId"],
			"exp":    claims["exp"],
		})
	} else {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "token analysisFail",
		})
	}
}
