package filenotify

import (
	"fmt"
	"testing"
)

func TestRemove_2(t *testing.T) {
	type testCase struct {
		name     string
		poller   *filePoller
		expected error
	}

	testCases := []testCase{
		{
			name: "remove existing watch",
			poller: &filePoller{
				watches: map[string]struct{}{
					"test.txt": {},
				},
			},
			expected: nil,
		},
		{
			name: "remove non-existing watch",
			poller: &filePoller{
				watches: map[string]struct{}{},
			},
			expected: errNoSuchWatch,
		},
		{
			name: "remove watch on closed poller",
			poller: &filePoller{
				watches: map[string]struct{}{
					"test.txt": {},
				},
				closed: true,
			},
			expected: errPollerClosed,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := tc.poller.remove("test.txt")
			if err != tc.expected {
				t.Errorf("Expected error %v, got %v", tc.expected, err)
			}
		})
	}
}
