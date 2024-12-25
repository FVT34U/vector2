package vector2

import (
	"fmt"
	"math"
)

type Vector2 struct {
	x float64
	y float64
}

func Add(a, b Vector2) Vector2 {
	return Vector2{a.x + b.x, a.y + b.y}
}

func Sub(a, b Vector2) Vector2 {
	return Vector2{a.x - b.x, a.y - b.y}
}

func Length(p Vector2) float64 {
	return math.Sqrt(math.Pow(p.x, 2) + math.Pow(p.y, 2))
}

func ScalarMul(p Vector2, a float64) Vector2 {
	return Vector2{p.x * a, p.y * a}
}

func VectorMul(a Vector2, b Vector2) float64 {
	return a.x*b.x + a.y*b.y
}

func GetAngleBetween2Vectors(a Vector2, b Vector2) float64 {
	return math.Acos(VectorMul(a, b)/(Length(a)*Length(b))) * (180.0 / math.Pi)
}

func RotateVectorByAngle(p Vector2, angle float64, isDegrees bool) Vector2 {
	if isDegrees {
		angle = angle * (math.Pi / 180)
	}
	return Vector2{p.x*math.Cos(angle) - p.y*math.Sin(angle), p.x*math.Sin(angle) + p.y*math.Cos(angle)}
}

func PrintV(p Vector2) {
	fmt.Printf("[%v, %v]", p.x, p.y)
}
