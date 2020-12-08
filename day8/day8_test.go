package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadInstruction(t *testing.T) {
	code := "nop +0"

	instr := ReadInstruction(code)

	require.Equal(t, "nop", instr.name)
	require.Equal(t, 0, instr.value)
	require.Equal(t, false, instr.hasRun)
}

func TestRunProgram(t *testing.T) {
	codes := []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}

	require.Equal(t, 5, RunProgram(codes))
}

func TestRunProgramFix(t *testing.T) {
	codes := []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}

	require.Equal(t, 8, RunProgramFix(codes))
}
