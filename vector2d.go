package main

import "math"

type Vector2D struct {
	x float64
	y float64
}

func (v *Vector2D) Addition(vector Vector2D) Vector2D {
	return Vector2D{x: v.x + vector.x, y: v.y + vector.y}
}

func (v *Vector2D) Subtraction(vector Vector2D) Vector2D {
	return Vector2D{x: v.x - vector.x, y: v.y - vector.y}
}

func (v *Vector2D) Multiply(vector Vector2D) Vector2D {
	return Vector2D{x: v.x * vector.x, y: v.y * vector.y}
}

func (v *Vector2D) AddV(d float64) Vector2D {
	return Vector2D{x: v.x + d, y: v.y + d}
}

func (v *Vector2D) SubtractionV(d float64) Vector2D {
	return Vector2D{x: v.x - d, y: v.y - d}
}

func (v *Vector2D) MultiplyV(d float64) Vector2D {
	return Vector2D{x: v.x * d, y: v.y * d}
}

func (v *Vector2D) DivisionV(d float64) Vector2D {
	return Vector2D{x: v.x / d, y: v.y / d}
}

func (v *Vector2D) limit(lower, upper float64) Vector2D {
	return Vector2D{x: math.Min(math.Max(v.x, lower), upper),
		y: math.Min(math.Max(v.y, lower), upper),
	}
}

func (v *Vector2D) Distance(v2 Vector2D) float64 {
	return math.Sqrt(math.Pow(v.x-v2.x, 2) + math.Pow(v.y-v2.y, 2))
}
