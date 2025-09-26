package urlreplacers

import (
	"fmt"
	"testing"
)

func TestPosAfterURL_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &absurllexer{
		content: []byte("http://example.com"),
		pos:     0,
	}

	result := l.posAfterURL([]byte(">"))
	expected := 1
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
