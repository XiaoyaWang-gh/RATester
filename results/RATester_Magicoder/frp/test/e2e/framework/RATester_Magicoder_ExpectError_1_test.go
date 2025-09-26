package framework

import (
	"fmt"
	"testing"
)

func TestExpectError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	reqExpect := &RequestExpect{}
	reqExpect.ExpectError(true)
	if reqExpect.expectError != true {
		t.Errorf("Expected ExpectError to be true, but got %v", reqExpect.expectError)
	}
}
