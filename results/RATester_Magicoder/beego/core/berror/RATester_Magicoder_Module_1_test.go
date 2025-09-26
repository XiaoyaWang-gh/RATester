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
		module: "testModule",
	}

	module := cd.Module()

	if module != "testModule" {
		t.Errorf("Expected module to be 'testModule', but got '%s'", module)
	}
}
