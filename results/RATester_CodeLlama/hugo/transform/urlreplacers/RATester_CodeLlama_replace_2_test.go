package urlreplacers

import (
	"fmt"
	"os"
	"testing"
)

func TestReplace_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &absurllexer{
		content: []byte("http://www.google.com/search?q=golang"),
		w:       os.Stdout,
	}
	l.replace()
	if l.pos != len(l.content) {
		t.Errorf("expected pos to be %d, got %d", len(l.content), l.pos)
	}
}
