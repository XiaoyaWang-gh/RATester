package context

import (
	"fmt"
	"testing"
)

func TestIsServerError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	output := &BeegoOutput{}
	output.Status = 500
	if !output.IsServerError() {
		t.Errorf("IsServerError() = %v, want %v", output.IsServerError(), true)
	}
}
