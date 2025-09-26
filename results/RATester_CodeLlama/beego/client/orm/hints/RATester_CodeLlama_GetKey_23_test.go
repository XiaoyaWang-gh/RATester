package hints

import (
	"fmt"
	"testing"
)

func TestGetKey_23(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Hint{key: "key"}
	if s.GetKey() != "key" {
		t.Errorf("GetKey() = %v, want %v", s.GetKey(), "key")
	}
}
