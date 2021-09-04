package main

import (
	"math/rand"
	"time"
)

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

func (d Deck) Shuffle() Deck {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
	return d
}
