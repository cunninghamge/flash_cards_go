package main

import (
	"reflect"
	"testing"
)

func TestCurrentCard(t *testing.T) {
	deck := Deck{[]Card{card1, card2, card3}}
	round := Round{Deck: deck}

	got := round.CurrentCard()
	if !reflect.DeepEqual(got, card1) {
		t.Errorf("got %v, want %v", got, card1)
	}
}

func TestTakeTurn(t *testing.T) {
	deck := Deck{[]Card{card1, card2, card3}}
	round := Round{Deck: deck}

	turn := round.TakeTurn("Juneau")
	if !turn.Correct() {
		t.Errorf("%v was incorrect but should have been correct", turn)
	}

	if !reflect.DeepEqual(round.Turns, []Turn{turn}) {
		t.Errorf("got %v want %v for turns taken", round.Turns, []Turn{turn})
	}

	if !reflect.DeepEqual(round.CurrentCard(), card2) {
		t.Errorf("got %v want %v for next card", round.CurrentCard(), card2)
	}
}

func TestNumberCorrect(t *testing.T) {
	deck := Deck{[]Card{card1, card2, card3}}
	round := Round{Deck: deck}

	round.TakeTurn("Juneau") //correct

	if round.NumberCorrect() != 1 {
		t.Errorf("got %d want 1", round.NumberCorrect())
	}

	round.TakeTurn("Venus") //incorrect

	if round.NumberCorrect() != 1 {
		t.Errorf("got %d want 1", round.NumberCorrect())
	}
}

func TestNumberCorrectByCategory(t *testing.T) {
	deck := Deck{[]Card{card1, card2, card3}}
	round := Round{Deck: deck}

	round.TakeTurn("Juneau") //correct
	round.TakeTurn("Venus")  //incorrect

	got := round.NumberCorrectByCategory("Geography")
	if got != 1 {
		t.Errorf("got %d want %d", got, 1)
	}

	got = round.NumberCorrectByCategory("STEM")
	if got != 0 {
		t.Errorf("got %d want %d", got, 0)
	}
}

func TestPercentCorrect(t *testing.T) {
	deck := Deck{[]Card{card1, card2, card3}}
	round := Round{Deck: deck}

	round.TakeTurn("Juneau") //correct
	round.TakeTurn("Venus")  //incorrect

	got := round.PercentCorrect()
	if got != 50.0 {
		t.Errorf("got %f want %f", got, 50.0)
	}

	round.TakeTurn("North north west") //correct

	got = round.PercentCorrect()
	if got != 66.7 {
		t.Errorf("got %f want %f", got, 66.7)
	}
}

func TestPercentCorrectByCategory(t *testing.T) {
	deck := Deck{[]Card{card1, card2, card3}}
	round := Round{Deck: deck}

	round.TakeTurn("Juneau")           //correct
	round.TakeTurn("Venus")            //incorrect
	round.TakeTurn("North north west") //correct

	testCases := map[string]float32{
		"Geography":   100.0,
		"STEM":        50.0,
		"Pop Culture": 0.0,
	}

	for category, want := range testCases {
		got := round.PercentCorrectByCategory(category)
		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	}
}
