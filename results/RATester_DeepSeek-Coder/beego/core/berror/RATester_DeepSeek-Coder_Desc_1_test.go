package berror

import (
	"fmt"
	"testing"
)

func TestDesc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cd := &codeDefinition{
		code:   1,
		module: "module",
		desc:   "description",
		name:   "name",
	}

	expected := "description"
	actual := cd.Desc()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
