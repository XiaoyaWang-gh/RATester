package middleware

import (
	"fmt"
	"testing"
)

func TestIgnore_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testWriter := &ignorableWriter{
		ignoreWrites: false,
	}

	testWriter.Ignore(true)

	if testWriter.ignoreWrites != true {
		t.Errorf("Expected ignoreWrites to be true, got %v", testWriter.ignoreWrites)
	}

	testWriter.Ignore(false)

	if testWriter.ignoreWrites != false {
		t.Errorf("Expected ignoreWrites to be false, got %v", testWriter.ignoreWrites)
	}
}
