package output

import (
	"fmt"
	"testing"
)

func TestIsZero_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := Format{}
	if !f.IsZero() {
		t.Errorf("Expected %v to be zero", f)
	}
}
