package main

import (
	"fmt"
	"log"
	"os"

	"github.com/acurry/goray/internal/scene"
	"gopkg.in/yaml.v3"
)

const version = "1.0.0"

type Root struct {
	Scene scene.Scene
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

	// Print out the new struct
	fmt.Printf("%+v\n", root)
}
