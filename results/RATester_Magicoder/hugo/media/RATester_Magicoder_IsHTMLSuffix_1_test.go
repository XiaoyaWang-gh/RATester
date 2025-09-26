package media

import (
	"fmt"
	"testing"
)

func TestIsHTMLSuffix_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	contentTypes := ContentTypes{
		HTML: Type{
			SuffixesCSV: "html,htm",
		},
	}

	tests := []struct {
		suffix string
		want   bool
	}{
		{"html", true},
		{"htm", true},
		{"txt", false},
	}

	for _, test := range tests {
		got := contentTypes.IsHTMLSuffix(test.suffix)
		if got != test.want {
			t.Errorf("IsHTMLSuffix(%q) = %v, want %v", test.suffix, got, test.want)
		}
	}
}
