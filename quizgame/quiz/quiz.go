package quiz

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func StartQuiz(file *os.File, duration time.Duration) {
	r := csv.NewReader(file)

	var i, sCounter int

	go quizDuration(duration, &i, &sCounter)

	for i = 1; ; i++ {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		question, response := record[0], record[1]

		fmt.Printf("Question %d - %s: ", i, question)

		scanner := bufio.NewScanner(os.Stdin)

		scanner.Scan()
		input := scanner.Text()

		if input == response {
			fmt.Println("Correct!")
			sCounter++
		} else {
			fmt.Println("Incorrect!")
		}
	}
}

func quizDuration(d time.Duration, n, i *int) {
	<-time.NewTimer(d).C
	printSummary(*n, *i)
	os.Exit(0)
}

func printSummary(n, i int) {
	fmt.Println()
	fmt.Println("You have reached the time limit!")
	fmt.Println("-- Quiz Summary --")
	fmt.Printf("Correct: %d, Incorrect: %d", i, n-i)
}
