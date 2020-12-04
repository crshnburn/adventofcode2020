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

func ReadPassports(lines []string) []map[string]string {
	var passportRecords []map[string]string
	currentPassport := make(map[string]string)
	for _, line := range lines {
		if len(line) == 0 {
			passportRecords = append(passportRecords, currentPassport)
			currentPassport = make(map[string]string)
		} else {
			fields := strings.Split(line, " ")
			for _, field := range fields {
				parts := strings.Split(field, ":")
				currentPassport[parts[0]] = parts[1]
			}
		}
	}
	passportRecords = append(passportRecords, currentPassport)
	return passportRecords
}

func ValidatePassport(passport map[string]string) bool {
	if len(passport) == 8 {
		return true
	} else if len(passport) == 7 {
		for key, _ := range passport {
			if key == "cid" {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func ValidatePassportFields(passport map[string]string) bool {
	for key, value := range passport {
		switch key {
		case "byr":
			year, err := strconv.Atoi(value)
			if err != nil {
				return false
			} else if year < 1920 || year > 2002 {
				return false
			}
			break
		case "iyr":
			year, err := strconv.Atoi(value)
			if err != nil {
				return false
			} else if year < 2010 || year > 2020 {
				return false
			}
		case "eyr":
			year, err := strconv.Atoi(value)
			if err != nil {
				return false
			} else if year < 2020 || year > 2030 {
				return false
			}
		case "hgt":
			regex, _ := regexp.Compile(`(\d+)(in|cm)`)
			if regex.MatchString(value) {
				values := regex.FindStringSubmatch(value)
				height, err := strconv.Atoi(values[1])
				if err != nil {
					return false
				}
				if values[2] == "in" {
					if height < 59 || height > 76 {
						return false
					}
				} else if values[2] == "cm" {
					if height < 150 || height > 193 {
						return false
					}
				} else {
					return false
				}
			} else {
				return false
			}
		case "hcl":
			match, _ := regexp.MatchString(`^#(\d|[a-f]){6}$`, value)
			if !match {
				return false
			}
		case "ecl":
			switch value {
			case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
			default:
				return false
			}
		case "pid":
			match, _ := regexp.MatchString(`^\d{9}$`, value)
			if !match {
				return false
			}
		}
	}
	return true
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	passports := ReadPassports(lines)
	var validPassports []map[string]string
	validCount := 0
	for _, passport := range passports {
		if ValidatePassport(passport) {
			validCount++
			validPassports = append(validPassports, passport)
		}
	}

	fmt.Println("Part 1: There are", validCount, "valid passports")
	validCount2 := 0
	for _, passport := range validPassports {
		if ValidatePassportFields(passport) {
			validCount2++
		}
	}
	fmt.Println("Part 2: There are", validCount2, "valid passports")
}
