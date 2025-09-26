package context

import (
	"fmt"
	"testing"
)

func TestIsSuccessful_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	output := &BeegoOutput{}
	output.Status = 200
	if !output.IsSuccessful() {
		t.Errorf("IsSuccessful() = %v, want %v", output.IsSuccessful(), true)
	}
}
