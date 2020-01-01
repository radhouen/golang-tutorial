package advanced_maths

import "math"

func Tan(angle float64) float64 {
	tan := math.Cos(angle) / math.Sin(angle)
	return tan
}

func Cotan(angle float64) float64  {
cotan := 1/ Tan(angle)
return cotan
}

func Expo(angle float64) float64  {
	return  1
}
