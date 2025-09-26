package publisher

import (
	"fmt"
	"reflect"
	"testing"
)

func TestExtractSingleQuotedStrings_1(t *testing.T) {
	type test struct {
		input string
		want  []string
	}

	tests := []test{
		{input: "", want: []string{}},
		{input: "hello world", want: []string{"hello", "world"}},
		{input: "hello 'world'", want: []string{"world"}},
		{input: "hello 'world' 'again'", want: []string{"world", "again"}},
		{input: "hello 'world ' 'again'", want: []string{"world ", "again"}},
		{input: "hello 'world ' ' again'", want: []string{"world ", "again"}},
		{input: "hello 'world ' ' again '", want: []string{"world ", "again "}},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := extractSingleQuotedStrings(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractSingleQuotedStrings(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
