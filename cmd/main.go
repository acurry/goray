package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/acurry/goray/internal/point"
	"github.com/acurry/goray/internal/scene"
	"github.com/acurry/goray/internal/utils"
	"gopkg.in/yaml.v3"
)

const version = "1.0.0"

type Root struct {
	Scene scene.Scene
}

func save(filePath string, img *image.NRGBA) {
	imgFile, err := os.Create(filePath)
	defer imgFile.Close()
	if err != nil {
		log.Println("Cannot create file:", err)
	}
	png.Encode(imgFile, img.SubImage(img.Rect))
}

func main() {
	// Load the file; returns []byte
	f, err := os.ReadFile("scene1.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var root Root

	if err := yaml.Unmarshal(f, &root); err != nil {
		log.Fatal(err)
	}

	root.Scene.Init()

	rect := image.Rect(0, 0, int(root.Scene.Width), int(root.Scene.Height))

	pix := make([]uint8, rect.Dx()*rect.Dy()*4)

	outer := utils.LinSpace(root.Scene.Screen.Top, root.Scene.Screen.Bottom, root.Scene.Height)
	inner := utils.LinSpace(root.Scene.Screen.Left, root.Scene.Screen.Right, root.Scene.Width)

	for i, x := range outer {
		for j, y := range inner {
			pixel := point.Point{X: x, Y: y, Z: 0.0}
			direction := pixel.Normalize()
		}
	}

	// Print out the new struct
	fmt.Printf("%+v\n", root)
}
