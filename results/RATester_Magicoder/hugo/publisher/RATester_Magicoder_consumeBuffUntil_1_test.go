package publisher

import (
	"bytes"
	"fmt"
	"testing"
)

func TestconsumeBuffUntil_1(t *testing.T) {
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
}
