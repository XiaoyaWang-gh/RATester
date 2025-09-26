package context

import (
	"fmt"
	"testing"
)

func TestIsClientError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	output := &BeegoOutput{}
	output.Status = 400
	if !output.IsClientError() {
		t.Errorf("IsClientError() = %v, want %v", output.IsClientError(), true)
	}
}
