package math

import (
	"fmt"
	"testing"
)

func TestRand_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}
	result := ns.Rand()
	if result < 0 || result > 1 {
		t.Errorf("Expected Rand() to return a value between 0 and 1, but got %f", result)
	}
}
