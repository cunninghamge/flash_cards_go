package main

import (
	"encoding/csv"
	"io"
	"os"
)

func createCardsFromFile(filepath string) ([]Card, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	var cards []Card
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		cards = append(cards, Card{Question: record[0], Answer: record[1], Category: record[2]})
	}

	return cards, nil
}
