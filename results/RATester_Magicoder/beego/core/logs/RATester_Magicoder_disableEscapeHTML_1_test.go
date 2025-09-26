package logs

import (
	"fmt"
	"testing"
)

func TestdisableEscapeHTML_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type testStruct struct {
		escapeHTML bool
	}

	test := testStruct{escapeHTML: true}
	disableEscapeHTML(test)

	if test.escapeHTML {
		t.Error("Expected escapeHTML to be false, but it was true")
	}
}
