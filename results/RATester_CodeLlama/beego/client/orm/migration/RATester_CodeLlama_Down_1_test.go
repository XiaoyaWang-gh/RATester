package migration

import (
	"fmt"
	"testing"
)

func TestDown_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &Migration{}
	m.ModifyType = "alter"
	m.Down()
	if m.ModifyType != "reverse" {
		t.Errorf("m.ModifyType = %v, want %v", m.ModifyType, "reverse")
	}
}
