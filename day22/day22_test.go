package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateDeck(t *testing.T) {
	input := []string{
		"Player 1:",
		"9",
		"2",
		"6",
		"3",
		"1",
	}
	deck := CreateDeck(input)
	require.Equal(t, 9, deck.cards[0])
}

func TestAddCards(t *testing.T) {
	deck := &Deck{cards: []int{1, 2, 3, 4, 5}}
	deck.addCards(4, 3)
	fmt.Println(deck)
	require.Equal(t, 4, deck.cards[len(deck.cards)-2])
	require.Equal(t, 3, deck.cards[len(deck.cards)-1])
}

func TestTopCard(t *testing.T) {
	deck := &Deck{cards: []int{1, 2, 3, 4, 5}}
	require.Equal(t, 1, deck.topCard())
	require.Equal(t, 2, deck.cards[0])
}

func TestScore(t *testing.T) {
	deck := &Deck{cards: []int{3, 2, 10, 6, 8, 5, 9, 4, 7, 1}}
	require.Equal(t, 306, deck.Score())
}

func TestPlayGame(t *testing.T) {
	player1 := []string{
		"Player 1:",
		"9",
		"2",
		"6",
		"3",
		"1",
	}
	player2 := []string{
		"Player 2:",
		"5",
		"8",
		"4",
		"7",
		"10",
	}
	deck1 := CreateDeck(player1)
	deck2 := CreateDeck(player2)
	game := &Game{player1: deck1, player2: deck2}
	require.Equal(t, 306, game.PlayGame())
}

func TestPlayRecursiveCombat(t *testing.T) {
	player1 := []string{
		"Player 1:",
		"9",
		"2",
		"6",
		"3",
		"1",
	}
	player2 := []string{
		"Player 2:",
		"5",
		"8",
		"4",
		"7",
		"10",
	}
	deck1 := CreateDeck(player1)
	deck2 := CreateDeck(player2)
	game := &Game{player1: deck1, player2: deck2}
	player1wins := game.RecursiveCombat()
	score := 0
	if player1wins {
		score = game.player1.Score()
	} else {
		score = game.player2.Score()
	}
	require.Equal(t, 291, score)
}
