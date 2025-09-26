package session

import (
	"fmt"
	"testing"
)

func TestID_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Session{
		id: "test_id",
	}

	id := s.ID()

	if id != "test_id" {
		t.Errorf("Expected ID to be 'test_id', but got %s", id)
	}
}
