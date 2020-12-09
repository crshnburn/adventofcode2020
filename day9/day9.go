package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func CheckNumberValid(value int, input []int) bool {
	// fmt.Println(value, input)
	for i, val := range input {
		for j, val1 := range input {
			if i != j {
				if val+val1 == value {
					return true
				}
			}
		}
	}
	return false
}

func FindInvalidNumber(input []int, preamble int) int {
	for i, val := range input[preamble:] {
		if !CheckNumberValid(val, input[i:i+preamble]) {
			return val
		}
	}
	return -1
}

func FindInvalidTotal(input []int, sum int) int {
	total := 0
	for i, _ := range input {
		for j, _ := range input[i+1:] {
			total = Sum(input[i : i+j])
			if total > sum {
				break
			} else if total == sum {
				numbers := input[i : i+j]
				sort.Ints(numbers)
				return numbers[0] + numbers[len(numbers)-1]
			}
		}
	}
	return -1
}

func Sum(vals []int) int {
	sum := 0

	for _, val := range vals {
		sum += val
	}

	return sum
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []int
	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		input = append(input, val)
	}
	invalidNumber := FindInvalidNumber(input, 25)
	fmt.Println("Part 1:", invalidNumber)
	fmt.Println("Part 2:", FindInvalidTotal(input, invalidNumber))
}
