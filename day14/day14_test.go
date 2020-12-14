package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadInt(t *testing.T) {
	value := int64(11)

	require.Equal(t, "000000000000000000000000000000001011", ReadInt(value))
}

func TestApplyMask(t *testing.T) {
	mask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"
	value := int64(11)

	require.Equal(t, "000000000000000000000000000001001001", ApplyMask(ReadInt(value), mask))
}

func TestChip(t *testing.T) {
	chip := NewChip()
	instructions := []string{
		"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
		"mem[8] = 11",
		"mem[7] = 101",
		"mem[8] = 0",
	}
	for _, instr := range instructions {
		chip.ReadInstruction(instr)
	}
	require.Equal(t, int64(165), chip.Total())
}

func TestGetAddresses(t *testing.T) {
	mask := "000000000000000000000000000000X1001X"
	value := int64(42)

	addresses := GetAddresses([]rune(ReadInt(value)), []rune(mask))
	for _, address := range addresses {
		fmt.Println(string(address))
	}
}

func TestChipv2(t *testing.T) {
	chip := NewChipv2()
	instructions := []string{
		"mask = 000000000000000000000000000000X1001X",
		"mem[42] = 100",
		"mask = 00000000000000000000000000000000X0XX",
		"mem[26] = 1",
	}
	for _, instr := range instructions {
		chip.ReadInstruction(instr)
	}
	require.Equal(t, 208, chip.Total())
}
