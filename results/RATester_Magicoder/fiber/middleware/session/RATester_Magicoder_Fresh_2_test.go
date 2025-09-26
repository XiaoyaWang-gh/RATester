package session

import (
	"fmt"
	"testing"
)

func TestFresh_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Session{
		fresh: true,
	}

	if !s.Fresh() {
		t.Error("Expected Fresh() to return true")
	}

	s.fresh = false

	if s.Fresh() {
		t.Error("Expected Fresh() to return false")
	}
}
