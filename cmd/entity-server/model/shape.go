package model

import "gorm.io/gorm"

/*
Define for shape:
Square: Length
Rectangle: Length, Length1
Triangle : Length, Length1, Length2
Diamond: Length, Length1
*/

type Shape struct {
	gorm.Model
	Type            string
	Length1         float64
	Length2         float64
	Length3         float64
	Area            float64
	Perimeter       float64
	CreatedByUserID uint
	CreatedByUser   User
	UpdatedByUserID uint
	UpdatedByUser   User
}
