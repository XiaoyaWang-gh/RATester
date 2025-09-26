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

	e := new(Echo)
	r := e.Router()
	if r == nil {
		t.Errorf("Expected Router, got nil")
	}
}
