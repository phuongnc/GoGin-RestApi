package router

import (
	"shapeservice/cmd/entity-server/api"
	"shapeservice/internal/pkg/ginutil"

	"github.com/gin-gonic/gin"
)

func initShapeRouter(Router *gin.RouterGroup) {
	shapeRouter := Router.Group("shapes").Use(ginutil.JWTAuth())
	{
		shapeRouter.POST("squares", api.AddSquare)
		shapeRouter.PUT("squares/:id", api.UpdateSquare)

		shapeRouter.POST("rectangles", api.AddRectangle)
		shapeRouter.PUT("rectangles/:id", api.UpdateRectangle)

		shapeRouter.POST("triangles", api.AddTriangle)
		shapeRouter.PUT("triangles/:id", api.UpdateTriangle)

		shapeRouter.POST("diamonds", api.AddDiamond)
		shapeRouter.PUT("diamonds/:id", api.UpdateDiamond)

		shapeRouter.GET("", api.GetShapes)
		shapeRouter.GET(":id", api.GetShape)
		shapeRouter.GET(":id/area", api.GetShapeArea)
		shapeRouter.GET(":id/perimeter", api.GetShapePerimeter)
		shapeRouter.DELETE(":id", api.DeleteShape)
	}
}
