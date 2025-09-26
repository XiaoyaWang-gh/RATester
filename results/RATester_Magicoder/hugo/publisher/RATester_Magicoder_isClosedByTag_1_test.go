package publisher

import (
	"fmt"
	"testing"
)

func TestisClosedByTag_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tests := []struct {
		b        []byte
		tagName  []byte
		expected bool
	}{
		{[]byte("</div>"), []byte("div"), true},
		{[]byte("</div>"), []byte("span"), false},
		{[]byte("<div>"), []byte("div"), false},
		{[]byte("<div/>"), []byte("div"), true},
		{[]byte("<div>Hello</div>"), []byte("div"), true},
		{[]byte("<div>Hello</span>"), []byte("div"), false},
		{[]byte("<div>Hello</div>"), []byte("span"), false},
		{[]byte("<div>Hello</span>"), []byte("span"), false},
		{[]byte("<div>Hello</div>"), []byte(""), false},
		{[]byte("<div>Hello</div>"), nil, false},
		{nil, []byte("div"), false},
		{nil, nil, false},
	}

	for _, test := range tests {
		result := isClosedByTag(test.b, test.tagName)
		if result != test.expected {
			t.Errorf("isClosedByTag(%v, %v) = %v; want %v", test.b, test.tagName, result, test.expected)
		}
	}
}
