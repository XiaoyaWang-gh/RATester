package limiter

import (
	"fmt"
	"testing"
)

func TestMsgsize_1(t *testing.T) {
	testCases := []struct {
		name string
		item item
		want int
	}{
		{
			name: "Test case 1",
			item: item{
				currHits: 1,
				prevHits: 2,
				exp:      3,
			},
			want: 35,
		},
		{
			name: "Test case 2",
			item: item{
				currHits: 10,
				prevHits: 20,
				exp:      30,
			},
			want: 43,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tc.item.Msgsize()
			if got != tc.want {
				t.Errorf("Expected %d, got %d", tc.want, got)
			}
		})
	}
}
