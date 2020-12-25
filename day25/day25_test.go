package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalcLoop(t *testing.T) {
	cardpub := 5764801

	require.Equal(t, 8, CalcLoop(cardpub))
	require.Equal(t, 11, CalcLoop(17807724))
}
