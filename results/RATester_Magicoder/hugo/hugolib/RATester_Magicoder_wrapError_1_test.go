package hugolib

import (
	"errors"
	"fmt"
	"testing"
)

func TestwrapError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Create a new instance of pageState
	ps := &pageState{
		// Initialize fields as needed
	}

	// Define the error to be passed to the method
	err := errors.New("test error")

	// Call the method under test
	result := ps.wrapError(err)

	// Assert that the result is as expected
	if result != err {
		t.Errorf("Expected %v, but got %v", err, result)
	}
}
