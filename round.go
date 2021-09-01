package main

import "math"

type Round struct {
	Deck  Deck
	Turns []Turn
}

func (r Round) CurrentCard() Card {
	return r.Deck.Cards[0]
}

func (r *Round) TakeTurn(guess string) Turn {
	turn := Turn{Guess: guess, Card: r.CurrentCard()}
	r.Turns = append(r.Turns, turn)
	r.Deck.Cards = r.Deck.Cards[1:]
	return turn
}

func (r Round) NumberCorrect() int {
	var count int
	for _, turn := range r.Turns {
		if turn.Correct() {
			count++
		}
	}
	return count
}

func (r Round) NumberCorrectByCategory(category string) int {
	var count int
	for _, turn := range r.Turns {
		if turn.Correct() && turn.Card.Category == category {
			count++
		}
	}
	return count
}

func (r Round) PercentCorrect() float64 {
	pct := float64(r.NumberCorrect()) / float64(len(r.Turns))
	return math.Round(pct*1000) / 10
}

func (r Round) PercentCorrectByCategory(category string) float64 {
	numCorrect := r.NumberCorrectByCategory(category)
	var numInCategory float64
	for _, turn := range r.Turns {
		if turn.Card.Category == category {
			numInCategory++
		}
	}

	if numInCategory == 0 {
		return 0
	}

	pct := float64(numCorrect) / numInCategory
	return math.Round(pct*1000) / 10
}
