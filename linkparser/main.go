package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/EdwinVesga/gophercises/linkparser/parser"
)

func main() {
	paths, err := filepath.Glob("./example/*.html")
	if err != nil {
		panic(err)
	}

	for i, path := range paths {
		file, _ := os.Open(path)
		if err != nil {
			panic(err)
		}

		links, err := parser.Parse(file)

		if err != nil {
			panic(err)
		}

		fmt.Printf("Exercise %d\n", i+1)
		fmt.Printf("%+v\n", links)
		fmt.Println("-----------------------------")
	}
}
