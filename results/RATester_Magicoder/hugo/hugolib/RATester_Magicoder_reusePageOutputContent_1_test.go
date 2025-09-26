package hugolib

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func TestreusePageOutputContent_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Create a new instance of pageState
	ps := &pageState{
		pageOutputTemplateVariationsState: &atomic.Uint32{},
	}

	// Set the value of pageOutputTemplateVariationsState
	ps.pageOutputTemplateVariationsState.Store(1)

	// Call the method under test
	result := ps.reusePageOutputContent()

	// Assert that the result is true
	if !result {
		t.Errorf("Expected true, but got false")
	}
}
