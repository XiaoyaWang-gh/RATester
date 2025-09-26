package accesslog

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func TestNextRequestCount_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	initialValue := atomic.LoadUint64(&requestCounter)
	nextValue := nextRequestCount()
	finalValue := atomic.LoadUint64(&requestCounter)

	if finalValue != initialValue+1 {
		t.Errorf("Expected %v, got %v", initialValue+1, finalValue)
	}

	if nextValue != initialValue+1 {
		t.Errorf("Expected %v, got %v", initialValue+1, nextValue)
	}
}
