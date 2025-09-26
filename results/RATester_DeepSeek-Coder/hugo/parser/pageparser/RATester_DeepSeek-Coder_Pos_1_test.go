package pageparser

import (
	"fmt"
	"testing"
)

func TestPos_1(t *testing.T) {
	testCases := []struct {
		name     string
		iterator *Iterator
		want     int
	}{
		{
			name: "Test case 1",
			iterator: &Iterator{
				items:   Items{},
				lastPos: 0,
			},
			want: 0,
		},
		{
			name: "Test case 2",
			iterator: &Iterator{
				items:   Items{},
				lastPos: 5,
			},
			want: 5,
		},
		{
			name: "Test case 3",
			iterator: &Iterator{
				items:   Items{},
				lastPos: 10,
			},
			want: 10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tc.iterator.Pos(); got != tc.want {
				t.Errorf("Pos() = %v, want %v", got, tc.want)
			}
		})
	}
}
