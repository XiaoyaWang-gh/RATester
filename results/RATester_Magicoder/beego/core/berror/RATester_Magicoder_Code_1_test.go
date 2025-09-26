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
		code:   123,
		module: "testModule",
		desc:   "testDesc",
		name:   "testName",
	}

	if cd.Code() != 123 {
		t.Errorf("Expected Code() to return 123, but got %d", cd.Code())
	}

	if cd.Module() != "testModule" {
		t.Errorf("Expected Module() to return 'testModule', but got '%s'", cd.Module())
	}

	if cd.Desc() != "testDesc" {
		t.Errorf("Expected Desc() to return 'testDesc', but got '%s'", cd.Desc())
	}

	if cd.Name() != "testName" {
		t.Errorf("Expected Name() to return 'testName', but got '%s'", cd.Name())
	}
}
