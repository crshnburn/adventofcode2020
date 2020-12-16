package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsIn(t *testing.T) {
	myRange := &Range{low: 1, high: 3}
	require.True(t, myRange.IsIn(3))
	require.False(t, myRange.IsIn(4))
}

func TestIsValid(t *testing.T) {
	myRange := &Range{low: 1, high: 3}
	myRange1 := &Range{low: 5, high: 7}
	field := &Field{name: "class", ranges: []*Range{myRange, myRange1}}
	require.True(t, field.IsValid(3))
	require.False(t, field.IsValid(4))
	require.True(t, field.IsValid(5))
}

func TestParseInput(t *testing.T) {
	input := []string{
		"class: 1-3 or 5-7",
		"row: 6-11 or 33-44",
		"seat: 13-40 or 45-50",
		"",
		"your ticket:",
		"7,1,14",
		"",
		"nearby tickets:",
		"7,3,47",
		"40,4,50",
		"55,2,20",
		"38,6,12",
	}

	ticket, myticket, scannedTickets := ParseInput(input)
	require.Equal(t, 3, len(ticket.fields))
	require.Equal(t, "7,1,14", myticket)
	require.Equal(t, "7,3,47", scannedTickets[0])
}

func TestParseField(t *testing.T) {
	input := "class: 1-3 or 5-7"
	field := ParseField(input)
	require.Equal(t, "class", field.name)

	require.Equal(t, 1, field.ranges[0].low)
	require.Equal(t, 3, field.ranges[0].high)
	require.Equal(t, 5, field.ranges[1].low)
	require.Equal(t, 7, field.ranges[1].high)
}

func TestCalcErrorRate(t *testing.T) {
	input := []string{
		"class: 1-3 or 5-7",
		"row: 6-11 or 33-44",
		"seat: 13-40 or 45-50",
		"",
		"your ticket:",
		"7,1,14",
		"",
		"nearby tickets:",
		"7,3,47",
		"40,4,50",
		"55,2,20",
		"38,6,12",
	}

	require.Equal(t, 71, CalcErrorRate(input))
}

func TestCalcField(t *testing.T) {
	input := []string{
		"class: 0-1 or 4-19",
		"row: 0-5 or 8-19",
		"seat: 0-13 or 16-19",
		"",
		"your ticket:",
		"11,12,13",
		"",
		"nearby tickets:",
		"3,9,18",
		"15,1,5",
		"5,14,9",
	}

	ticket, _, scannedTickets := ParseInput(input)

	fieldMap := make(map[string]*Field)
	// fieldOrder := []*Field{}
	for _, field := range ticket.fields {
		fieldMap[field.name] = field
	}
	var col []string
	for _, tick := range scannedTickets {
		parts := strings.Split(tick, ",")
		col = append(col, parts[0])
	}
	require.Equal(t, "row", CalcField(fieldMap, col))
}

func TestCalcOrder(t *testing.T) {
	input := []string{
		"class: 0-1 or 4-19",
		"row: 0-5 or 8-19",
		"seat: 0-13 or 16-19",
		"",
		"your ticket:",
		"11,12,13",
		"",
		"nearby tickets:",
		"3,9,18",
		"15,1,5",
		"5,14,9",
	}

	ticket, _, scannedTickets := ParseInput(input)
	rows := CalcOrder(ticket, scannedTickets)
	fmt.Println(rows)
}
