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

	s := &Session{}
	s.id = "1234567890"
	if s.ID() != "1234567890" {
		t.Errorf("ID() = %v, want %v", s.ID(), "1234567890")
	}
}
