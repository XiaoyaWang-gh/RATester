package metrics

import (
	"fmt"
	"testing"

	"github.com/go-kit/kit/metrics"
)

func TestNewCounterWithNoopHeaders_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var counter metrics.Counter
	counterWithNoopHeaders := NewCounterWithNoopHeaders(counter)
	if counterWithNoopHeaders.counter != counter {
		t.Errorf("counterWithNoopHeaders.counter = %v, want %v", counterWithNoopHeaders.counter, counter)
	}
}
