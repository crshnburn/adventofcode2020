package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	name   string
	value  int
	hasRun bool
}

func (i *Instruction) run() {
	i.hasRun = true
}

func ReadInstruction(code string) *Instruction {
	parts := strings.Split(code, " ")
	value, _ := strconv.Atoi(parts[1])
	return &Instruction{name: parts[0], value: value, hasRun: false}
}

func RunProgram(codes []string) int {
	program := []*Instruction{}
	for _, code := range codes {
		program = append(program, ReadInstruction(code))
	}
	ptr := 0
	acc := 0
	var currentInstr *Instruction
	currentInstr = program[ptr]
	for !currentInstr.hasRun {
		// fmt.Println(currentInstr, "|", ptr, acc)
		switch currentInstr.name {
		case "jmp":
			ptr += currentInstr.value
		case "nop":
			ptr++
		case "acc":
			acc += currentInstr.value
			ptr++
		}
		currentInstr.run()
		currentInstr = program[ptr]
	}
	return acc
}

func RunProgramFix(codes []string) int {
	program := []*Instruction{}
	for _, code := range codes {
		program = append(program, ReadInstruction(code))
	}
	ptr := 0
	acc := 0
	completed := false
	switchPos := 0
	var currentInstr *Instruction
	currentInstr = program[ptr]
	for !completed {
		for !currentInstr.hasRun && ptr < len(program) {
			switch currentInstr.name {
			case "jmp":
				if ptr == switchPos {
					ptr++
				} else {
					ptr += currentInstr.value
				}
			case "nop":
				if ptr == switchPos {
					ptr += currentInstr.value
				} else {
					ptr++
				}
			case "acc":
				acc += currentInstr.value
				ptr++
			}
			currentInstr.run()
			if ptr < len(program) {
				currentInstr = program[ptr]
			}
		}
		if ptr < len(program) {
			ptr = 0
			acc = 0
			switchPos++
			currentInstr = program[ptr]
			for _, instr := range program {
				instr.hasRun = false
			}
		} else {
			completed = true
		}
	}
	return acc
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

	fmt.Println("Part 1", RunProgram(lines))
	fmt.Println("Part 2:", RunProgramFix(lines))
}
