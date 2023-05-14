package sphere

import (
	"math"

	"github.com/acurry/goray/internal/point"
)

type Sphere struct {
	radius     float64
	shininess  int64
	reflection float64
	center     *point.Point
	ambient    *point.Point
	diffuse    *point.Point
	specular   *point.Point
}

func New() *Sphere {
	return &Sphere{
		radius:     0.0,
		shininess:  0,
		reflection: 0.0,
		center:     point.New(),
		ambient:    point.New(),
		diffuse:    point.New(),
		specular:   point.New(),
	}
}

func (s *Sphere) CheckIntersect(origin *point.Point, direction *point.Point) float64 {
	b := 2.0 * direction.Dot(origin.Sub(s.center))
	x := origin.Sub(s.center)
	c := x.Normal() - (s.radius * s.radius)

	delta := (b * b) - (4.0 * c)
	if delta > 0.0 {
		t1 := (-b + math.Sqrt(delta)) / 2
		t2 := (-b - math.Sqrt(delta)) / 2

		if t1 < t2 {
			return t1
		} else if t1 > t2 {
			return t2
		}
	}

	return -1.0
}
