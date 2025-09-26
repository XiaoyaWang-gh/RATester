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

	w := &ignorableWriter{}
	w.Ignore(true)
	if w.ignoreWrites != true {
		t.Error("Expected ignoreWrites to be true")
	}
	w.Ignore(false)
	if w.ignoreWrites != false {
		t.Error("Expected ignoreWrites to be false")
	}
}
