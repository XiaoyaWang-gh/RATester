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
		{"Test Home", KindHome, true},
		{"Test Section", "section", false},
		{"Test Taxonomy", "taxonomy", false},
		{"Test Term", "term", false},
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
