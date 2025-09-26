package page_generate

import (
	"fmt"
	"testing"
)

func TestImportsString_1(t *testing.T) {
	type test struct {
		input []string
		want  string
	}

	tests := []test{
		{input: []string{}, want: ""},
		{input: []string{"fmt"}, want: `import "fmt"`},
		{input: []string{"fmt", "os"}, want: `import (
"fmt"
"os"
)`},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("input=%v", tt.input), func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := importsString(tt.input)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
