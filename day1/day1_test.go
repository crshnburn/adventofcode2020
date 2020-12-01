package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExpenseFind(t *testing.T) {
	expenses := []int{1721, 979, 366, 299, 675, 1456}

	require.Equal(t, 514579, FindExpenses(expenses))
}

func TestThreeExpenseFind(t *testing.T) {
	expenses := []int{1721, 979, 366, 299, 675, 1456}

	require.Equal(t, 241861950, FindThreeExpenses(expenses))
}
