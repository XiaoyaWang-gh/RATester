package migration

import (
	"fmt"
	"testing"
)

func TestUp_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &Migration{}
	m.ModifyType = "reverse"
	m.Up()
	if m.ModifyType != "alter" {
		t.Errorf("Expected m.ModifyType to be 'alter', got %s", m.ModifyType)
	}
}
