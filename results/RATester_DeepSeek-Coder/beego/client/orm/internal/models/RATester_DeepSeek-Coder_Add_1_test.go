package models

import (
	"fmt"
	"strings"
	"testing"
)

func TestAdd_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fi := &FieldInfo{
		Name: "test",
	}

	f := &Fields{
		Columns:   make(map[string]*FieldInfo),
		Fields:    make(map[string]*FieldInfo),
		FieldsLow: make(map[string]*FieldInfo),
	}

	added := f.Add(fi)

	if !added {
		t.Errorf("Expected true, got %v", added)
	}

	if f.Fields[fi.Name] != fi {
		t.Errorf("Expected %v, got %v", fi, f.Fields[fi.Name])
	}

	if f.Columns[fi.Name] != fi {
		t.Errorf("Expected %v, got %v", fi, f.Columns[fi.Name])
	}

	if f.FieldsLow[strings.ToLower(fi.Name)] != fi {
		t.Errorf("Expected %v, got %v", fi, f.FieldsLow[strings.ToLower(fi.Name)])
	}
}
