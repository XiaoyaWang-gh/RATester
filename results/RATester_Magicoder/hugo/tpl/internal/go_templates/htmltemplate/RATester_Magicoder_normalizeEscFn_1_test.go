package template

import (
	"fmt"
	"testing"
)

func TestnormalizeEscFn_1(t *testing.T) {
	tests := []struct {
		name string
		e    string
		want string
	}{
		{
			name: "Test case 1",
			e:    "field name",
			want: "field want",
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

			if got := normalizeEscFn(tt.e); got != tt.want {
				t.Errorf("normalizeEscFn() = %v, want %v", got, tt.want)
			}
		})
	}
}
