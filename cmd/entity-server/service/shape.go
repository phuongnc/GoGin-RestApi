package service

import (
	"shapeservice/cmd/entity-server/model"
	"shapeservice/cmd/entity-server/repository"
	"shapeservice/internal/pkg/util"

	log "shapeservice/internal/pkg/logger"
)

//
type SquareReq struct {
	Length float64 `json:"length" valid:"required~Param is required, length~Param must be > 0"`
}

type SquareRes struct {
	ID     uint    `json:"id"`
	Length float64 `json:"length"`
}

//
type RectangleReq struct {
	Length float64 `json:"length" valid:"required~Param is required, length~Param must be > 0"`
	Width  float64 `json:"width" valid:"required~Param is required, length~Param must be > 0"`
}
type RectangleRes struct {
	ID     uint    `json:"id"`
	Length float64 `json:"length"`
	Width  float64 `json:"width"`
}

//
type TriangleReq struct {
	Length1 float64 `json:"length1" valid:"required~Param is required, length~Param must be > 0"`
	Length2 float64 `json:"length2" valid:"required~Param is required, length~Param must be > 0"`
	Length3 float64 `json:"length3" valid:"required~Param is required, length~Param must be > 0"`
}

type TriangleRes struct {
	ID      uint    `json:"id"`
	Length1 float64 `json:"length1"`
	Length2 float64 `json:"length2"`
	Length3 float64 `json:"length3"`
}

//
type DiamondReq struct {
	Diagonal1 float64 `json:"diagonal1" valid:"required~Param is required, length~Param must be > 0"`
	Diagonal2 float64 `json:"diagonal2" valid:"required~Param is required, length~Param must be > 0"`
}

type DiamondRes struct {
	ID        uint    `json:"id"`
	Diagonal1 float64 `json:"diagonal1"`
	Diagonal2 float64 `json:"diagonal2"`
}

//
// type ShapeCommonReq struct {
// 	ID         uint   `json:"id"`
// 	ShapeType  string `json:"shape_type"`
// 	PageLimit  uint   `json:"page_limit"`
// 	PageOffset uint   `json:"page_offset"`
// }

type Shape struct {
	model *model.Shape
}

type ShapeBuilder interface {
	WithID(ID uint) ShapeBuilder
	WithShapeType(st string) ShapeBuilder
	WithUserCreated(ID uint) ShapeBuilder
	WithUserUpdated(ID uint) ShapeBuilder
	WithSquare(s SquareReq) ShapeBuilder
	WithRectangle(r RectangleReq) ShapeBuilder
	WithTriangle(t TriangleReq) ShapeBuilder
	WithDiamond(d DiamondReq) ShapeBuilder
	Build() Shape
}

// type ShapeFunc interface {
// 	Add()
// 	Update()
// 	Get()
// 	GetAll()
// 	Delete()
// }

func (obj *Shape) WithID(ID uint) ShapeBuilder {
	obj.model.ID = ID
	return obj
}

func (obj *Shape) WithShapeType(st string) ShapeBuilder {
	obj.model.Type = st
	return obj
}

func (obj *Shape) WithUserCreated(ID uint) ShapeBuilder {
	obj.model.CreatedByUserID = ID
	return obj
}

func (obj *Shape) WithUserUpdated(ID uint) ShapeBuilder {
	obj.model.UpdatedByUserID = ID
	return obj
}

