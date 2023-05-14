package scene

import (
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
