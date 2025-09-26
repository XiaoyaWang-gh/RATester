package herrors

import (
	"fmt"
	"testing"
)

func TestRecover_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Recover() = %v, want %v", r, nil)
		}
	}()

	Recover("test")
}
