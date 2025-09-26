package pageparser

import (
	"fmt"
	"testing"
)

func TestLexShortcodeParamVal_1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Test 1",
			input: "[shortcode param=value]",
			want:  "value",
		},
		{
			name:  "Test 2",
			input: "[shortcode param=another value]",
			want:  "another value",
		},
		{
			name:  "Test 3",
			input: "[shortcode param=\"multi word value\"]",
			want:  "multi word value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			l := &pageLexer{
				input: []byte(tt.input),
			}
			l.state = lexShortcodeParamVal
			l.run()
			got := string(l.current())
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
