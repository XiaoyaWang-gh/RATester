package helpers

import (
	"bytes"
	"fmt"
	"testing"
)

func TestTable_1(t *testing.T) {
	testCases := []struct {
		name     string
		stats    *ProcessingStats
		expected string
	}{
		{
			name: "Test case 1",
			stats: &ProcessingStats{
				Name:            "Test",
				Pages:           10,
				PaginatorPages:  5,
				Static:          3,
				ProcessedImages: 2,
				Files:           1,
			},
			expected: "Pages:\t10\nPaginator Pages:\t5\nStatic:\t3\nProcessed Images:\t2\nFiles:\t1\n",
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			w := &bytes.Buffer{}
			tc.stats.Table(w)
			if got := w.String(); got != tc.expected {
				t.Errorf("Expected:\n%s\nGot:\n%s", tc.expected, got)
			}
		})
	}
}
