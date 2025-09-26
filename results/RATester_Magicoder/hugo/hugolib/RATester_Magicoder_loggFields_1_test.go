package hugolib

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func TestloggFields_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	counter := &buildCounters{
		contentRenderCounter: atomic.Uint64{},
		pageRenderCounter:    atomic.Uint64{},
	}

	counter.contentRenderCounter.Store(1)
	counter.pageRenderCounter.Store(2)

	fields := counter.loggFields()

	if len(fields) != 2 {
		t.Errorf("Expected 2 fields, got %d", len(fields))
	}

	if fields[0].Name != "pages" || fields[0].Value != uint64(2) {
		t.Errorf("Expected pages field with value 2, got %s: %v", fields[0].Name, fields[0].Value)
	}

	if fields[1].Name != "content" || fields[1].Value != uint64(1) {
		t.Errorf("Expected content field with value 1, got %s: %v", fields[1].Name, fields[1].Value)
	}
}
