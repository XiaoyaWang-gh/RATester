package context

import (
	"fmt"
	"testing"
)

func TestIsOk_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	output := &BeegoOutput{
		Status: 200,
	}
	if !output.IsOk() {
		t.Errorf("IsOk() = %v, want %v", output.IsOk(), true)
	}
}
