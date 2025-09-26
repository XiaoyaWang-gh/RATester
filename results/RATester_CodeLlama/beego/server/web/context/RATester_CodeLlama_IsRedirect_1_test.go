package context

import (
	"fmt"
	"testing"
)

func TestIsRedirect_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	output := &BeegoOutput{}
	output.Status = 301
	if !output.IsRedirect() {
		t.Errorf("IsRedirect() = %v, want %v", output.IsRedirect(), true)
	}
}
