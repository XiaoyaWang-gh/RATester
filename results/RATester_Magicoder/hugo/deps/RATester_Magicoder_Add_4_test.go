package deps

import (
	"fmt"
	"testing"
)

func TestAdd_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &Listeners{}

	called := false
	l.Add(func() {
		called = true
	})

	l.Notify()

	if !called {
		t.Error("Expected listener to be called")
	}
}
