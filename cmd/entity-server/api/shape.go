package api

import (
	"net/http"
	"shapeservice/cmd/entity-server/service"
	"shapeservice/internal/pkg/ginutil"
	"shapeservice/internal/pkg/msg"
	"shapeservice/internal/pkg/util"
	"strconv"

	"shapeservice/configs"

	"github.com/gin-gonic/gin"
)

func AddSquare(c *gin.Context) {
	appG := Gin{C: c}

	claims, _ := c.Get("claims")
	user := claims.(*util.CustomClaims)

	var reqObj service.SquareReq
	isValid := appG.BindAndValidate(&reqObj)
	if isValid {
		shapeBuilder := service.NewShapeBuilder().
			WithSquare(reqObj).
			WithUserCreated(user.Id).
			Build()
		objRes, err := shapeBuilder.Add()
		if err != nil {
			appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.ERROR_ADD_FAIL), nil, nil)
			return
		}
		appG.Response(http.StatusOK, true, msg.GetMsg(msg.SUCCESS), objRes, nil)
	}
}

func UpdateSquare(c *gin.Context) {
	appG := Gin{C: c}

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.INVALID_PARAMS), nil, nil)
		return
	}

	claims, _ := c.Get("claims")
	user := claims.(*util.CustomClaims)

	var reqObj service.SquareReq
	isValid := appG.BindAndValidate(&reqObj)
	if isValid {
		shapeBuilder := service.NewShapeBuilder().
			WithID(uint(ID)).
			WithSquare(reqObj).
			WithUserUpdated(user.Id).
			Build()
		objRes, err := shapeBuilder.Update()
		if err != nil {
			appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.ERROR_UPDATE_FAIL), nil, nil)
			return
		}
		appG.Response(http.StatusOK, true, msg.GetMsg(msg.SUCCESS), objRes, nil)
	}
}

func AddRectangle(c *gin.Context) {
	appG := Gin{C: c}

	claims, _ := c.Get("claims")
	user := claims.(*util.CustomClaims)

	var reqObj service.RectangleReq
	isValid := appG.BindAndValidate(&reqObj)
	if isValid {
		shapeBuilder := service.NewShapeBuilder().
			WithRectangle(reqObj).
			WithUserCreated(user.Id).
			Build()
		objRes, err := shapeBuilder.Add()
		if err != nil {
			appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.ERROR_ADD_FAIL), nil, nil)
			return
		}
		appG.Response(http.StatusOK, true, msg.GetMsg(msg.SUCCESS), objRes, nil)
	}
}

func UpdateRectangle(c *gin.Context) {
	appG := Gin{C: c}

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.INVALID_PARAMS), nil, nil)
		return
	}

	claims, _ := c.Get("claims")
	user := claims.(*util.CustomClaims)

	var reqObj service.RectangleReq
	isValid := appG.BindAndValidate(&reqObj)
	if isValid {
		shapeBuilder := service.NewShapeBuilder().
			WithID(uint(ID)).
			WithRectangle(reqObj).
			WithUserUpdated(user.Id).
			Build()
		objRes, err := shapeBuilder.Update()
		if err != nil {
			appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.ERROR_UPDATE_FAIL), nil, nil)
			return
		}
		appG.Response(http.StatusOK, true, msg.GetMsg(msg.SUCCESS), objRes, nil)
	}
}

func AddTriangle(c *gin.Context) {
	appG := Gin{C: c}

	claims, _ := c.Get("claims")
	user := claims.(*util.CustomClaims)

	var reqObj service.TriangleReq
	isValid := appG.BindAndValidate(&reqObj)
	if isValid {
		shapeBuilder := service.NewShapeBuilder().
			WithTriangle(reqObj).
			WithUserCreated(user.Id).
			Build()
		objRes, err := shapeBuilder.Add()
		if err != nil {
			appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.ERROR_ADD_FAIL), nil, nil)
			return
		}
		appG.Response(http.StatusOK, true, msg.GetMsg(msg.SUCCESS), objRes, nil)
	}
}

func UpdateTriangle(c *gin.Context) {
	appG := Gin{C: c}

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.INVALID_PARAMS), nil, nil)
		return
	}

	claims, _ := c.Get("claims")
	user := claims.(*util.CustomClaims)

	var reqObj service.TriangleReq
	isValid := appG.BindAndValidate(&reqObj)
	if isValid {
		shapeBuilder := service.NewShapeBuilder().
			WithID(uint(ID)).
			WithTriangle(reqObj).
			WithUserUpdated(user.Id).
			Build()
		objRes, err := shapeBuilder.Update()
		if err != nil {
			appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.ERROR_UPDATE_FAIL), nil, nil)
			return
		}
		appG.Response(http.StatusOK, true, msg.GetMsg(msg.SUCCESS), objRes, nil)
	}
}

