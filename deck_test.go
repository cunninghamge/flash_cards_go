package main

import (
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	deck := Deck{Cards: testCards[:2]}

	got := deck.Count()
	if got != 2 {
		t.Errorf("got %d want %d", got, 2)
	}

	deck.Cards = append(deck.Cards, testCards[2])

	got = deck.Count()
	if got != 3 {
		t.Errorf("got %d want %d", got, 3)
	}
}

func TestCardsInCategory(t *testing.T) {
	t.Run("one card in category", func(t *testing.T) {
		got := testDeck.CardsInCategory("Geography")
		want := []Card{testCards[0]}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("multiple cards in category", func(t *testing.T) {
		got := testDeck.CardsInCategory("STEM")
		want := []Card{testCards[1], testCards[2], testCards[5], testCards[6]}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("no cards in category", func(t *testing.T) {
		got := testDeck.CardsInCategory("Pop Culture")
		if !reflect.DeepEqual(got, []Card{}) {
			t.Errorf("got %v, want %v", got, []Card{})
		}
	})
}
