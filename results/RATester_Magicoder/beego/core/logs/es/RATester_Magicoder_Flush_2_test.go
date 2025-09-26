package es

import (
	"fmt"
	"testing"
)

func TestFlush_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Create a new instance of esLogger
	el := &esLogger{}

	// Call the method under test
	el.Flush()

	// Add assertions to check the expected behavior
	// For example, if Flush() does not return an error, you can add:
	// assert.NoError(t, err)
}
