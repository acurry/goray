package sphere

import (
	"math"

	"github.com/acurry/goray/internal/point"
)

type Sphere struct {
	Radius     float64
	Shininess  int64
	Reflection float64
	Center     point.Point `yaml:",flow"`
	Ambient    point.Point `yaml:",flow"`
	Diffuse    point.Point `yaml:",flow"`
	Specular   point.Point `yaml:",flow"`
}

func (s *Sphere) CheckIntersect(origin *point.Point, direction *point.Point) float64 {
	b := 2.0 * direction.Dot(origin.Sub(&s.Center))
	x := origin.Sub(&s.Center)
	c := x.Normal() - (s.Radius * s.Radius)

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
