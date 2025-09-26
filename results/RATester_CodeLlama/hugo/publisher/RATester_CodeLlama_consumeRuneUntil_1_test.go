package publisher

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestConsumeRuneUntil_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	w := &htmlElementsCollectorWriter{}
	condition := func(r rune) bool {
		return r == 'a'
	}
	resolve := func(w *htmlElementsCollectorWriter) htmlCollectorStateFunc {
		return nil
	}
	// when
	s := w.consumeRuneUntil(condition, resolve)
	// then
	assert.NotNil(t, s)
}
