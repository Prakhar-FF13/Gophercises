package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type Question struct {
	question, answer string
}

func main() {
	/*
		Read time of quiz from user
	*/
	t := flag.Int64("time", 30, "Quiz timmings (seconds)")

	flag.Parse()

	/*
		Open csv file.
	*/
	ior, err := os.Open("./problems.csv")
	if err != nil {
		log.Fatal("Problem opening problems csv\n")
	}
	defer ior.Close()

	/*
		Read csv file.
	*/
	r := csv.NewReader(ior)
	x, err := r.ReadAll()

	if err != nil {
		log.Fatal("Problem reading csv")
	}

	/*
		Convert read csv file into Question type array.
	*/
	q := []Question{}
	for i := 0; i < len(x); i++ {
		z := Question{x[i][0], x[i][1]}
		q = append(q, z)
	}

	/*
		Define a timer for the quiz.
	*/
	timer := time.NewTimer(time.Duration(*t) * time.Second)

	// store no. of correct responses.
	correct := 0

	for i, ques := range q {
		fmt.Printf("Question #%d: %s = ", i+1, ques.question)

		// Take input in a go routine so that it does not block the main thread.
		// If we dont do this then even if timer runs out, program wont close until user enters input.
		// used to get answer from go routine
		answerChan := make(chan string)
		go func() {
			var x string
			fmt.Scan(&x)
			answerChan <- x
		}()

		// select makes sure whatever matches it runs.
		select {
		// if timer passes print result.
		case <-timer.C:
			fmt.Printf("\nYou scored: %d\n", correct)
			return
		case answer := <-answerChan:
			if answer == ques.answer {
				correct++
			}
		}
	}
	fmt.Printf("\nYou scored: %d\n", correct)
}
