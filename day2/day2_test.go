package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScanLine(t *testing.T) {
	line := "1-3 a: abcde"
	var min int
	var max int
	var character string
	var password string
	min, max, character, password = ScanLine(line)

	require.Equal(t, 1, min)
	require.Equal(t, 3, max)
	require.Equal(t, "a", character)
	require.Equal(t, "abcde", password)
}

func TestValidatePassword(t *testing.T) {
	require.True(t, ValidatePassword("1-3 a: abcde"))
	require.False(t, ValidatePassword("1-3 b: cdefg"))
	require.True(t, ValidatePassword("2-9 c: ccccccccc"))
}

func TestValidatePassword2(t *testing.T) {
	require.True(t, ValidatePassword2("1-3 a: abcde"))
	require.False(t, ValidatePassword2("1-3 b: cdefg"))
	require.False(t, ValidatePassword2("2-9 c: ccccccccc"))
}
