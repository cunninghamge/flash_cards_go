package main

type Round struct {
	Deck  Deck
	Turns []Turn
}

func (r Round) CurrentCard() Card {
	return Card{}
}

func (r Round) TakeTurn(guess string) Turn {
	return Turn{}
}

func (r Round) NumberCorrect() int {
	return 0
}

func (r Round) NumberCorrectByCategory(category string) int {
	return 0
}

func (r Round) PercentCorrect() float32 {
	return 0
}

func (r Round) PercentCorrectByCategory(category string) float32 {
	return 0
}
