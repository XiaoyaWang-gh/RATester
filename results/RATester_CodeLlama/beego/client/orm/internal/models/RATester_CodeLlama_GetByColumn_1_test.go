package models

import (
	"fmt"
	"testing"
)

func TestGetByColumn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &Fields{}
	f.Columns = map[string]*FieldInfo{
		"id": {
			Name: "id",
		},
	}
	fi := f.GetByColumn("id")
	if fi == nil {
		t.Errorf("GetByColumn failed")
	}
	if fi.Name != "id" {
		t.Errorf("GetByColumn failed")
	}
}
