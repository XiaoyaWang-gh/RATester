package urlreplacers

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEmit_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &absurllexer{
		content: []byte("hello world"),
		w:       &bytes.Buffer{},
		path:    []byte("."),
	}
	l.emit()
	if l.w.(*bytes.Buffer).String() != "hello world" {
		t.Errorf("expected %s, got %s", "hello world", l.w.(*bytes.Buffer).String())
	}
}
