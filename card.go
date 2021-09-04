package main

import "errors"

const errInvalidRecords = "invalid record: must contain exactly 3 fields"

type Card struct {
	Question string
	Answer   string
	Category string
}

func createCards(records [][]string) ([]Card, error) {
	if len(records[0]) != 3 {
		return nil, errors.New(errInvalidRecords)
	}

	var cards []Card
	for _, record := range records {
		cards = append(cards, Card{Question: record[0], Answer: record[1], Category: record[2]})
	}
	return cards, nil
}
