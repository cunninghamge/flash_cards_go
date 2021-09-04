package reader

import (
	"errors"
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	os.WriteFile("tmp.csv", []byte("a,b,c\nd,e"), 0644)
	defer os.Remove("tmp.csv")

	testCases := map[string]struct {
		filepath   string
		expRecords bool
		expError   error
	}{
		"success": {
			filepath:   "../fixtures/cards.csv",
			expRecords: true,
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
			records, err := ReadFile(tc.filepath)
			if (tc.expRecords && len(records) == 0) || (!tc.expRecords && len(records) > 0) {
				t.Errorf("got unexpected result for records: %v", records)
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