func (obj *Shape) WithSquare(s SquareReq) ShapeBuilder {
	obj.model.Type = util.SQUARE
	obj.model.Length1 = s.Length
	obj.model.Area = util.CalculateArea(util.SQUARE, s.Length, 0, 0)
	obj.model.Perimeter = util.CalculatePerimeter(util.SQUARE, s.Length, 0, 0)
	return obj
}
func (obj *Shape) WithRectangle(r RectangleReq) ShapeBuilder {
	obj.model.Type = util.RECTANGLE
	obj.model.Length1 = r.Width
	obj.model.Length2 = r.Length
	obj.model.Area = util.CalculateArea(util.RECTANGLE, r.Width, r.Length, 0)
	obj.model.Perimeter = util.CalculatePerimeter(util.RECTANGLE, r.Width, r.Length, 0)
	return obj
}
func (obj *Shape) WithTriangle(t TriangleReq) ShapeBuilder {
	obj.model.Type = util.TRIANGLE
	obj.model.Length1 = t.Length1
	obj.model.Length2 = t.Length2
	obj.model.Length3 = t.Length3
	obj.model.Area = util.CalculateArea(util.TRIANGLE, t.Length1, t.Length2, t.Length3)
	obj.model.Perimeter = util.CalculatePerimeter(util.RECTANGLE, t.Length1, t.Length2, t.Length3)
	return obj
}
func (obj *Shape) WithDiamond(d DiamondReq) ShapeBuilder {
	obj.model.Type = util.DIAMOND
	obj.model.Length1 = d.Diagonal1
	obj.model.Length2 = d.Diagonal2
	obj.model.Area = util.CalculateArea(util.DIAMOND, d.Diagonal1, d.Diagonal2, 0)
	obj.model.Perimeter = util.CalculatePerimeter(util.DIAMOND, d.Diagonal1, d.Diagonal2, 0)
	return obj
}

func (obj Shape) Build() Shape {
	return obj
}

func NewShapeBuilder() ShapeBuilder {
	return &Shape{model: new(model.Shape)}
}

func (obj *Shape) getRepository() *repository.ShapeRepository {
	return &repository.ShapeRepository{Model: obj.model}
}

func (obj *Shape) Gets(maps map[string]interface{}) (interface{}, error) {
	lstObject, err := obj.getRepository().Gets(maps)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, err
	}

	var res []interface{}
	for _, item := range lstObject {
		res = append(res, toResponse(item))
	}
	return res, nil
}

func (obj *Shape) Get() (interface{}, error) {
	resObj, err := obj.getRepository().Get()
	if err != nil {
		log.GetLogger().Error(err)
		return nil, err
	}
	return toResponse(resObj), nil
}

func (obj *Shape) GetShapeArea() (float64, error) {
	resObj, err := obj.getRepository().Get()
	if err != nil {
		log.GetLogger().Error(err)
		return 0, err
	}
	return resObj.Area, nil
}

func (obj *Shape) GetShapePerimeter() (float64, error) {
	resObj, err := obj.getRepository().Get()
	if err != nil {
		log.GetLogger().Error(err)
		return 0, err
	}
	return resObj.Perimeter, nil
}

func (obj *Shape) Add() (interface{}, error) {
	resObj, err := obj.getRepository().Add()
	if err != nil {
		log.GetLogger().Error(err)
		return nil, err
	}
	return toResponse(resObj), nil
}

func (obj *Shape) Update() (interface{}, error) {
	resObj, err := obj.getRepository().Update()
	if err != nil {
		log.GetLogger().Error(err)
		return nil, err
	}
	return toResponse(resObj), nil
}

func (obj *Shape) Delete() error {
	err := obj.getRepository().Delete()
	if err != nil {
		log.GetLogger().Error(err)
	}
	return err
}

func toResponse(model *model.Shape) interface{} {
	switch model.Type {
	case util.SQUARE:
		return SquareRes{ID: model.ID, Length: model.Length1}
	case util.RECTANGLE:
		return RectangleRes{ID: model.ID, Width: model.Length1, Length: model.Length2}
	case util.TRIANGLE:
		return TriangleRes{ID: model.ID, Length1: model.Length1, Length2: model.Length2, Length3: model.Length3}
	case util.DIAMOND:
		return DiamondRes{ID: model.ID, Diagonal1: model.Length1, Diagonal2: model.Length2}
	}
	return nil
}
