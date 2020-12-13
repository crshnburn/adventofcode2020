package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func FindTime(start int, busIds []int) (int, int) {
	busTime := start
	timeFound := false
	for !timeFound {
		for _, id := range busIds {
			if busTime%id == 0 {
				return busTime, id
			}
		}
		busTime++
	}
	return -1, -1
}

func FindSpecialTime(busIdMap map[int64]int64) int64 {
	timestamp := int64(0)
	inc := busIdMap[0]
	var offsets []int64
	for offset, _ := range busIdMap {
		// fmt.Println(offset)
		offsets = append(offsets, offset)
	}
	// fmt.Println(offsets)
	for i := 1; i <= len(offsets); i++ {
		// fmt.Println(offsets[:i])
		for !checkValue(timestamp, offsets[:i], busIdMap) {
			timestamp += inc
		}
		inc = multiple(offsets[:i], busIdMap)
	}

	return timestamp
}

func multiple(offsets []int64, buses map[int64]int64) int64 {
	mult := int64(1)
	for _, offset := range offsets {
		mult *= buses[int64(offset)]
	}
	return mult
}

func checkValue(timestamp int64, offsets []int64, buses map[int64]int64) bool {
	for _, offset := range offsets {
		// fmt.Println(i, offset, buses[int64(i)])
		if (timestamp+offset)%buses[offset] != 0 {
			return false
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
	scanner.Scan()
	startTime, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	idString := scanner.Text()
	var busIds []int
	busIdMap := make(map[int64]int64)
	for i, id := range strings.Split(idString, ",") {
		busId, err := strconv.Atoi(id)
		if err == nil {
			busIds = append(busIds, busId)
			busIdMap[int64(i)] = int64(busId)
		}
	}
	busTime, id := FindTime(startTime, busIds)
	fmt.Println("Part 1:", (busTime-startTime)*id)
	fmt.Println("Part 2:", FindSpecialTime(busIdMap))
}
