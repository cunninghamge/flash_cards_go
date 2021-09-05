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

func playRound(osArgs []string, reader io.Reader, writer io.Writer) {
	round := newRound(osArgs, writer)
	roundLength := round.Deck.Count()

	displayWelcome(writer, roundLength)
	scanner := bufio.NewScanner(reader)

	for i := 0; i < roundLength; i++ {
		playTurn(writer, scanner, i+1, roundLength, &round)
	}

	displaySummary(writer, &round)
	if round.PercentCorrect() < 100 {
		retryMissedCards(writer, scanner, &round)
	}
	fmt.Fprintln(writer, "Deck Complete!")
}

func newRound(osArgs []string, writer io.Writer) Round {
	var source = "./fixtures/default_cards.csv"
	if len(osArgs) > 0 {
		source = osArgs[0]
	}

	records, err := reader.ReadFile(source)
	if err != nil {
		exitWithError(err)
	}
	cards, err := createCards(records)
	if err != nil {
		exitWithError(err)
	}

	return Round{Deck: Deck{cards}.Shuffle()}
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
	fmt.Fprintln(w, turn.Feedback())
	if !turn.Correct() {
		fmt.Fprintf(w, fmt.Sprintf("Answer: %s\n", turn.Card.Answer))
	}
	fmt.Fprint(w, "\n")
}

func displaySummary(w io.Writer, round *Round) {
	fmt.Fprintln(w, gameOver)
	fmt.Fprintf(w, "You had %d correct guesses out of %d for a total score of %.1f%%.\n", round.NumberCorrect(), len(round.Turns), round.PercentCorrect())
	categories := round.ListCategories()
	for _, category := range categories {
		fmt.Fprintf(w, "%s - %.1f%% correct\n", category, round.PercentCorrectByCategory(category))
	}
}

func retryPrompt(w io.Writer, s *bufio.Scanner) {
	for {
		fmt.Fprint(w, "Retry incorrect guesses? [y/n] ")
		s.Scan()
		answer := s.Text()
		if strings.ToLower(answer) == "n" {
			os.Exit(0)
		}
		if strings.ToLower(answer) == "y" {
			return
		}
	}
}

func retryMissedCards(w io.Writer, s *bufio.Scanner, round *Round) {
	retryPrompt(w, s)

	subRound := round.MissedGuesses()
	roundLength := subRound.Deck.Count()
	for i := 0; i < roundLength; i++ {
		playTurn(w, s, i+1, roundLength, &subRound)
	}
	if subRound.PercentCorrect() < 100 {
		retryMissedCards(w, s, &subRound)
	}
}

func exitWithError(err error) {
	fmt.Printf("ERROR: %v\n", err)
	os.Exit(1)
}
