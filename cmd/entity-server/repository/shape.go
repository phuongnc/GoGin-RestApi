package repository

import (
	"errors"
	"fmt"
	"shapeservice/cmd/entity-server/model"
	"shapeservice/internal/pkg/msg"

	"gorm.io/gorm"
)

type ShapeRepository struct {
	Model     *model.Shape
	pageNum   int
	pageLimit int
	order     map[string]interface{}
}

func (obj *ShapeRepository) Gets(maps map[string]interface{}) ([]*model.Shape, error) {
	var list []*model.Shape

	order := ""
	if maps["order"] != "" && maps["order_by"] != "" {
		order = fmt.Sprintf("%s %s", maps["order_by"], maps["order"])
	}

	err := db.
		Where("type = ?", obj.Model.Type).
		Order(order).
		Offset(maps["page_offset"].(int)).
		Limit(maps["page_limit"].(int)).
		Find(&list).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return list, nil
}

func (obj *ShapeRepository) Get() (*model.Shape, error) {
	err := db.
		Where("id = ?", obj.Model.ID).
		First(obj.Model).Error
	if err != nil {
		return nil, err
	}
	return obj.Model, nil
}

func (obj *ShapeRepository) Add() (*model.Shape, error) {
	if err := db.Create(obj.Model).Error; err != nil {
		return nil, err
	}
	return obj.Model, nil
}

func (obj *ShapeRepository) Update() (*model.Shape, error) {
	tmpObj := new(model.Shape)
	if err := db.Where("id = ?", obj.Model.ID).Find(tmpObj).Error; err != nil {

		return nil, err
	}
	db.Model(tmpObj).Updates(obj.Model)
	return obj.Model, nil
}

func (obj *ShapeRepository) Delete() error {
	resObj, err := obj.Get()
	if err != nil {
		return err
	}
	if resObj.ID > 0 {
		if err := db.Where("id = ?", resObj.ID).Delete(&resObj).Error; err != nil {
			return err
		}
		return nil
	}
	return errors.New(msg.GetMsg(msg.ERROR_NOT_EXIST))
}
