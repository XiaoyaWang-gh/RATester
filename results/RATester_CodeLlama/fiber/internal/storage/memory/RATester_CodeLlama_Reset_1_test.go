package memory

import (
	"fmt"
	"testing"
)

func TestReset_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Storage{}
	s.Reset()
	if len(s.db) != 0 {
		t.Errorf("expected empty db, got %v", s.db)
	}
}
