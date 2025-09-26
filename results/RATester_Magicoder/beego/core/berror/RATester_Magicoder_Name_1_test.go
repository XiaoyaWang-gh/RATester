package berror

import (
	"fmt"
	"testing"
)

func TestName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cd := &codeDefinition{
		name: "testName",
	}

	name := cd.Name()
	if name != "testName" {
		t.Errorf("Expected 'testName', but got '%s'", name)
	}
}
