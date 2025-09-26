package navigation

import (
	"fmt"
	"testing"
)

func TestKeyName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &MenuEntry{
		MenuConfig: MenuConfig{
			Identifier: "foo",
			Name:       "bar",
		},
	}

	if m.KeyName() != "foo" {
		t.Errorf("Expected %q, got %q", "foo", m.KeyName())
	}

	m.Identifier = ""

	if m.KeyName() != "bar" {
		t.Errorf("Expected %q, got %q", "bar", m.KeyName())
	}
}
