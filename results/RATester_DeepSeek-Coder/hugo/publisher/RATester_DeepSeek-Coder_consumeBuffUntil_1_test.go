package publisher

import (
	"bytes"
	"fmt"
	"testing"
)

func TestConsumeBuffUntil_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &htmlElementsCollectorWriter{
		buff: bytes.Buffer{},
	}

	condition := func() bool {
		return w.buff.Len() > 5
	}

	resolve := func(w *htmlElementsCollectorWriter) htmlCollectorStateFunc {
		return nil
	}

	w.consumeBuffUntil(condition, resolve)

	if w.buff.Len() != 0 {
		t.Errorf("Expected buffer length to be 0, got %d", w.buff.Len())
	}

	w.buff.WriteString("1234567")

	nextState := w.consumeBuffUntil(condition, resolve)

	if w.buff.Len() != 0 {
		t.Errorf("Expected buffer length to be 0, got %d", w.buff.Len())
	}

	if nextState == nil {
		t.Errorf("Expected next state to be not nil")
	}
}
