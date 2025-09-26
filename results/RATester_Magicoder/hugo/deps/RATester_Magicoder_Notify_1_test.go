package deps

import (
	"fmt"
	"testing"
)

func TestNotify_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &Listeners{
		listeners: make([]func(), 0),
	}

	called := false
	l.Add(func() {
		called = true
	})

	l.Notify()

	if !called {
		t.Error("Notify did not call the added function")
	}
}
