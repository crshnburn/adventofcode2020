package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func CountQuestions(lines []string) int {
	questions := make(map[string]bool)
	for _, line := range lines {
		for _, question := range line {
			questions[string(question)] = true
		}
	}
	return len(questions)
}

func CountQuestions2(lines []string) int {
	questions := make(map[string]int)
	for _, line := range lines {
		for _, question := range line {
			if _, ok := questions[string(question)]; !ok {
				questions[string(question)] = 1
			} else {
				questions[string(question)] = questions[string(question)] + 1
			}
		}
	}
	allanswered := 0
	for _, answered := range questions {
		if answered == len(lines) {
			allanswered++
		}
	}
	return allanswered
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	currentPass := []string{}
	questionTotal := 0
	question2Total := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			questionTotal += CountQuestions(currentPass)
			question2Total += CountQuestions2(currentPass)
			currentPass = []string{}
		} else {
			currentPass = append(currentPass, line)
		}
	}
	questionTotal += CountQuestions(currentPass)
	question2Total += CountQuestions2(currentPass)
	fmt.Println("Total questions answered", questionTotal)
	fmt.Println("Total questions answered by everyone", question2Total)
}
