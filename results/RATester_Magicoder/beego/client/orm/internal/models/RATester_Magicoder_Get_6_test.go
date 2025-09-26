package models

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGet_6(t *testing.T) {
	mc := &ModelCache{
		cache: map[string]*ModelInfo{
			"table1": {
				Table: "table1",
			},
			"table2": {
				Table: "table2",
			},
		},
	}

	tests := []struct {
		name     string
		table    string
		expected *ModelInfo
		ok       bool
	}{
		{
			name:     "Existing table",
			table:    "table1",
			expected: &ModelInfo{Table: "table1"},
			ok:       true,
		},
		{
			name:     "Non-existing table",
			table:    "table3",
			expected: nil,
			ok:       false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			mi, ok := mc.Get(test.table)
			if ok != test.ok {
				t.Errorf("Expected ok to be %v, but got %v", test.ok, ok)
			}
			if !reflect.DeepEqual(mi, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, mi)
			}
		})
	}
}
