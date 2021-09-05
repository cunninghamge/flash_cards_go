package main

import (
	"reflect"
	"testing"
)

func TestCurrentCard(t *testing.T) {
	round := Round{Deck: testDeck}

	got := round.CurrentCard()
	if !reflect.DeepEqual(got, testCards[0]) {
		t.Errorf("got %v, want %v", got, testCards[0])
	}
}

func TestTakeTurn(t *testing.T) {
	deck := Deck{[]Card{testCards[0], testCards[1], testCards[2]}}
	round := Round{Deck: deck}

	turn := round.TakeTurn("Juneau")
	if !reflect.DeepEqual(turn.Card, testCards[0]) {
		t.Errorf("got %v, want %v for turn's card", turn.Card, testCards[0])
	}

	if !turn.Correct() && turn.Feedback() != "Correct!" {
		t.Errorf("%v was incorrect but should have been correct", turn)
	}

	if !reflect.DeepEqual(round.Turns, []Turn{turn}) {
		t.Errorf("got %v want %v for turns taken", round.Turns, []Turn{turn})
	}

	if !reflect.DeepEqual(round.CurrentCard(), testCards[1]) {
		t.Errorf("got %v want %v for next card", round.CurrentCard(), testCards[1])
	}
}

func TestNumberCorrect(t *testing.T) {
	deck := Deck{[]Card{testCards[0], testCards[1], testCards[2]}}
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
	deck := Deck{[]Card{testCards[0], testCards[1], testCards[2]}}
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
	deck := Deck{[]Card{testCards[0], testCards[1], testCards[2]}}
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
	deck := Deck{[]Card{testCards[0], testCards[1], testCards[2]}}
	round := Round{Deck: deck}

	round.TakeTurn("Juneau")           //correct
	round.TakeTurn("Venus")            //incorrect
	round.TakeTurn("North north west") //correct

	testCases := map[string]float64{
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

func TestListCategories(t *testing.T) {
	deck := Deck{[]Card{testCards[0], testCards[1], testCards[2]}}
	round := Round{Deck: deck}
	for i := 0; i < 3; i++ {
		round.TakeTurn("doesn't matter")
	}

	got := round.ListCategories()
	want := []string{"Geography", "STEM"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestMissedGuesses(t *testing.T) {
	deck := Deck{testCards}
	round := Round{Deck: deck}

	round.TakeTurn("Juneau")
	round.TakeTurn("Mars")
	round.TakeTurn("South")

	newRound := round.MissedGuesses()
	got := newRound.Deck.Cards[0]
	want := testCards[2]
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
