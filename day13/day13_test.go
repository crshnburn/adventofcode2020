package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindTime(t *testing.T) {
	start := 939
	busIds := []int{7, 13, 59, 31, 19}
	busTime, id := FindTime(start, busIds)
	require.Equal(t, 295, (busTime-start)*id)
}

func TestFindSpecialTime(t *testing.T) {
	busIdMap := make(map[int64]int64)
	busIdMap[0] = 7
	busIdMap[1] = 13
	busIdMap[4] = 59
	busIdMap[6] = 31
	busIdMap[7] = 19
	timestamp := FindSpecialTime(busIdMap)
	require.Equal(t, int64(1068781), timestamp)
}
