package main

import "testing"

func TestCorrect(t *testing.T) {
	card := Card{
		Question: "What is the capital of Alaska?",
		Answer:   "Juneau",
		Category: "Geography",
	}

	testCases := map[string]struct {
		guess string
		want  bool
	}{
		"returns true for correct guesses": {
			guess: "Juneau",
			want:  true,
		},
		"return false for incorrect guesses": {
			guess: "Fairbanks",
			want:  false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			turn := Turn{
				Guess: tc.guess,
				Card:  card,
			}

			got := turn.Correct()
			if got != tc.want {
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}
}

func TestFeedback(t *testing.T) {
	card := Card{
		Question: "Which planet is closest to the sun?",
		Answer:   "Mercury",
		Category: "STEM",
	}

	testCases := map[string]struct {
		guess string
		want  string
	}{
		"returns 'Correct!' for correct answers": {
			guess: "Mercury",
			want:  "Correct!",
		},
		"returns 'Incorrect.' for incorrect answers": {
			guess: "Saturn",
			want:  "Incorrect.",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			turn := Turn{tc.guess, card}

			got := turn.Feedback()
			if got != tc.want {
				t.Errorf("got %s want %s", got, tc.want)
			}
		})
	}
}
