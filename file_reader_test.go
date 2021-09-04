package main

import (
	"errors"
	"os"
	"testing"
)

func TestCreateCardsFromFile(t *testing.T) {
	os.WriteFile("tmp.csv", []byte("a,b,c\nd,e"), 0644)
	defer os.Remove("tmp.csv")

	testCases := map[string]struct {
		filepath string
		expCards bool
		expError error
	}{
		"success": {
			filepath: "cards.csv",
			expCards: true,
		},
		"error opening file": {
			filepath: "notarealfile",
			expError: errors.New("open notarealfile: no such file or directory"),
		},
		"error reading file": {
			filepath: "tmp.csv",
			expError: errors.New("record on line 2: wrong number of fields"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			cards, err := createCardsFromFile(tc.filepath)
			if (tc.expCards && len(cards) == 0) || (!tc.expCards && len(cards) > 0) {
				t.Errorf("got unexpected result for cards: %v", cards)
			}

			if tc.expError == nil && err != nil {
				t.Errorf("got %v want %v", err, tc.expError)
			}

			if tc.expError != nil && err.Error() != tc.expError.Error() {
				t.Errorf("got %v want %v", err, tc.expError)
			}
		})
	}
}
