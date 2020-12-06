package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountQuestions(t *testing.T) {
	input := []string{"abc"}

	require.Equal(t, 3, CountQuestions(input))
	require.Equal(t, 3, CountQuestions([]string{"a", "b", "c"}))
	require.Equal(t, 3, CountQuestions([]string{"ab", "ac"}))
	require.Equal(t, 1, CountQuestions([]string{"a", "a", "a", "a"}))
	require.Equal(t, 1, CountQuestions([]string{"b"}))
}

func TestCountQuestions2(t *testing.T) {
	input := []string{"abc"}

	require.Equal(t, 3, CountQuestions2(input))
	require.Equal(t, 0, CountQuestions2([]string{"a", "b", "c"}))
	require.Equal(t, 1, CountQuestions2([]string{"ab", "ac"}))
	require.Equal(t, 1, CountQuestions2([]string{"a", "a", "a", "a"}))
	require.Equal(t, 1, CountQuestions2([]string{"b"}))
}
