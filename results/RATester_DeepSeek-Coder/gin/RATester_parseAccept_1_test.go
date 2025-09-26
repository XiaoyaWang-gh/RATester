package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseAccept_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		input string
		want  []string
	}

	tests := []test{
		{"text/html, application/xhtml+xml, application/xml;q=0.9, */*;q=0.8", []string{"text/html", "application/xhtml+xml", "application/xml", "*/*"}},
		{"application/json", []string{"application/json"}},
		{"text/html;level=1, text/html;level=2;q=0.5, */*;q=0.1", []string{"text/html;level=1", "text/html;level=2", "*/*"}},
	}

	for _, test := range tests {
		got := parseAccept(test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("parseAccept(%q) = %q, want %q", test.input, got, test.want)
		}
	}
}