func AddDiamond(c *gin.Context) {
	appG := Gin{C: c}

	claims, _ := c.Get("claims")
	user := claims.(*util.CustomClaims)

	var reqObj service.DiamondReq
	isValid := appG.BindAndValidate(&reqObj)
	if isValid {
		shapeBuilder := service.NewShapeBuilder().
			WithDiamond(reqObj).
			WithUserCreated(user.Id).
			Build()
		objRes, err := shapeBuilder.Add()
		if err != nil {
			appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.ERROR_ADD_FAIL), nil, nil)
			return
		}
		appG.Response(http.StatusOK, true, msg.GetMsg(msg.SUCCESS), objRes, nil)
	}
}

func UpdateDiamond(c *gin.Context) {
	appG := Gin{C: c}

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.INVALID_PARAMS), nil, nil)
		return
	}

	claims, _ := c.Get("claims")
	user := claims.(*util.CustomClaims)

	var reqObj service.DiamondReq
	isValid := appG.BindAndValidate(&reqObj)
	if isValid {
		shapeBuilder := service.NewShapeBuilder().
			WithID(uint(ID)).
			WithDiamond(reqObj).
			WithUserUpdated(user.Id).
			Build()
		objRes, err := shapeBuilder.Update()
		if err != nil {
			appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.ERROR_UPDATE_FAIL), nil, nil)
			return
		}
		appG.Response(http.StatusOK, true, msg.GetMsg(msg.SUCCESS), objRes, nil)
	}
}

func GetShapes(c *gin.Context) {
	appG := Gin{C: c}

	shapeType := c.Query("type")
	if shapeType == "" {
		appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.INVALID_PARAMS), nil, nil)
		return
	}
	offset, limit := ginutil.GetPage(c, configs.GetConfig().DefaultPageNum, configs.GetConfig().DefaultPageLimit)

	maps := map[string]interface{}{
		"page_limit":  limit,
		"page_offset": offset,
		"order_by":    c.Query("order_by"),
		"order":       c.Query("order"),
	}

	shapeBuilder := service.NewShapeBuilder().
		WithShapeType(shapeType).
		Build()

	objRes, err := shapeBuilder.Gets(maps)
	if err != nil {
		appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.ERROR_GET_FAIL), nil, nil)
		return
	}
	appG.Response(http.StatusOK, true, msg.GetMsg(msg.SUCCESS), objRes, nil)

}

func GetShape(c *gin.Context) {
	appG := Gin{C: c}

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.INVALID_PARAMS), nil, nil)
		return
	}

	shapeBuilder := service.NewShapeBuilder().
		WithID(uint(ID)).
		Build()
	objRes, err := shapeBuilder.Get()
	if err != nil {
		appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.ERROR_GET_FAIL), nil, nil)
		return
	}
	appG.Response(http.StatusOK, true, msg.GetMsg(msg.SUCCESS), objRes, nil)

}

func GetShapeArea(c *gin.Context) {
	appG := Gin{C: c}

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.INVALID_PARAMS), nil, nil)
		return
	}

	shapeBuilder := service.NewShapeBuilder().
		WithID(uint(ID)).
		Build()
	objRes, err := shapeBuilder.GetShapeArea()
	if err != nil {
		appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.ERROR_GET_FAIL), nil, nil)
		return
	}
	appG.Response(http.StatusOK, true, msg.GetMsg(msg.SUCCESS), objRes, nil)
}

func GetShapePerimeter(c *gin.Context) {
	appG := Gin{C: c}

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.INVALID_PARAMS), nil, nil)
		return
	}

	shapeBuilder := service.NewShapeBuilder().
		WithID(uint(ID)).
		Build()
	objRes, err := shapeBuilder.GetShapePerimeter()
	if err != nil {
		appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.ERROR_GET_FAIL), nil, nil)
		return
	}
	appG.Response(http.StatusOK, true, msg.GetMsg(msg.SUCCESS), objRes, nil)
}

func DeleteShape(c *gin.Context) {
	appG := Gin{C: c}

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.INVALID_PARAMS), nil, nil)
		return
	}

	shapeBuilder := service.NewShapeBuilder().
		WithID(uint(ID)).
		Build()

	err = shapeBuilder.Delete()
	if err != nil {
		appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.ERROR_DELETE_FAIL), nil, nil)
		return
	}
	appG.Response(http.StatusOK, true, msg.GetMsg(msg.SUCCESS), nil, nil)

}
