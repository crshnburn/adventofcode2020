package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func FindExpenses(expenses []int) int {
	return FindExpensesSum(expenses, 2020)
}

func FindExpensesSum(expenses []int, total int) int {
	result := -1
	sort.Ints(expenses)
	for _, expense := range expenses {
		searchFor := total - expense
		index := sort.SearchInts(expenses, searchFor)
		if index < len(expenses) && expenses[index] == searchFor {
			result = expense * searchFor
			break
		}
	}
	return result
}

func FindThreeExpenses(expenses []int) int {
	result := -1
	sort.Ints(expenses)
	for _, expense := range expenses {
		searchFor := 2020 - expense
		pairResult := FindExpensesSum(expenses, searchFor)
		if pairResult != -1 {
			result = pairResult * expense
			break
		}
	}
	return result
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var expenses []int
	for scanner.Scan() {
		expense, _ := strconv.Atoi(scanner.Text())
		expenses = append(expenses, expense)
	}

	fmt.Println("Tell the elves", FindExpenses(expenses))
	fmt.Println("Tell the elves (3)", FindThreeExpenses(expenses))
}
