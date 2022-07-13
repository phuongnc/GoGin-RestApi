package util

import (
	"errors"
	"shapeservice/configs"
	"shapeservice/internal/pkg/msg"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	SigningKey []byte
}

type CustomClaims struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type TokenInfo struct {
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiredAt    time.Time `json:"expired_at"`
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(configs.GetConfig().TokenSecret),
	}
}

func NewJWTSignedWithPasswordHash(hash string) *JWT {
	return &JWT{
		[]byte(hash),
	}
}

func (j *JWT) GenerateToken(id uint, email string) (TokenInfo, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Hour)

	claims := CustomClaims{
		id,
		email,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "Shapeservice",
		},
	}
	token, err := j.CreateToken(claims)
	refreshToken, _ := j.RefreshToken(token)
	return TokenInfo{
		Token:        token,
		RefreshToken: refreshToken,
		ExpiredAt:    expireTime,
	}, err
}

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New(msg.GetMsg(msg.ERROR_AUTH_CHECK_TOKEN_FAIL))
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New(msg.GetMsg(msg.ERROR_AUTH_CHECK_TOKEN_TIMEOUT))
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New(msg.GetMsg(msg.ERROR_AUTH_CHECK_TOKEN_FAIL))
			} else {
				return nil, errors.New(msg.GetMsg(msg.ERROR_AUTH_CHECK_TOKEN_FAIL))
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, errors.New(msg.GetMsg(msg.ERROR_AUTH_CHECK_TOKEN_FAIL))

	} else {
		return nil, errors.New(msg.GetMsg(msg.ERROR_AUTH_CHECK_TOKEN_FAIL))
	}
}

func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(6 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", errors.New(msg.GetMsg(msg.ERROR_AUTH_CHECK_TOKEN_FAIL))
}
