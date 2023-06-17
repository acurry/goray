package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/acurry/goray/internal/scene"
	"gopkg.in/yaml.v3"
)

type Root struct {
	Scene scene.Scene
}

func Save(filePath string, img *image.NRGBA) {
	imgFile, err := os.Create(filePath)
	if err != nil {
		log.Println("Cannot create file:", err)
	}
	defer imgFile.Close()
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

	// Print out the new struct
	fmt.Printf("%+v\n", root)
}
