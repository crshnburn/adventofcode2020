package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func ScanLine(line string) (int, int, string, string) {
	var min int
	var max int
	var character string
	var password string
	fmt.Sscanf(line, "%d-%d %1s: %s", &min, &max, &character, &password)
	return min, max, character, password
}

func ValidatePassword(line string) bool {
	min, max, character, password := ScanLine(line)
	chars := strings.Split(password, "")
	sort.Strings(chars)
	sortedPassword := strings.Join(chars, "")
	if strings.LastIndex(sortedPassword, character) > -1 {
		charLength := (strings.LastIndex(sortedPassword, character) + 1) - strings.Index(sortedPassword, character)
		return charLength >= min && charLength <= max
	} else {
		return false
	}
}

func ValidatePassword2(line string) bool {
	first, second, char, password := ScanLine(line)
	return (string(password[first-1]) == char) != (string(password[second-1]) == char)
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	validPasswords := 0
	validPasswords2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		if ValidatePassword(line) {
			validPasswords += 1
		}
		if ValidatePassword2(line) {
			validPasswords2 += 1
		}
	}

	fmt.Println("ValidPasswords", validPasswords)
	fmt.Println("ValidPasswords2", validPasswords2)
}
