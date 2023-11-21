package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question/answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file: %s\n", *csvFilename))
	}

	r := csv.NewReader(file)
	record, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse provided CSV file")
	}
	problems := parseLines(record)

	var correct int
	for i, problem := range problems {
		fmt.Printf("Question #%d: %s =\n", i+1, problem.question)
		var answer string
		_, _ = fmt.Scanf("%s\n", &answer)
		if answer == problem.answer {
			fmt.Println("correct!")
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	rtn := make([]problem, len(lines))
	for i, line := range lines {
		rtn[i] = problem{
			question: line[0],
			answer:   line[1],
		}
	}
	return rtn
}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
