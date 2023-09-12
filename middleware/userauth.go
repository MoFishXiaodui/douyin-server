package middleware

import (
	"dy/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

func UserAuth(c *gin.Context) {
	// 参考 https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-Parse-Hmac

	// Get the token
	tokenString := c.Query("token")
	if tokenString == "" {
		//	如果在Query参数找不到，就试着在 application/form-data 找一下
		tokenString = c.Request.FormValue("token")
	}

	// 解码验证
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetJWTconfig()), nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"status_code": -1,
			"status_msg":  "token解析出错",
		})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		// 确认有无过期
		exp := claims["exp"].(float64)
		if float64(time.Now().Unix()) > exp {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
				"status_code": -1,
				"status_msg":  "token expired",
			})
			return
		}
		// 暂存userId
		c.Set("UserId", claims["UserId"].(float64))

		// 继续
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"status_code": -1,
			"status_msg":  "token失效",
		})
		return
	}
}
