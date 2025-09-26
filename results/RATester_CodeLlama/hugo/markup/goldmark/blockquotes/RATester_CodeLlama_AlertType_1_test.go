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

	c := &blockquoteContext{}
	c.typ = "test"
	if c.AlertType() != "test" {
		t.Errorf("AlertType() = %v, want %v", c.AlertType(), "test")
	}
}
