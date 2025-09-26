package publisher

import (
	"fmt"
	"testing"
)

func TestBackup_3(t *testing.T) {
	type testCase struct {
		name string
		l    *htmlElementsCollectorWriter
		want *htmlElementsCollectorWriter
	}

	testCases := []testCase{
		{
			name: "Test backup",
			l: &htmlElementsCollectorWriter{
				pos:   5,
				input: []byte("Hello"),
			},
			want: &htmlElementsCollectorWriter{
				pos:   4,
				input: []byte("Hello"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tc.l.backup()
			if tc.l.pos != tc.want.pos {
				t.Errorf("Expected position %d, got %d", tc.want.pos, tc.l.pos)
			}
		})
	}
}
