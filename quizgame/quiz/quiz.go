package quiz

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type problem struct {
	q string
	a string
}

func StartQuiz(file *os.File, duration time.Duration) error {
	r := csv.NewReader(file)

	records, err := r.ReadAll()

	if err != nil {
		return err
	}

	problems := parseProblems(records)

	checkProblems(duration, problems)

	return nil
}

func checkProblems(duration time.Duration, problems []problem) {

	timer := time.NewTimer(duration)
	var c int

problemsloop:
	for i, p := range problems {
		fmt.Printf("Problem %d - %s: ", i, p.q)

		answerChan := make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChan <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("\nYou have reached the time limit!")
			break problemsloop
		case ans := <-answerChan:
			if ans == p.a {
				c++
			}
		}
	}

	printSummary(len(problems), c)

}

func parseProblems(records [][]string) []problem {
	ret := make([]problem, len(records))
	for i, r := range records {
		ret[i] = problem{
			r[0],
			r[1],
		}
	}
	return ret
}

func printSummary(n, i int) {
	fmt.Println("-- Quiz Summary --")
	fmt.Printf("You scored %d out of %d\n", i, n)
}
