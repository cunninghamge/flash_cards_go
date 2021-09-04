package main

import (
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	deck := Deck{Cards: []Card{card1, card2}}

	got := deck.Count()
	if got != 2 {
		t.Errorf("got %d want %d", got, 2)
	}

	deck.Cards = append(deck.Cards, card3)

	got = deck.Count()
	if got != 3 {
		t.Errorf("got %d want %d", got, 3)
	}
}

func TestCardsInCategory(t *testing.T) {
	deck := Deck{Cards: []Card{card1, card2, card3}}

	t.Run("one card in category", func(t *testing.T) {
		got := deck.CardsInCategory("Geography")
		want := []Card{card1}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("multiple cards in category", func(t *testing.T) {
		got := deck.CardsInCategory("STEM")
		want := []Card{card2, card3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("no cards in category", func(t *testing.T) {
		got := deck.CardsInCategory("Pop Culture")
		if !reflect.DeepEqual(got, []Card{}) {
			t.Errorf("got %v, want %v", got, []Card{})
		}
	})
}
