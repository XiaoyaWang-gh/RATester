package echo

import (
	"fmt"
	"testing"
)

func TestRouter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	router := e.Router()

	if router == nil {
		t.Error("Expected a router, but got nil")
	}
}
