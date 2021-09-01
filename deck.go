package main

type Deck struct {
	Cards []Card
}

func (d Deck) Count() int {
	return len(d.Cards)
}

func (d Deck) CardsInCategory(category string) []Card {
	var cards = []Card{}
	for _, card := range d.Cards {
		if card.Category == category {
			cards = append(cards, card)
		}
	}
	return cards
}
