package berror

import (
	"fmt"
	"testing"
)

func TestCode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cd := &codeDefinition{
		code:   12345,
		module: "testModule",
		desc:   "testDesc",
		name:   "testName",
	}

	if cd.Code() != 12345 {
		t.Errorf("Expected code to be 12345, got %d", cd.Code())
	}
}
