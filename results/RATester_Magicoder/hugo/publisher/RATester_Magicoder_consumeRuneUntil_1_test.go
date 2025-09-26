package publisher

import (
	"fmt"
	"testing"
)

func TestconsumeRuneUntil_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &htmlElementsCollectorWriter{
		r: 'a',
		state: func(w *htmlElementsCollectorWriter) htmlCollectorStateFunc {
			return w.consumeRuneUntil(func(r rune) bool {
				return r == 'b'
			}, func(w *htmlElementsCollectorWriter) htmlCollectorStateFunc {
				return nil
			})
		},
	}
	w.state = w.state(w)
	if w.state == nil {
		t.Error("Expected state to be not nil")
	}
}
