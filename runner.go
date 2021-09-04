package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"flash_cards/reader"
)

var (
	lineBreak = strings.Repeat("-", 40)
	gameOver  = strings.Repeat("*", 10) + " Game Over! " + strings.Repeat("*", 10)
)

func playRound(cardSource string, reader io.Reader, writer io.Writer) {
	round := newRound(cardSource, writer)
	roundLength := round.Deck.Count()

	displayWelcome(writer, roundLength)
	scanner := bufio.NewScanner(reader)

	for i := 0; i < roundLength; i++ {
		playTurn(writer, scanner, i+1, roundLength, &round)
	}

	displaySummary(writer, &round)
}

func newRound(source string, writer io.Writer) Round {
	if len(source) < 1 {
		source = "./fixtures/default_cards.csv"
	}
	records, err := reader.ReadFile(source)
	if err != nil {
		exitWithError(err)
	}
	cards, err := createCards(records)
	if err != nil {
		exitWithError(err)
	}
	return Round{Deck: Deck{cards}}
}

func displayWelcome(w io.Writer, count int) {
	fmt.Fprintf(w, "Welcome! You're playing with %d cards.\n", count)
	fmt.Fprintln(w, lineBreak)
}

func playTurn(w io.Writer, s *bufio.Scanner, cardNumber, roundLength int, round *Round) {
	fmt.Fprintf(w, "This is card number %d out of %d.\n", cardNumber, roundLength)
	fmt.Fprintf(w, "Question: %s\n", round.CurrentCard().Question)

	s.Scan()
	answer := s.Text()

	turn := round.TakeTurn(answer)
	fmt.Fprintln(w, turn.Feedback()+"\n")
}

func displaySummary(w io.Writer, round *Round) {
	fmt.Fprintln(w, gameOver)
	fmt.Fprintf(w, "You had %d correct guesses out of %d for a total score of %.1f%%.\n", round.NumberCorrect(), len(round.Turns), round.PercentCorrect())
	categories := round.ListCategories()
	for _, category := range categories {
		fmt.Fprintf(w, "%s - %.1f%% correct\n", category, round.PercentCorrectByCategory(category))
	}
}

func exitWithError(err error) {
	fmt.Printf("ERROR: %v\n", err)
	os.Exit(1)
}
