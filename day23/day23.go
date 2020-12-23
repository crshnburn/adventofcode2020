package main

import (
	"container/ring"
	"fmt"
)

func findPos(circle []int, cup int) int {
	for i, val := range circle {
		if val == cup {
			return i
		}
	}
	return -1
}

func GetNextThree(circle []int, cup int) []int {
	cupPos := findPos(circle, cup)
	nextcups := []int{}
	for i := cupPos + 1; i <= cupPos+3; i++ {
		if i >= len(circle) {
			nextcups = append(nextcups, circle[i-len(circle)])
		} else {
			nextcups = append(nextcups, circle[i])
		}
	}
	return nextcups
}

func RemoveFound(circle []int, cup int) []int {
	toRemove := GetNextThree(circle, cup)
	for _, remove := range toRemove {
		circle = removeCup(circle, remove)
	}
	return circle
}

func FindInsert(circle []int, cup int) int {
	toFind := (cup - 1) % (len(circle) + 3)
	pos := -1
	for pos == -1 {
		pos = findPos(circle, toFind)
		toFind = toFind - 1
		if toFind < 1 {
			toFind = len(circle) + 3
		}
	}
	return pos
}

func InsertCups(circle []int, cup int, removed []int) []int {
	insertPos := FindInsert(circle, cup)
	newCircle := []int{}
	newCircle = append(newCircle, circle[:insertPos+1]...)
	newCircle = append(newCircle, removed...)
	newCircle = append(newCircle, circle[insertPos+1:]...)
	return newCircle
}

func removeCup(circle []int, cup int) []int {
	pos := findPos(circle, cup)
	return append(circle[:pos], circle[pos+1:]...)
}

func Game(circle []int, rounds int) []int {
	pos := 0
	for i := 0; i < rounds; i++ {
		// fmt.Println("cups:", circle)
		cup := circle[pos]
		// fmt.Println("Cup:", pos, cup)
		removed := GetNextThree(circle, cup)
		// fmt.Println("Pick up:", removed)
		circle = RemoveFound(circle, cup)
		circle = InsertCups(circle, cup, removed)
		pos = (findPos(circle, cup) + 1) % len(circle)
		// fmt.Println()
	}
	return circle
}

func answerString(circle []int) string {
	posOne := findPos(circle, 1)
	answer := ""
	// fmt.Println(circle)
	for i := 1; i < len(circle); i++ {
		answer += fmt.Sprint(circle[(posOne+i)%len(circle)])
	}
	return answer
}

func RingGame(start []int, iterations int) *ring.Ring {
	cups := ring.New(len(start))
	cupPos := make([]*ring.Ring, len(start))
	for _, cup := range start {
		cups.Value = cup
		cupPos[cup-1] = cups
		cups = cups.Next()
	}
	for i := 0; i < iterations; i++ {
		removed := cups.Unlink(3)
		removedVals := []int{}
		for i := 0; i < removed.Len(); i++ {
			removedVals = append(removedVals, removed.Value.(int))
			removed = removed.Next()
		}
		// fmt.Println(removedVals)
		cupVal := cups.Value.(int) - 1
		if cupVal < 1 {
			cupVal = cups.Len() + 3
		}
		// fmt.Println(cupVal, removedVals)
		for contains(removedVals, cupVal) {
			if cupVal <= 1 {
				cupVal = cups.Len() + 3
			} else {
				cupVal--
			}
		}
		// fmt.Println(cupVal)

		cupPos[cupVal-1].Link(removed)
		cups = cups.Next()
	}
	return cupPos[0]
}

func part1Answer(cups *ring.Ring) string {
	answer := ""
	for i := 0; i < 8; i++ {
		cups = cups.Next()
		answer += fmt.Sprint(cups.Value)
	}
	return answer
}

func part2Answer(posOne *ring.Ring) int {
	return posOne.Next().Value.(int) * posOne.Next().Next().Value.(int)
}

func contains(list []int, val int) bool {
	for _, l := range list {
		if val == l {
			return true
		}
	}
	return false
}

func main() {
	circle := []int{2, 1, 9, 7, 4, 8, 3, 6, 5}
	fmt.Println("Part 1:", answerString(Game(circle, 100)))
	circle = []int{2, 1, 9, 7, 4, 8, 3, 6, 5}
	fmt.Println("Part 1(ring):", part1Answer(RingGame(circle, 100)))
	circle = []int{2, 1, 9, 7, 4, 8, 3, 6, 5}
	for i := 10; i <= 1000000; i++ {
		circle = append(circle, i)
	}
	fmt.Println("Part 2:", part2Answer(RingGame(circle, 10000000)))
}
