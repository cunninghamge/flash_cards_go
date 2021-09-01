package main

type Turn struct {
	Guess string
	Card  Card
}

func (t Turn) Correct() bool {
	return t.Guess == t.Card.Answer
}

func (t Turn) Feedback() string {
	if t.Correct() {
		return "Correct!"
	}

	return "Incorrect."
}
