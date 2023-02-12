package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/EdwinVesga/gophercises/cyoa/story"
)

func main() {
	filename := flag.String("file", "gopher.json", "Json file of the CYOA story.")
	port := flag.Int("port", 3000, "Defines the http port of cyoa app.")

	f, err := os.Open(*filename)

	if err != nil {
		panic(err)
	}

	st, err := story.JsonStory(f)

	if err != nil {
		panic(err)
	}

	h := story.NewHandler(st)
	fmt.Printf("Starting server on port: %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
