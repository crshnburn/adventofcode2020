package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	input := []string{
		"0: 4 1 5",
		"1: 2 3 | 3 2",
		"2: 4 4 | 5 5",
		"3: 4 5 | 5 4",
		"4: \"a\"",
		"5: \"b\"",
	}
	rules := Parse(input)

	require.Equal(t, "4 1 5", rules["0"])
}

func TestGetRule(t *testing.T) {
	input := []string{
		"0: 4 1 5",
		"1: 2 3 | 3 2",
		"2: 4 4 | 5 5",
		"3: 4 5 | 5 4",
		"4: \"a\"",
		"5: \"b\"",
		"6: 4 5",
	}
	rules := Parse(input)
	// parts := strings.Split("4 5 | 5 4", " | ")
	// for _, part := range parts {
	// 	fmt.Println(part)
	// }
	fmt.Println(GetRule(rules, "0"))
}
