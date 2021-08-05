package main

import (
	"bufio"
	"fmt"
	"os"
)

type Problem struct {
	Question      string
	CorrectAnswer string
}

type Answer struct {
	Answer  string
	Correct bool
}

func main() {
	problems := getProblems()
	var answers []*Answer

	scanner := bufio.NewScanner(os.Stdin)

	for i, problem := range problems {
		fmt.Print(fmt.Sprintf("Problem #%d: %s = ", i+1, problem.Question))
		scanner.Scan()

		answer := scanner.Text()

		answers = append(answers, &Answer{
			Answer:  answer,
			Correct: answer == problem.CorrectAnswer,
		})
	}

	var score uint
	for _, a := range answers {
		if a.Correct {
			score += 1
		}
	}
	fmt.Println(fmt.Sprintf("You scored %d out of %d.", score, len(problems)))
}

func getProblems() []*Problem {
	return []*Problem{
		{Question: "1+1", CorrectAnswer: "2"},
		{Question: "2+1", CorrectAnswer: "3"},
	}
}
