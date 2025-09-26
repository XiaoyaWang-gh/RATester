package images

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/disintegration/gift"
)

func TestColorBalance_1(t *testing.T) {
	filters := &Filters{}

	tests := []struct {
		name            string
		percentageRed   any
		percentageGreen any
		percentageBlue  any
		want            gift.Filter
	}{
		{
			name:            "Test case 1",
			percentageRed:   0.5,
			percentageGreen: 0.5,
			percentageBlue:  0.5,
			want:            gift.ColorBalance(0.5, 0.5, 0.5),
		},
		{
			name:            "Test case 2",
			percentageRed:   0.25,
			percentageGreen: 0.25,
			percentageBlue:  0.25,
			want:            gift.ColorBalance(0.25, 0.25, 0.25),
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := filters.ColorBalance(tt.percentageRed, tt.percentageGreen, tt.percentageBlue)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ColorBalance() = %v, want %v", got, tt.want)
			}
		})
	}
}
