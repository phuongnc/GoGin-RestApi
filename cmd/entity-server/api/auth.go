package api

import (
	"net/http"
	"shapeservice/cmd/entity-server/service"
	"shapeservice/internal/pkg/msg"
	"shapeservice/internal/pkg/util"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	appG := Gin{C: c}

	var objService service.LoginReq
	isValid := appG.BindAndValidate(&objService)

	if isValid {
		user, err := objService.Login()
		if err != nil {
			appG.Response(http.StatusUnauthorized, false, msg.GetMsg(msg.ERROR_AUTH), nil, nil)
			return
		}

		j := util.NewJWT()
		tokenInfo, err := j.GenerateToken(user.ID, user.Email)
		if err != nil {
			appG.Response(http.StatusBadRequest, false, err.Error(), nil, nil)
			return
		}
		user.AccessToken = tokenInfo.Token
		user.RefreshToken = tokenInfo.RefreshToken
		user.ExpiredAt = tokenInfo.ExpiredAt
		appG.Response(http.StatusOK, true, msg.GetMsg(msg.SUCCESS), user, nil)
	}
}

func RefreshToken(c *gin.Context) {
	appG := Gin{C: c}

	var objService service.RefreshTokenReq
	isValid := appG.BindAndValidate(&objService)
	if isValid {
		j := util.NewJWT()
		user, err := j.ParseToken(objService.RefreshToken)
		if err != nil {
			appG.Response(http.StatusUnauthorized, false, msg.GetMsg(msg.ERROR_AUTH_TOKEN), nil, nil)
			return
		}
		tokenInfo, err := j.GenerateToken(user.Id, user.Email)
		if err != nil {
			appG.Response(http.StatusBadRequest, false, err.Error(), nil, nil)
			return
		}
		accessToken := service.AccessTokenRes{
			AccessToken:  tokenInfo.Token,
			ExpiredAt:    tokenInfo.ExpiredAt,
			RefreshToken: tokenInfo.RefreshToken,
		}
		appG.Response(http.StatusOK, true, msg.GetMsg(msg.SUCCESS), accessToken, nil)
	}
}

func Register(c *gin.Context) {
	appG := Gin{C: c}

	var objService service.RegisterReq
	isValid := appG.BindAndValidate(&objService)

	if isValid {
		err := objService.Register()
		if err != nil {
			appG.Response(http.StatusBadRequest, false, err.Error(), nil, nil)
			return
		}
		appG.Response(http.StatusOK, true, msg.GetMsg(msg.SUCCESS), nil, nil)
	}
}
