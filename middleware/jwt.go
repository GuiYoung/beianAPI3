package middlewares

import (
	"beianAPI/utils/errmsg"
	"beianAPI/utils/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// check token string
func JWY() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := errmsg.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = errmsg.ERROR_TOKEN_WRONG
		} else {
			claims, err := jwt.ParseToken(token)
			if err != nil {
				code = errmsg.ERROR_TOKEN_WRONG
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = errmsg.ERROR_TOKEN_EXPIRED
			}
		}

		if code != errmsg.SUCCESS {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			return
		}
		c.Next()
	}
}
