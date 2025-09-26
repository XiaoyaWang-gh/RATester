package models

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetByName_1(t *testing.T) {
	fields := &Fields{
		Fields: map[string]*FieldInfo{
			"field1": {Name: "field1"},
			"field2": {Name: "field2"},
		},
	}

	testCases := []struct {
		name     string
		field    string
		expected *FieldInfo
	}{
		{"ExistingField", "field1", &FieldInfo{Name: "field1"}},
		{"NonExistingField", "field3", nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := fields.GetByName(tc.field)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
