package main

import (
	"fmt"
	"log"
	"os"

	"github.com/acurry/goray/internal/point"
	"gopkg.in/yaml.v3"
)

const version = "1.0.0"

type Position struct {
	Pos point.Point `yaml:",flow"`
}

func main() {
	// Load the file; returns []byte
	f, err := os.ReadFile("test.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var position Position

	if err := yaml.Unmarshal(f, &position); err != nil {
		log.Fatal(err)
	}

	// Print out the new struct
	fmt.Printf("%+v\n", position)
}
