package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Chip struct {
	mask   string
	memory map[int]string
}

type Chipv2 struct {
	mask   string
	memory map[int64]int
}

func ReadInt(value int64) string {
	bin := strconv.FormatInt(value, 2)
	return fmt.Sprintf("%036s", bin)
}

func ApplyMask(value string, mask string) string {
	valueArr := []rune(value)
	maskArr := []rune(mask)
	for i, value := range maskArr {
		switch value {
		case '1':
			valueArr[i] = '1'
		case '0':
			valueArr[i] = '0'
		}
	}
	return string(valueArr)
}

func GetAddresses(orig []rune, mask []rune) [][]rune {
	addresses := make([][]rune, 1)
	addresses[0] = []rune{}

	for z, bit := range mask {
		// fmt.Println(bit)
		switch bit {
		case '0':
			for j := range addresses {
				addresses[j] = append(addresses[j], orig[z])
			}
		case '1':
			for j := range addresses {
				addresses[j] = append(addresses[j], '1')
			}
		case 'X':
			lenVal := len(addresses)
			for i := 0; i < lenVal; i++ {
				newArr := make([]rune, len(addresses[i]))
				copy(newArr, addresses[i])
				addresses[i] = append(addresses[i], '0')
				newArr = append(newArr, '1')
				addresses = append(addresses, newArr)
			}

		}
	}
	return addresses
}

func NewChip() *Chip {
	memory := make(map[int]string)
	return &Chip{mask: "", memory: memory}
}

func NewChipv2() *Chipv2 {
	memory := make(map[int64]int)
	return &Chipv2{mask: "", memory: memory}
}
func (c *Chip) ReadInstruction(instr string) {
	parts := strings.Split(instr, " = ")
	if parts[0] == "mask" {
		c.mask = parts[1]
	} else {
		var pos int
		fmt.Sscanf(parts[0], "mem[%d]", &pos)
		value, _ := strconv.Atoi(parts[1])
		c.memory[pos] = ApplyMask(ReadInt(int64(value)), c.mask)
	}
}

func (c *Chipv2) ReadInstruction(instr string) {
	parts := strings.Split(instr, " = ")
	if parts[0] == "mask" {
		c.mask = parts[1]
	} else {
		var pos int
		fmt.Sscanf(parts[0], "mem[%d]", &pos)
		value, _ := strconv.Atoi(parts[1])
		addresses := GetAddresses([]rune(ReadInt(int64(pos))), []rune(c.mask))
		for _, address := range addresses {
			val, _ := strconv.ParseInt(string(address), 2, 64)
			c.memory[val] = value
		}
	}
}

func (c *Chip) Total() int64 {
	total := int64(0)
	for _, value := range c.memory {
		val, _ := strconv.ParseInt(value, 2, 64)
		total += val
	}
	return total
}

func (c *Chipv2) Total() int {
	total := 0
	for _, value := range c.memory {
		total += value
	}
	return total
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	chip := NewChip()
	chipv2 := NewChipv2()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instr := scanner.Text()
		chip.ReadInstruction(instr)
		chipv2.ReadInstruction(instr)
	}
	fmt.Println("Part 1:", chip.Total())
	fmt.Println("Part 2:", chipv2.Total())
}
