package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

const defaultFileName string = "problems.csv"
const defaultTimeLimit int = 30

func main() {
	fmt.Println("Quiz Game - Gophercise #1")

	csvFileName := flag.String("file", defaultFileName, "Csv file name with 'question,response' format.")
	quizTimeLimit := flag.Int("limit", defaultTimeLimit, "the time limit for the quiz in seconds.")

	flag.Parse()

	f, err := os.Open(*csvFileName)

	if err != nil {
		fmt.Printf("Failed to open CSV file: %s.", *csvFileName)
	}

	r := csv.NewReader(f)

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		for value := range record {
			fmt.Printf("%s\n", record[value])
		}
	}

}
