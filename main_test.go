package main

import (
	"encoding/csv"
	"io"
	"os"
	"testing"
)

var (
	testDeck  Deck
	testCards []Card
)

func TestMain(m *testing.M) {
	file, _ := os.Open("./fixtures/default_cards.csv")
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		testCards = append(testCards, Card{record[0], record[1], record[2]})
	}

	testDeck = Deck{testCards}

	code := m.Run()
	os.Exit(code)
}
