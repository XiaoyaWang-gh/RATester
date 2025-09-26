package context

import (
	"fmt"
	"testing"
)

func TestData_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{}
	data := input.Data()
	if data == nil {
		t.Error("Expected data to be initialized, but it was nil")
	}
}
