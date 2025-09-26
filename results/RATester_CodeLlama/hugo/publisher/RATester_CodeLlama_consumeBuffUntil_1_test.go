package publisher

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestConsumeBuffUntil_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	w := &htmlElementsCollectorWriter{}
	condition := func() bool {
		return true
	}
	resolve := func(w *htmlElementsCollectorWriter) htmlCollectorStateFunc {
		return nil
	}
	// when
	s := w.consumeBuffUntil(condition, resolve)
	// then
	assert.NotNil(t, s)
}
