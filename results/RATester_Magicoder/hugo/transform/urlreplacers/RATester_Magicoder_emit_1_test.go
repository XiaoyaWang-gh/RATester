package urlreplacers

import (
	"bytes"
	"fmt"
	"testing"
)

func Testemit_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &absurllexer{
		content: []byte("test content"),
		w:       &bytes.Buffer{},
		path:    []byte("test path"),
		pos:     5,
		start:   0,
	}

	l.emit()

	expected := "test content"
	actual := l.w.(*bytes.Buffer).String()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
