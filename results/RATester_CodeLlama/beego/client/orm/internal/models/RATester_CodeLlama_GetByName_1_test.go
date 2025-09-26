package models

import (
	"fmt"
	"testing"
)

func TestGetByName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &Fields{
		Fields: map[string]*FieldInfo{
			"name": {
				Name: "name",
			},
		},
	}
	name := "name"
	fi := f.GetByName(name)
	if fi == nil {
		t.Errorf("GetByName(%s) should not be nil", name)
	}
	if fi.Name != name {
		t.Errorf("GetByName(%s) should be %s", name, name)
	}
}
