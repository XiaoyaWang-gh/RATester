package blockquotes

import (
	"fmt"
	"testing"
)

func TestAlertType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &blockquoteContext{
		typ: "info",
	}

	expected := "info"
	actual := ctx.AlertType()

	if actual != expected {
		t.Errorf("Expected: %s, but got: %s", expected, actual)
	}
}
