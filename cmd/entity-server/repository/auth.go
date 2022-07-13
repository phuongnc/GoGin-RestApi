package repository

import (
	"shapeservice/cmd/entity-server/model"

	log "shapeservice/internal/pkg/logger"

	"gorm.io/gorm"
)

type AuthRepository struct {
}

func (auth *AuthRepository) Login(email string) (*model.User, error) {
	var u model.User
	if err := db.Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (auth *AuthRepository) Register(u *model.User) error {
	return db.Create(u).Error
}

func (obj *AuthRepository) IsEmailExist(email string) (bool, error) {
	var u model.User
	err := db.Where("email = ?", email).Unscoped().First(&u).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.GetLogger().Error(err)
		return false, err
	}
	if u.ID > 0 {
		return true, nil
	}
	return false, nil
}
