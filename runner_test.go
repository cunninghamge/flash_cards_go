package main

import (
	"bufio"
	"bytes"
	"errors"
	"flash_cards/reader"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestPlayRound(t *testing.T) {
	writer := &bytes.Buffer{}
	reader := &bytes.Buffer{}

	reader.Write([]byte("Fairbanks\ny\nJuneau"))

	playRound([]string{"./fixtures/one_card.csv"}, reader, writer)
	gameLog := writer.String()
	want := `Welcome! You're playing with 1 cards.
----------------------------------------
This is card number 1 out of 1.
Question: What is the capital of Alaska?
Incorrect.
Answer: Juneau

********** Game Over! **********
You had 0 correct guesses out of 1 for a total score of 0.0%.
Geography - 0.0% correct
Retry incorrect guesses? [y/n] This is card number 1 out of 1.
Question: What is the capital of Alaska?
Correct!

Deck Complete!
`
	if gameLog != want {
		t.Errorf("got %s want %s", gameLog, want)
	}
}

func TestNewRound(t *testing.T) {
	t.Run("default deck", func(t *testing.T) {
		defaultRecords, _ := reader.ReadFile("./fixtures/default_cards.csv")
		defaultCards, _ := createCards(defaultRecords)
		sort.Slice(defaultCards, func(i, j int) bool {
			return defaultCards[i].Answer < defaultCards[j].Answer
		})

		round := newRound([]string{}, &bytes.Buffer{})
		sort.Slice(round.Deck.Cards, func(i, j int) bool {
			return round.Deck.Cards[i].Answer < round.Deck.Cards[j].Answer
		})
		got := round.Deck.Cards
		if !reflect.DeepEqual(got, defaultCards) {
			t.Errorf("got %v defaultCards %v", got, defaultCards)
		}
	})

	t.Run("with file source", func(t *testing.T) {
		question := "What is the capital of Alaska?"
		round := newRound([]string{"fixtures/test_cards.csv"}, &bytes.Buffer{})
		for _, card := range round.Deck.Cards {
			if card.Question == question {
				return
			}
		}

		t.Errorf("round does not contain expected card with question %s:\n%v", question, round.Deck)
	})
}

func TestNewRoundExit(t *testing.T) {
	os.WriteFile("tmp.csv", []byte("a,b,c,d\n"), 0644)
	defer os.Remove("tmp.csv")

	testCases := map[string]string{
		"file reader error":  "notarealfile.txt",
		"card creator error": "tmp.csv",
	}

	for name, file := range testCases {
		t.Run(name, func(t *testing.T) {
			if os.Getenv("OS_EXIT_CALLED") == "1" {
				newRound([]string{file}, &bytes.Buffer{})
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
		"Correct!\n\n"
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
		"Incorrect.\n" +
		"Answer: " + testCards[1].Answer + "\n\n"
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

func TestRetryPrompt(t *testing.T) {
	t.Run("repeats until y or n is entered", func(t *testing.T) {
		reader := &bytes.Buffer{}
		writer := &bytes.Buffer{}

		reader.Write([]byte("z\ny"))
		scanner := bufio.NewScanner(reader)

		retryPrompt(writer, scanner)
		got := writer.String()
		want := "Retry incorrect guesses? [y/n] Retry incorrect guesses? [y/n] "
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("exits if n is entered", func(t *testing.T) {
		if os.Getenv("OS_EXIT_CALLED") == "1" {
			reader := &bytes.Buffer{}
			reader.Write([]byte("n"))
			scanner := bufio.NewScanner(reader)
			retryPrompt(&bytes.Buffer{}, scanner)
			return
		}
		subTest := exec.Command(os.Args[0], "-test.run=TestRetryPrompt")
		subTest.Env = append(os.Environ(), "OS_EXIT_CALLED=1")
		err := subTest.Run()
		if exitError, ok := err.(*exec.ExitError); ok && !exitError.Success() {
			t.Error("process exited with error, wanted exit status 0")
		}
	})
}

func TestRetryMissedCards(t *testing.T) {
	round := Round{Turns: []Turn{{Guess: "Anchorage", Card: testCards[0]}}}
	reader := &bytes.Buffer{}
	writer := &bytes.Buffer{}
	reader.Write([]byte("y\nAnchorage\ny\nJuneau"))
	scanner := bufio.NewScanner(reader)
	retryMissedCards(writer, scanner, &round)
	got := writer.String()
	want := `Retry incorrect guesses? [y/n] This is card number 1 out of 1.
Question: What is the capital of Alaska?
Incorrect.
Answer: Juneau

Retry incorrect guesses? [y/n] This is card number 1 out of 1.
Question: What is the capital of Alaska?
Correct!`
	if strings.TrimSpace(got) != strings.TrimSpace(want) {
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
