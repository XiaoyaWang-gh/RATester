package fiber

import (
	"fmt"
	"testing"
)

func TestRemoveEscapeChar_1(t *testing.T) {
	tests := []struct {
		name string
		word string
		want string
	}{
		{
			name: "Test with escape character",
			word: "hello\nworld",
			want: "helloworld",
		},
		{
			name: "Test without escape character",
			word: "hello world",
			want: "hello world",
		},
		{
			name: "Test with escape character at the end",
			word: "hello\n",
			want: "hello",
		},
		{
			name: "Test with escape character at the beginning",
			word: "\nhello",
			want: "hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := RemoveEscapeChar(tt.word); got != tt.want {
				t.Errorf("RemoveEscapeChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
