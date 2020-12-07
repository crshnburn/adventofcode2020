package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseRules(t *testing.T) {
	rule := "light red bags contain 1 bright white bag, 2 muted yellow bags."

	bags := ParseRules([]string{rule})

	require.Equal(t, "light red bags", bags["light red bags"].Name)
}
