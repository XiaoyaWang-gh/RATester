package template

import (
	"fmt"
	"testing"
)

func TestErrRecover_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var err error
	defer errRecover(&err)
	panic("generic error")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
