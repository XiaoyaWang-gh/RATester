package orm

import (
	"fmt"
	"testing"
)

func TestgetDefault_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ac := &_dbCache{
		cache: map[string]*alias{
			"default": {
				Name: "default",
				// other fields
			},
		},
	}

	al := ac.getDefault()

	if al == nil {
		t.Error("Expected alias, got nil")
	}

	if al.Name != "default" {
		t.Errorf("Expected alias name to be 'default', got '%s'", al.Name)
	}
}
