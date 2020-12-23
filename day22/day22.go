package main

import (
	"fmt"
	"strconv"
)

type Deck struct {
	cards []int
}

type Game struct {
	player1 *Deck
	player2 *Deck
}

type GameState struct {
	player1 int
	player2 int
}

func (d *Deck) topCard() int {
	topCard := d.cards[0]
	d.cards = d.cards[1:]
	return topCard
}

func CreateDeck(input []string) *Deck {
	cards := []int{}
	for i := 1; i < len(input); i++ {
		val, _ := strconv.Atoi(input[i])
		cards = append(cards, val)
	}
	return &Deck{cards: cards}
}

func (d *Deck) addCards(a, b int) {
	d.cards = append(d.cards, a)
	d.cards = append(d.cards, b)
}

func (d *Deck) Score() int {
	score := 0
	for i := 1; i <= len(d.cards); i++ {
		score += i * d.cards[len(d.cards)-i]
	}
	return score
}

func (d *Deck) hash() int {
	hash := 0
	for _, card := range d.cards {
		hash += card
	}
	return hash
}

func contains(states []GameState, state GameState) bool {
	for _, local := range states {
		if local.player1 == state.player1 && local.player2 == state.player2 {
			return true
		}
	}
	return false
}

func (g *Game) PlayGame() int {
	for len(g.player1.cards) > 0 && len(g.player2.cards) > 0 {
		card1 := g.player1.topCard()
		card2 := g.player2.topCard()
		if card1 > card2 {
			g.player1.addCards(card1, card2)
		} else {
			g.player2.addCards(card2, card1)
		}
	}
	if len(g.player1.cards) == 0 {
		return g.player2.Score()
	} else {
		return g.player1.Score()
	}
}

func (g *Game) RecursiveCombat() bool {
	previousRounds := []GameState{}
	for len(g.player1.cards) > 0 && len(g.player2.cards) > 0 {
		state := GameState{player1: g.player1.Score(), player2: g.player2.Score()}
		if contains(previousRounds, state) {
			return true
		}
		previousRounds = append(previousRounds, state)

		// fmt.Println("Player 1 Deck:", g.player1)
		// fmt.Println("Player 2 Deck:", g.player2)
		card1 := g.player1.topCard()
		card2 := g.player2.topCard()
		// fmt.Println("Player 1 plays:", card1)
		// fmt.Println("Player 2 plays:", card2)
		if len(g.player1.cards) >= card1 && len(g.player2.cards) >= card2 {
			subdeck1 := &Deck{cards: makeSubDeck(g.player1.cards, card1)}
			subdeck2 := &Deck{cards: makeSubDeck(g.player2.cards, card2)}
			// fmt.Println("**Player 1 Deck:", g.player1)
			// fmt.Println("**Player 2 Deck:", g.player2)
			subgame := &Game{player1: subdeck1, player2: subdeck2}
			play1wins := subgame.RecursiveCombat()
			// fmt.Println(play1wins, card1, card2)
			// fmt.Println("***Player 1 Deck:", g.player1)
			// fmt.Println("***Player 2 Deck:", g.player2)
			if play1wins {
				g.player1.addCards(card1, card2)
			} else {
				g.player2.addCards(card2, card1)
			}
		} else {
			if card1 > card2 {
				g.player1.addCards(card1, card2)
			} else {
				g.player2.addCards(card2, card1)
			}
		}
		// fmt.Println()
	}
	return len(g.player1.cards) != 0
}

func makeSubDeck(cards []int, num int) []int {
	newDeck := []int{}
	for i := 0; i < num; i++ {
		newDeck = append(newDeck, cards[i])
	}
	return newDeck
}

func main() {
	player1 := []string{
		"Player 1:",
		"23",
		"32",
		"46",
		"47",
		"27",
		"35",
		"1",
		"16",
		"37",
		"50",
		"15",
		"11",
		"14",
		"31",
		"4",
		"38",
		"21",
		"39",
		"26",
		"22",
		"3",
		"2",
		"8",
		"45",
		"19",
	}

	player2 := []string{
		"Player 2:",
		"13",
		"20",
		"12",
		"28",
		"9",
		"10",
		"30",
		"25",
		"18",
		"36",
		"48",
		"41",
		"29",
		"24",
		"49",
		"33",
		"44",
		"40",
		"6",
		"34",
		"7",
		"43",
		"42",
		"17",
		"5",
	}

	game := &Game{player1: CreateDeck(player1), player2: CreateDeck(player2)}
	fmt.Println("Part 1:", game.PlayGame())
	game = &Game{player1: CreateDeck(player1), player2: CreateDeck(player2)}
	player1wins := game.RecursiveCombat()
	score := 0
	if player1wins {
		score = game.player1.Score()
	} else {
		score = game.player2.Score()
	}
	fmt.Println("Part 2:", score)

}
