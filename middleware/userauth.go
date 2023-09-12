package middleware

import (
	"dy/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strconv"
	"time"
)

func UserAuth(c *gin.Context) {
	// 参考 https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-Parse-Hmac

	// Get the token
	tokenString := c.Query("token")

	// 解码验证
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetJWTconfig()), nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"msg": "token解析出错",
		})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		// 确认有无过期
		exp := claims["exp"].(float64)
		if float64(time.Now().Unix()) > exp {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
				"msg": "token expired",
			})
			return
		}

		// 确认用户是否正确
		fmt.Println(claims["UserId"])
		// 先把query参数转成整数
		userId, errUserId := strconv.Atoi(c.Query("user_id"))
		if errUserId != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"msg": "unexpected user_id",
			})
		}
		//fmt.Println(reflect.TypeOf(claims["UserId"]))	// float64
		//fmt.Println(reflect.TypeOf(userId))	// int

		if float64(userId) == claims["UserId"].(float64) {
			// 继续
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"msg": "token 校验失败",
				// 冒用他人userID
			})
			return
		}
	} else {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"msg": "token失效",
		})
		return
	}
}
