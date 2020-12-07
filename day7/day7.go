package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Bag struct {
	Name     string
	Contains map[*Bag]int
}

func ParseRules(rules []string) map[string]*Bag {
	bags := make(map[string]*Bag)
	for _, rule := range rules {
		parts := strings.Split(rule, " contain ")
		bagName := parts[0]
		currentBag, ok := bags[bagName]
		if !ok {
			currentBag = &Bag{Name: bagName, Contains: make(map[*Bag]int)}
			bags[bagName] = currentBag
		}
		for _, containsRule := range strings.Split(parts[1], ", ") {
			regex, _ := regexp.Compile(`(\d)\s(.*)`)
			values := regex.FindStringSubmatch(containsRule)
			if len(values) == 3 {
				number, _ := strconv.Atoi(values[1])
				subBagName := values[2]

				if strings.HasSuffix(subBagName, ".") {
					subBagName = string(subBagName[0 : len(subBagName)-1])
				}
				if !(strings.HasSuffix(subBagName, "s")) {
					subBagName += "s"
				}
				subBag, found := bags[subBagName]
				if !found {
					subBag = &Bag{Name: subBagName, Contains: make(map[*Bag]int)}
					bags[subBagName] = subBag
				}
				currentBag.Contains[subBag] = number
			}
		}
	}
	return bags
}

func FindBag(b Bag, target string) bool {
	for c := range b.Contains {
		if c.Name == target || FindBag(*c, target) {
			return true
		}
	}
	return false
}

func CountBag(b Bag) int {
	fmt.Println(b)
	count := 0
	for c, number := range b.Contains {
		count += number + (number * CountBag(*c))
	}
	return count
}

func main() {
	file, err := os.Open(("./input.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rules := []string{}
	for scanner.Scan() {
		rules = append(rules, scanner.Text())
	}
	bags := ParseRules(rules)
	count := 0
	for _, bag := range bags {
		if FindBag(*bag, "shiny gold bags") {
			count++
		}
	}

	fmt.Println("There are", count, "bags")
	fmt.Println("Shiny gold bags contain", CountBag(*bags["shiny gold bags"]), "bags")
}
