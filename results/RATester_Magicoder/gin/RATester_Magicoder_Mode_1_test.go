package gin

import (
	"fmt"
	"testing"
)

func TestMode_1(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Test Mode",
			want: "Test Mode",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Mode(); got != tt.want {
				t.Errorf("Mode() = %v, want %v", got, tt.want)
			}
		})
	}
}
