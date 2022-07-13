package ginutil

import (
	"net/http"
	"shapeservice/internal/pkg/msg"
	jwt "shapeservice/internal/pkg/util"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			Response(c, http.StatusUnauthorized, false, msg.GetMsg(msg.ERROR_AUTH_TOKEN), gin.H{
				"reload": true,
			}, nil)
			c.Abort()
			return
		}
		j := jwt.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			Response(c, http.StatusUnauthorized, false, err.Error(), gin.H{
				"reload": true,
			}, nil)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
