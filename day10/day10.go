package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func JoltRange(adapters []int) (int, int, int) {
	noOnes := 0
	noTwos := 0
	noThrees := 1
	currentJolt := 0
	sort.Ints(adapters)
	for _, adapter := range adapters {
		diff := adapter - currentJolt
		switch diff {
		case 1:
			noOnes++
		case 2:
			noTwos++
		case 3:
			noThrees++
		}
		currentJolt = adapter
	}
	return noOnes, noTwos, noThrees
}

func JoltCombos(adapters []int) int {
	adapterMap := map[int]bool{}
	sort.Ints(adapters)
	for _, adapter := range adapters {
		adapterMap[adapter] = true
	}
	deviceJolts := adapters[len(adapters)-1]

	cache := map[string]int{}
	adapterMap[deviceJolts] = true

	return search(0, deviceJolts, adapterMap, cache)
}

func search(jolt, deviceJolts int, adapters map[int]bool, cache map[string]int) int {
	if jolt == deviceJolts {
		return 1
	}
	count := 0
	for diff := 1; diff <= 3; diff++ {
		if adapters[jolt+diff] {
			key := fmt.Sprintf("%d,%d", jolt, diff)
			if _, found := cache[key]; !found {
				subCount := search(jolt+diff, deviceJolts, adapters, cache)
				cache[key] = subCount
			}
			count += cache[key]
		}
	}
	return count
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var adapters []int
	for scanner.Scan() {
		adapter, _ := strconv.Atoi(scanner.Text())
		adapters = append(adapters, adapter)
	}

	noOnes, _, noThrees := JoltRange(adapters)
	fmt.Println("Part 1:", noOnes*noThrees)
	fmt.Println("Part 2:", JoltCombos(adapters))
}
