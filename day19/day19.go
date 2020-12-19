package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func Parse(input []string) map[string]string {
	rules := make(map[string]string)
	for _, rule := range input {
		parts := strings.Split(rule, ": ")
		rules[parts[0]] = parts[1]
	}
	return rules
}

func GetRule(rules map[string]string, number string) []string {
	var values []string
	instr := rules[number]
	if instr == "\"a\"" {
		values = append(values, "a")
	} else if instr == "\"b\"" {
		values = append(values, "b")
	} else {
		var pipeParts [][]string
		parts := strings.Split(instr, " | ")
		for _, part := range parts {
			var subvals [][]string
			numbers := strings.Split(part, " ")
			for _, num := range numbers {
				subvals = append(subvals, GetRule(rules, num))
			}
			pipeParts = append(pipeParts, CombineRules(subvals))
		}
		for _, part := range pipeParts {
			for _, val := range part {
				values = append(values, val)
			}
		}
	}
	return values
}

func CombineRules(rules [][]string) []string {
	if len(rules) == 1 {
		return rules[0]
	} else {
		toReturn := []string{}
		subRules := CombineRules(rules[1:])
		for _, part1 := range rules[0] {
			for _, part2 := range subRules {
				toReturn = append(toReturn, part1+part2)
			}
		}
		return toReturn
	}
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	rawRules := []string{}
	for len(line) > 0 {
		rawRules = append(rawRules, line)
		scanner.Scan()
		line = scanner.Text()
	}
	data := []string{}
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	count := 0
	rules := Parse(rawRules)
	rule0 := GetRule(rules, "0")
	for _, line := range data {
		for _, rule := range rule0 {
			if line == rule {
				count++
				break
			}
		}
	}
	fmt.Println("Part 1:", count)

	//Rule 0: 8 11
	//Rule 8: 42 | 42 8
	//Rule 11: 42 31 | 42 11 31
	//Expanded: 42+ 42+ 31+

	regex42 := strings.Join(GetRule(rules, "42"), "|")
	regex31 := strings.Join(GetRule(rules, "31"), "|")
	// pattern := "(" + regex42 + ")+("
	// for i := 1; i < 10; i++ {
	// 	pattern += fmt.Sprintf("(%s){%d}(%s){%d}}|", regex42, i, regex31, i)
	// }
	// pattern += "(" + regex42 + "({10}(" + regex31 + "){10}})"
	pattern := "^(" + regex42 + ")+((" + regex42 + ")(" + regex31 + ")|(" + regex42 + "){2}(" + regex31 + "){2}|(" + regex42 + "){3}(" + regex31 + "){3}|(" + regex42 + "){4}(" + regex31 + "){4})$"
	regex, _ := regexp.Compile(pattern)
	count1 := 0
	for _, line := range data {
		if regex.MatchString(line) {
			count1++
		}
	}
	fmt.Println("Part 2:", count1)
}
