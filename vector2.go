package vector2

import (
	"fmt"
	"math"
)

type Vector2 struct {
	X float64
	Y float64
}

func Add(a, b Vector2) Vector2 {
	return Vector2{a.X + b.X, a.Y + b.Y}
}

func Sub(a, b Vector2) Vector2 {
	return Vector2{a.X - b.X, a.Y - b.Y}
}

func Length(p Vector2) float64 {
	return math.Sqrt(math.Pow(p.X, 2) + math.Pow(p.Y, 2))
}

func ScalarMul(p Vector2, a float64) Vector2 {
	return Vector2{p.X * a, p.Y * a}
}

func VectorMul(a Vector2, b Vector2) float64 {
	return a.X*b.X + a.Y*b.Y
}

func GetAngleBetween2Vectors(a Vector2, b Vector2) float64 {
	return math.Acos(VectorMul(a, b)/(Length(a)*Length(b))) * (180.0 / math.Pi)
}

func RotateVectorByAngle(p Vector2, angle float64, isDegrees bool) Vector2 {
	if isDegrees {
		angle = angle * (math.Pi / 180)
	}
	return Vector2{p.X*math.Cos(angle) - p.Y*math.Sin(angle), p.X*math.Sin(angle) + p.Y*math.Cos(angle)}
}

func PrintV(p Vector2) {
	fmt.Printf("[%v, %v]", p.X, p.Y)
}
