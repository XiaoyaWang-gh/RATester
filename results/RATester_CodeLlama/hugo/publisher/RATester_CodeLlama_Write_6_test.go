package publisher

import (
	"fmt"
	"testing"
)

func TestWrite_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &htmlElementsCollectorWriter{}
	w.collector = &htmlElementsCollector{}
	w.state = w.defaultLexElementInside
	w.input = []byte("hello")
	w.pos = 0
	w.err = nil
	w.inQuote = 0
	w.buff.Reset()

	n, err := w.Write(w.input)

	if err != nil {
		t.Errorf("Write() error = %v, wantErr %v", err, nil)
		return
	}

	if n != len(w.input) {
		t.Errorf("Write() n = %v, want %v", n, len(w.input))
	}

	if w.pos != len(w.input) {
		t.Errorf("Write() w.pos = %v, want %v", w.pos, len(w.input))
	}

	if w.err != nil {
		t.Errorf("Write() w.err = %v, want %v", w.err, nil)
	}

	if w.inQuote != 0 {
		t.Errorf("Write() w.inQuote = %v, want %v", w.inQuote, 0)
	}

	if w.buff.String() != "hello" {
		t.Errorf("Write() w.buff.String() = %v, want %v", w.buff.String(), "hello")
	}
}
