package text

import (
	"fmt"
	"testing"
)

func TestRemoveAccentsString_1(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "Test with string with accents",
			args: "éèêëàäâ",
			want: "eeeeaaa",
		},
		{
			name: "Test with string without accents",
			args: "abc",
			want: "abc",
		},
		{
			name: "Test with empty string",
			args: "",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := RemoveAccentsString(tt.args); got != tt.want {
				t.Errorf("RemoveAccentsString() = %v, want %v", got, tt.want)
			}
		})
	}
}
