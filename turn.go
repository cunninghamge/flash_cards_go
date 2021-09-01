package main

type Turn struct {
	Guess string
	Card  Card
}

func (t Turn) Correct() bool {
	return false
}

func (t Turn) Feedback() string {
	return ""
}
