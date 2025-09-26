package kinds

import (
	"fmt"
	"testing"
)

func TestIsBranch_1(t *testing.T) {
	tests := []struct {
		name string
		kind string
		want bool
	}{
		{"Home", KindHome, true},
		{"Section", "section", true},
		{"Taxonomy", "taxonomy", true},
		{"Term", "term", true},
		{"Other", "other", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := IsBranch(tt.kind); got != tt.want {
				t.Errorf("IsBranch() = %v, want %v", got, tt.want)
			}
		})
	}
}
