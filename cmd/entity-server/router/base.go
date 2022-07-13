package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitRouter(env string) *gin.Engine {
	var router = gin.Default()
	apiGroup := router.Group(fmt.Sprintf("%s/api", env))
	initAuthRouter(apiGroup)
	initShapeRouter(apiGroup)
	return router
}
