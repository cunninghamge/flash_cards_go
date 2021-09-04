package main

import "testing"

func TestBuildDeck(t *testing.T) {
	testCases := map[string]struct {
		records    [][]string
		deckLength int
		wantError  bool
	}{
		"success": {
			records: [][]string{
				{"question?", "answer.", "category"},
				{"question?", "answer.", "category"},
			},
			deckLength: 2,
		},
		"too few fields": {
			records: [][]string{
				{"question?", "answer."},
				{"question?", "answer."},
			},
			wantError: true,
		},
		"too many fields": {
			records: [][]string{
				{"question?", "answer.", "category", "extra field"},
				{"question?", "answer.", "category", "extra field"},
			},
			wantError: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			cards, err := createCards(tc.records)
			got := len(cards)
			if got != tc.deckLength {
				t.Errorf("created %d card but should have created %d", got, tc.deckLength)
			}

			if tc.wantError && err.Error() != errInvalidRecords {
				t.Errorf("got %v want %s", err, errInvalidRecords)
			}

			if !tc.wantError && err != nil {
				t.Errorf("got %v want %v", err, nil)
			}
		})
	}
}
