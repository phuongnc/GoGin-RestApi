package service

import (
	"errors"
	"shapeservice/cmd/entity-server/model"
	"shapeservice/cmd/entity-server/repository"
	"shapeservice/internal/pkg/msg"
	"shapeservice/internal/pkg/util"
	"time"

	log "shapeservice/internal/pkg/logger"
)

type LoginReq struct {
	Email    string `json:"email" valid:"email~Email is not valid"`
	Password string `json:"password" valid:"stringlength(6|50)~Password is at least 6 characters"`
}

type RegisterReq struct {
	Email     string `json:"email" valid:"email~Email is not valid"`
	Password  string `json:"password" valid:"stringlength(6|50)~Password is at least 6 characters"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type RefreshTokenReq struct {
	RefreshToken string `json:"refresh_token"`
}

type AccessTokenRes struct {
	AccessToken  string    `json:"token"`
	ExpiredAt    time.Time `json:"expired_at"`
	RefreshToken string    `json:"refresh_token"`
}

type UserRes struct {
	ID           uint      `json:"id"`
	Email        string    `json:"email"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiredAt    time.Time `json:"expired_at"`
}

func toUserRes(user *model.User) *UserRes {
	var userRes = &UserRes{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	return userRes
}

func (obj *LoginReq) Login() (*UserRes, error) {
	user, err := (&repository.AuthRepository{}).Login(obj.Email)
	if err != nil {
		return nil, err
	}
	if !util.CheckPasswordHash(obj.Password, user.Password) {
		return nil, errors.New("Login failed")
	}
	return toUserRes(user), nil
}

func (obj *RegisterReq) Register() error {
	passHash, err := util.HashPassword(obj.Password)
	if err != nil {
		return err
	}
	model := model.User{
		Email:     obj.Email,
		Password:  passHash,
		FirstName: obj.FirstName,
		LastName:  obj.LastName,
	}
	authRepository := repository.AuthRepository{}

	isExists, _ := authRepository.IsEmailExist(obj.Email)
	if isExists {
		return errors.New(msg.GetMsg(msg.ERROR_EXIST_ACCOUNT))
	}
	if err = authRepository.Register(&model); err != nil {
		log.GetLogger().Error(err)
		return errors.New(msg.GetMsg(msg.ERROR_ADD_ACCOUNT_FAIL))
	} else {
		return nil
	}
}
