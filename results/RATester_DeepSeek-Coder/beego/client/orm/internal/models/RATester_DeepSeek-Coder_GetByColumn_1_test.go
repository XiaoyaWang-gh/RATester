package models

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetByColumn_1(t *testing.T) {
	fields := &Fields{
		Columns: map[string]*FieldInfo{
			"column1": {Name: "name1"},
			"column2": {Name: "name2"},
		},
	}

	tests := []struct {
		name     string
		fields   *Fields
		column   string
		expected *FieldInfo
	}{
		{
			name:     "TestGetByColumn_ExistingColumn",
			fields:   fields,
			column:   "column1",
			expected: &FieldInfo{Name: "name1"},
		},
		{
			name:     "TestGetByColumn_NonExistingColumn",
			fields:   fields,
			column:   "column3",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tt.fields.GetByColumn(tt.column)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Fields.GetByColumn() = %v, want %v", got, tt.expected)
			}
		})
	}
}
