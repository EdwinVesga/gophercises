package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/EdwinVesga/gophercises/quizgame/quiz"
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

	d := time.Second * time.Duration(*quizTimeLimit)

	err = quiz.StartQuiz(f, d)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
