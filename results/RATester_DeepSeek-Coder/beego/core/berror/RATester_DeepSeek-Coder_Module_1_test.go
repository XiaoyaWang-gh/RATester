package berror

import (
	"fmt"
	"testing"
)

func TestModule_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cd := &codeDefinition{
		code:   1234,
		module: "testModule",
		desc:   "testDesc",
		name:   "testName",
	}

	expected := "testModule"
	result := cd.Module()

	if result != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, result)
	}
}
