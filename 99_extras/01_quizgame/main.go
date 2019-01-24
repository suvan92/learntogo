package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// QuizOptions hold all the flag customization options
type QuizOptions struct {
	// Name of the file containing the questions
	File string
	// Duration of the quiz in seconds
	Duration int
}

// Question is a struct that associates a question
// and an answer
type Question struct {
	// The question string
	Q string `json:"q"`
	// The answer string
	A string `json:"a"`
}

// parseFlags
func parseFlags() QuizOptions {
	// define flags
	qFile := flag.String("f", "problems.csv", "name of the csv file containing questions and answers")
	duration := flag.Int("d", 30, "duration of the quiz in seconds")

	// parse flags
	flag.Parse()

	// convert to options
	return QuizOptions{
		File:     *qFile,
		Duration: *duration,
	}
}

func readFile(filename string) (questions []Question) {
	// open file
	csvFile, err := os.Open(filename)
	defer csvFile.Close()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// create reader
	r := csv.NewReader(csvFile)

	// parse lines
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		questions = append(questions, Question{
			Q: line[0],
			A: line[1],
		})
	}

	return questions
}

func ask(q string) (a string) {
	// Present question
	fmt.Print(q, "? ")
	// Read input
	fmt.Scanln(&a)
	return a
}

func main() {
	opts := parseFlags()
	questions := readFile(opts.File)

	welcomeMessage := "\nWelcome to the quiz! You will have %v seconds to answer %v questions.\nPress any key to begin...\n"

	fmt.Printf(welcomeMessage, opts.Duration, len(questions))
	fmt.Scanln()

	var score int
	var answered int

	timer := time.NewTimer(time.Duration(opts.Duration) * time.Second)
	go func() {
		<-timer.C
		fmt.Printf("\nTimes up! You answered %v of %v questions and got %v correct\n", answered, len(questions), score)
		os.Exit(0)
	}()

	for _, q := range questions {
		a := ask(q.Q)

		if a == q.A {
			score++
		}

		answered++
	}

	fmt.Println("You got", score, "out of", len(questions), "correct")
}
