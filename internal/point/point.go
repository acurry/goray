package point

import (
	"math"
)

type Point struct {
	X float64
	Y float64
	Z float64
}

func (a *Point) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var pointDetails []float64
	if err := unmarshal(&pointDetails); err != nil {
		return err
	}

	a.X = pointDetails[0]
	a.Y = pointDetails[1]
	a.Z = pointDetails[2]

	return nil
}

func Dot(a *Point, b *Point) float64 {
	return (a.X * b.X) + (a.Y * b.Y) + (a.Z * b.Z)
}

func Cross(a *Point, b *Point) *Point {
	return &Point{
		(a.Y * b.Z) - (a.Z * b.Y),
		(a.Z * b.X) - (a.X * b.Z),
		(a.X * b.Y) - (a.Y * b.X),
	}
}

func (a *Point) Dot(b *Point) float64 {
	return (a.X * b.X) + (a.Y * b.Y) + (a.Z * b.Z)
}

func (a *Point) Cross(b *Point) *Point {
	return &Point{
		(a.Y * b.Z) - (a.Z * b.Y),
		(a.Z * b.X) - (a.X * b.Z),
		(a.X * b.Y) - (a.Y * b.X),
	}
}

func (a *Point) Normal() float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z)
}

func (a *Point) Normalize() *Point {
	norm := a.Normal()
	return &Point{
		a.X / norm,
		a.Y / norm,
		a.Z / norm,
	}
}

func Reflected(vector *Point, axis *Point) *Point {
	/* return vector - 2 * np.dot(vector, axis) * axis */
	x := Mul(axis, 2*Dot(vector, axis))
	return Sub(vector, x)
}

func Add(a *Point, b *Point) *Point {
	return &Point{
		a.X + b.X,
		a.Y + b.Y,
		a.Z + b.Z,
	}
}

func (a *Point) Add(b *Point) *Point {
	return &Point{
		a.X + b.X,
		a.Y + b.Y,
		a.Z + b.Z,
	}
}

func Sub(a *Point, b *Point) *Point {
	return &Point{
		a.X - b.X,
		a.Y - b.Y,
		a.Z - b.Z,
	}
}

func (a *Point) Sub(b *Point) *Point {
	return &Point{
		a.X - b.X,
		a.Y - b.Y,
		a.Z - b.Z,
	}
}

func Mul(a *Point, x float64) *Point {
	return &Point{
		a.X * x,
		a.Y * x,
		a.Z * x,
	}
}

func (a *Point) Mul(x float64) *Point {
	return &Point{
		a.X * x,
		a.Y * x,
		a.Z * x,
	}
}

func Div(a *Point, x float64) *Point {
	return &Point{
		a.X / x,
		a.Y / x,
		a.Z / x,
	}
}

func (a *Point) Div(x float64) *Point {
	return &Point{
		a.X / x,
		a.Y / x,
		a.Z / x,
	}
}

func Clip(point *Point, min, max float64) *Point {
	var x float64
	var y float64
	var z float64

	if point.X >= max {
		x = max
	} else if point.X >= min && point.X < max {
		x = point.X
	} else {
		x = min
	}

	if point.Y >= max {
		y = max
	} else if point.Y >= min && point.Y < max {
		y = point.Y
	} else {
		y = min
	}

	if point.Z >= max {
		z = max
	} else if point.Z >= min && point.Z < max {
		z = point.Z
	} else {
		z = min
	}

	return &Point{
		x, y, z,
	}
}
