package pageparser

import (
	"fmt"
	"testing"
)

func TesthasPrefix_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &pageLexer{
		input: []byte("Hello, world!"),
		pos:   0,
	}

	tests := []struct {
		prefix []byte
		want   bool
	}{
		{[]byte("Hello"), true},
		{[]byte("world"), false},
		{[]byte(""), true},
		{[]byte("Hello, world!"), true},
	}

	for _, test := range tests {
		got := l.hasPrefix(test.prefix)
		if got != test.want {
			t.Errorf("hasPrefix(%q) = %v, want %v", test.prefix, got, test.want)
		}
	}
}
