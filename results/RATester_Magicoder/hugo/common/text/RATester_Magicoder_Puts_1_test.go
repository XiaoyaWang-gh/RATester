package text

import (
	"fmt"
	"testing"
)

func TestPuts_1(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "empty string",
			s:    "",
			want: "",
		},
		{
			name: "string with newline",
			s:    "Hello\n",
			want: "Hello\n",
		},
		{
			name: "string without newline",
			s:    "Hello",
			want: "Hello\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Puts(tt.s); got != tt.want {
				t.Errorf("Puts() = %v, want %v", got, tt.want)
			}
		})
	}
}
