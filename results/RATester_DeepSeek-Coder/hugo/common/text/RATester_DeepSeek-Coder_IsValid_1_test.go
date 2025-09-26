package text

import (
	"fmt"
	"testing"
)

func TestIsValid_1(t *testing.T) {
	tests := []struct {
		name string
		pos  Position
		want bool
	}{
		{
			name: "valid position",
			pos:  Position{LineNumber: 1},
			want: true,
		},
		{
			name: "invalid position",
			pos:  Position{LineNumber: 0},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.pos.IsValid(); got != tt.want {
				t.Errorf("Position.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
