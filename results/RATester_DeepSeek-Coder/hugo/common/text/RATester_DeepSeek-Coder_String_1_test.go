package text

import (
	"fmt"
	"testing"
)

func TestString_1(t *testing.T) {
	tests := []struct {
		name string
		pos  Position
		want string
	}{
		{
			name: "Test with valid position",
			pos:  Position{Filename: "test.go", Offset: 10, LineNumber: 5, ColumnNumber: 2},
			want: "test.go:5:2",
		},
		{
			name: "Test with no filename",
			pos:  Position{Offset: 10, LineNumber: 5, ColumnNumber: 2},
			want: "<stream>:5:2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.pos.String(); got != tt.want {
				t.Errorf("Position.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
