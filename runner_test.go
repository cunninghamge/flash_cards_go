package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"testing"
)

func TestPlayRound(t *testing.T) {
	writer := &bytes.Buffer{}
	reader := &bytes.Buffer{}

	var answers []string
	for _, card := range testDeck.Cards {
		answers = append(answers, card.Answer)
	}
	reader.Write([]byte(strings.Join(answers, "\n")))
	playRound("", reader, writer)

	gameLog := writer.String()
	prefix := "Welcome! You're playing with 9 cards.\n" + lineBreak + "\n"
	if !strings.HasPrefix(gameLog, prefix) {
		t.Errorf("incorrect welcome: got %s want %s", gameLog, prefix)
	}

	for i, card := range testDeck.Cards {
		want := fmt.Sprintf("This is card number %d out of 9.", i+1) + "\n" +
			"Question: " + card.Question + "\n" +
			"Correct!" + "\n\n"
		if !strings.Contains(gameLog, want) {
			t.Errorf("Game Log %s missing substring %s", gameLog, want)
		}
	}

	suffix := gameOver + "\n" +
		"You had 9 correct guesses out of 9 for a total score of 100.0%.\n" +
		"Geography - 100.0% correct\n" +
		"STEM - 100.0% correct\n" +
		"Sports - 100.0% correct\n" +
		"History - 100.0% correct\n"
	if !strings.HasSuffix(gameLog, suffix) {
		t.Errorf("incorrect summary: got %s want %s", gameLog, suffix)
	}
}

func TestNewRound(t *testing.T) {
	t.Run("default deck", func(t *testing.T) {
		got := newRound("", &bytes.Buffer{})
		want := Round{Deck: testDeck}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("with file source", func(t *testing.T) {
		t.Skip()
		round := newRound("fixtures/cards.csv", &bytes.Buffer{})
		got := round.Deck.Cards[0].Question
		want := "What is the official state sport of Alaska?"
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}

func TestNewRoundExit(t *testing.T) {
	t.Skip()
	os.WriteFile("tmp.csv", []byte("a,b,c,d\n"), 0644)
	defer os.Remove("tmp.csv")

	testCases := map[string]string{
		"file reader error":  "notarealfile.txt",
		"card creator error": "tmp.csv",
	}

	for name, file := range testCases {
		t.Run(name, func(t *testing.T) {
			if os.Getenv("OS_EXIT_CALLED") == "1" {
				newRound(file, &bytes.Buffer{})
				return
			}
			subTest := exec.Command(os.Args[0], "-test.run=TestNewRoundExit")
			subTest.Env = append(os.Environ(), "OS_EXIT_CALLED=1")
			err := subTest.Run()
			if exitError, ok := err.(*exec.ExitError); !ok || exitError.Success() {
				t.Error("process exited with no error, wanted exit status 1")
			}
		})
	}
}

func TestDisplayWelcome(t *testing.T) {
	writer := &bytes.Buffer{}
	count := 5

	displayWelcome(writer, count)

	got := writer.String()
	want := fmt.Sprintf("Welcome! You're playing with %d cards.\n", count) +
		lineBreak + "\n"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestPlayTurn(t *testing.T) {
	deck := Deck{[]Card{testCards[0], testCards[1], testCards[2]}}
	round := Round{Deck: deck}

	writer := &bytes.Buffer{}
	reader := &bytes.Buffer{}
	reader.Write([]byte("Juneau"))
	scanner := bufio.NewScanner(reader)
	playTurn(writer, scanner, 1, 3, &round)

	got := writer.String()
	want := "This is card number 1 out of 3.\n" +
		"Question: " + testCards[0].Question + "\n" +
		"Correct!" + "\n\n"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

	writer.Reset()
	reader.Write([]byte("Venus"))
	scanner = bufio.NewScanner(reader)
	playTurn(writer, scanner, 2, 3, &round)

	got = writer.String()
	want = "This is card number 2 out of 3.\n" +
		"Question: " + testCards[1].Question + "\n" +
		"Incorrect.\n\n"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestDisplaySummary(t *testing.T) {
	deck := Deck{[]Card{testCards[0], testCards[1], testCards[2]}}
	round := Round{Deck: deck}

	round.TakeTurn("Juneau")
	round.TakeTurn("Mars")
	round.TakeTurn("South")

	var writer bytes.Buffer
	displaySummary(&writer, &round)

	got := writer.String()
	want := gameOver + "\n" +
		"You had 2 correct guesses out of 3 for a total score of 66.7%.\n" +
		"Geography - 100.0% correct\n" +
		"STEM - 50.0% correct\n"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestExitWithError(t *testing.T) {
	if os.Getenv("OS_EXIT_CALLED") == "1" {
		exitWithError(errors.New("error"))
		return
	}
	subTest := exec.Command(os.Args[0], "-test.run=TestExitWithError")
	subTest.Env = append(os.Environ(), "OS_EXIT_CALLED=1")
	err := subTest.Run()
	if exitError, ok := err.(*exec.ExitError); !ok || exitError.Success() {
		t.Error("process exited with no error, wanted exit status 1")
	}
}
