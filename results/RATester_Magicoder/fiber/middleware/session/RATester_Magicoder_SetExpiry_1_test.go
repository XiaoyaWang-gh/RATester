package session

import (
	"fmt"
	"testing"
	"time"
)

func TestSetExpiry_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Session{
		exp: 0,
	}

	exp := time.Duration(10)
	s.SetExpiry(exp)

	if s.exp != exp {
		t.Errorf("Expected expiry to be %v, but got %v", exp, s.exp)
	}
}
