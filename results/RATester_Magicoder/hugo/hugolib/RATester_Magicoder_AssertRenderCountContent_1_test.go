package hugolib

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func TestAssertRenderCountContent_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &IntegrationTestBuilder{
		counters: &buildCounters{
			contentRenderCounter: atomic.Uint64{},
		},
	}

	b.counters.contentRenderCounter.Store(5)
	b.AssertRenderCountContent(5)

	b.counters.contentRenderCounter.Store(10)
	b.AssertRenderCountContent(10)
}
