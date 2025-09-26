package context

import (
	"fmt"
	"testing"
)

func TestSetStatus_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	output := &BeegoOutput{}
	status := 200
	output.SetStatus(status)
	if output.Status != status {
		t.Errorf("Expected status %d, but got %d", status, output.Status)
	}
}
