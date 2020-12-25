package main

import "fmt"

func CalcLoop(pubkey int) int {
	subject := 7
	value := 1
	loop := 0
	for value != pubkey {
		value *= subject
		value %= 20201227
		loop++
	}
	return loop
}

func Encrypt(val int, count int) int {
	encrypted := 1
	for i := 0; i < count; i++ {
		encrypted *= val
		encrypted %= 20201227
	}
	return encrypted
}

func main() {
	cardpub := 3469259
	doorpub := 13170438

	// cardLoop := CalcLoop(cardpub)
	doorLoop := CalcLoop(doorpub)

	fmt.Println("Part 1:", Encrypt(cardpub, doorLoop))
}
