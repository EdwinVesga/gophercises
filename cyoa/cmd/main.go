package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/EdwinVesga/gophercises/cyoa/story"
)

func main() {
	filename := flag.String("file", "gopher.json", "Json file of the CYOA story.")

	f, err := os.Open(*filename)

	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(f)

	var story story.Story

	err = decoder.Decode(&story)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", story)
}
