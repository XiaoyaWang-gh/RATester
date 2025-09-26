package resources

import (
	"fmt"
	"testing"
)

func TestHash_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given: A test context.
	l := &genericResource{}

	// when: The hash method is called.
	result := l.hash()

	// then: The result is 0.
	if result != 0 {
		t.Errorf("Expected result to be 0 but was %d", result)
	}
}
