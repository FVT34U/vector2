package vector2

import (
	"fmt"
	"math"
)

type vector2 struct {
	x float64
	y float64
}

func Add(a, b vector2) vector2 {
	return vector2{a.x + b.x, a.y + b.y}
}

func Sub(a, b vector2) vector2 {
	return vector2{a.x - b.x, a.y - b.y}
}

func Length(p vector2) float64 {
	return math.Sqrt(math.Pow(p.x, 2) + math.Pow(p.y, 2))
}

func ScalarMul(p vector2, a float64) vector2 {
	return vector2{p.x * a, p.y * a}
}

func VectorMul(a vector2, b vector2) float64 {
	return a.x*b.x + a.y*b.y
}

func GetAngleBetween2Vectors(a vector2, b vector2) float64 {
	return math.Acos(VectorMul(a, b)/(Length(a)*Length(b))) * (180.0 / math.Pi)
}

func RotateVectorByAngle(p vector2, angle float64, isDegrees bool) vector2 {
	if isDegrees {
		angle = angle * (math.Pi / 180)
	}
	return vector2{p.x*math.Cos(angle) - p.y*math.Sin(angle), p.x*math.Sin(angle) + p.y*math.Cos(angle)}
}

func PrintV(p vector2) {
	fmt.Printf("[%v, %v]", p.x, p.y)
}
