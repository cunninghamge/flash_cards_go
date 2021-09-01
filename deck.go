package main

type Deck struct {
	Cards []Card
}

func (d Deck) Count() int {
	return 0
}

func (d Deck) CardsInCategory(cat string) []Card {
	return []Card{}
}
