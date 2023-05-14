package scene

import (
	"math"

	"github.com/acurry/goray/internal/light"
	"github.com/acurry/goray/internal/point"
	"github.com/acurry/goray/internal/sphere"
)

type screen struct {
	Left   float64
	Top    float64
	Right  float64
	Bottom float64
}

type Scene struct {
	Width    int32
	Height   int32
	MaxDepth int32 `yaml:"max_depth"`
	Light    light.Light
	Shapes   []sphere.Sphere
	Ratio    float64
	*screen  `yaml:",omitempty"`
	Camera   point.Point `yaml:",flow"`
}

func (s *Scene) Init() {
	s.Ratio = float64(s.Width) / float64(s.Height)
	s.Left = -1.0
	s.Top = 1.0 / s.Ratio
	s.Right = 1.0
	s.Bottom = -1.0 / s.Ratio
}

func (s *Scene) NearestIntersect(
	origin *point.Point,
	direction *point.Point,
) (*sphere.Sphere, float64) {
	distances := make([]float64, len(s.Shapes))
	for _, shape := range s.Shapes {
		distances = append(distances, shape.CheckIntersect(origin, direction))
	}

	var nearest *sphere.Sphere
	minDistance := math.Inf(1)
	for i, distance := range distances {
		if distance > 0 && distance < minDistance {
			minDistance = distance
			nearest = &s.Shapes[i]
		}
	}

	return nearest, minDistance
}
