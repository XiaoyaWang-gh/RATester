package parse

import (
	"fmt"
	"testing"
)

func TestPosition_2(t *testing.T) {
	tests := []struct {
		name string
		p    Pos
		want Pos
	}{
		{
			name: "Test case 1",
			p:    10,
			want: 10,
		},
		{
			name: "Test case 2",
			p:    20,
			want: 20,
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

			if got := tt.p.Position(); got != tt.want {
				t.Errorf("Pos.Position() = %v, want %v", got, tt.want)
			}
		})
	}
}
