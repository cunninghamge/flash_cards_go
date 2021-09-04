package main

import "strings"

type Turn struct {
	Guess string
	Card  Card
}

func (t Turn) Correct() bool {
	return strings.ToLower(t.Guess) == strings.ToLower(t.Card.Answer)
}

func (t Turn) Feedback() string {
	if t.Correct() {
		return "Correct!"
	}

	return "Incorrect."
}
