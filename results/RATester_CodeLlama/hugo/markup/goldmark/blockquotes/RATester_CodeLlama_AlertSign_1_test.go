package blockquotes

import (
	"fmt"
	"testing"
)

func TestAlertSign_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &blockquoteContext{}
	c.alert.sign = "test"
	if c.AlertSign() != "test" {
		t.Errorf("AlertSign() = %v, want %v", c.AlertSign(), "test")
	}
}
