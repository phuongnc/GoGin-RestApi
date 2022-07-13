package router

import (
	"shapeservice/cmd/entity-server/api"

	"github.com/gin-gonic/gin"
)

func initAuthRouter(Router *gin.RouterGroup) {
	authenRouter := Router.Group("auth")
	{
		authenRouter.POST("login", api.Login)
		authenRouter.POST("register", api.Register)
		authenRouter.POST("refresh-token", api.RefreshToken)
	}
}
