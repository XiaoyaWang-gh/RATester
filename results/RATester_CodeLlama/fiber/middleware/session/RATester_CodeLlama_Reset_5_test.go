package session

import (
	"fmt"
	"testing"
)

func TestReset_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Store{}
	err := s.Reset()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
