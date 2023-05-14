package scene

import (
	"github.com/acurry/goray/internal/light"
	"github.com/acurry/goray/internal/point"
	"github.com/acurry/goray/internal/sphere"
)

type screen struct {
	left   float64
	top    float64
	right  float64
	bottom float64
}

type Scene struct {
	width    int32
	height   int32
	maxDepth int32
	light    *light.Light
	shapes   []sphere.Sphere
	ratio    float64
	*screen
	camera *point.Point
}

func New() *Scene {
	return &Scene{
		width:    0,
		height:   0,
		maxDepth: 0,
		light:    light.New(),
		shapes:   make([]sphere.Sphere, 10),
		ratio:    0.0,
		screen: &screen{
			left:   0.0,
			top:    0.0,
			right:  0.0,
			bottom: 0.0,
		},
	}
}
