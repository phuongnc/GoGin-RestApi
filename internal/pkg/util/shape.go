package util

import "math"

// shape mapping
const (
	SQUARE    string = "square"
	RECTANGLE        = "rectangle"
	TRIANGLE         = "triangle"
	DIAMOND          = "diamond"
)

func CalculateArea(t string, l1 float64, l2 float64, l3 float64) float64 {
	var area float64
	switch t {
	case SQUARE:
		area = math.Pow(l1, 2)
	case RECTANGLE:
		area = l1 * l2
	case TRIANGLE:
		//Heron method
		halfPer := (l1 + l2 + l3) / 2
		area = math.Sqrt(halfPer * (halfPer - l1) * (halfPer - l2) * (halfPer - l3))
	case DIAMOND:
		area = l1 * l1 / 2
	}
	return area
}

func CalculatePerimeter(t string, l1 float64, l2 float64, l3 float64) float64 {
	var perimeter float64
	switch t {
	case SQUARE:
		perimeter = l1 * 4
	case RECTANGLE:
		perimeter = (l1 + l2) * 2
	case TRIANGLE:
		perimeter = l1 + l2 + l3
	case DIAMOND:
		perimeter = 2 * math.Sqrt(math.Pow(l1, 2)+math.Pow(l2, 2))
	}
	return perimeter
}
