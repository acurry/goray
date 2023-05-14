package internal

import "math"

type Point struct {
	x float64
	y float64
	z float64
}

func (a *Point) Dot(b *Point) float64 {
	return (a.x * b.x) + (a.y * b.y) + (a.z * b.z)
}

func (a *Point) Cross(b *Point) Point {
	return Point{
		(a.y * b.z) - (a.z * b.y),
		(a.z * b.x) - (a.x * b.z),
		(a.x * b.y) - (a.y * b.x),
	}
}

func (a *Point) Normal() float64 {
	return math.Sqrt(a.x*a.x + a.y*a.y + a.z*a.z)
}

func (a *Point) Normalize() Point {
	norm := a.Normal()
	return Point{
		a.x / norm,
		a.y / norm,
		a.z / norm,
	}
}

func (a *Point) Add(b *Point) Point {
	return Point{
		a.x + b.x,
		a.y + b.y,
		a.z + b.z,
	}
}

func (a *Point) Sub(b *Point) Point {
	return Point{
		a.x - b.x,
		a.y - b.y,
		a.z - b.z,
	}
}

func (a *Point) Div(x float64) Point {
	return Point{
		a.x / x,
		a.y / x,
		a.z / x,
	}
}
