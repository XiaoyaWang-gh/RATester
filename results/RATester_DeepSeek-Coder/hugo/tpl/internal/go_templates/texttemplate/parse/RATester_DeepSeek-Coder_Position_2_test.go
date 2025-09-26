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
		{"Test 1", 10, 10},
		{"Test 2", 20, 20},
		{"Test 3", 30, 30},
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
