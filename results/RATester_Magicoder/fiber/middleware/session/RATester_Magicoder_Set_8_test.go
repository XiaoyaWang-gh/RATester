package session

import (
	"fmt"
	"testing"
)

func TestSet_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Session{
		data: &data{},
	}

	key := "testKey"
	val := "testValue"

	s.Set(key, val)

	got := s.data.Get(key)
	if got != val {
		t.Errorf("Expected %v, got %v", val, got)
	}
}
