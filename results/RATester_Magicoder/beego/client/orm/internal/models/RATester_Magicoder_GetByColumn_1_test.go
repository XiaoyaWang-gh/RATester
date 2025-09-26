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
		name   string
		column string
		want   *FieldInfo
	}{
		{"TestGetByColumn_ExistingColumn", "column1", &FieldInfo{Name: "name1"}},
		{"TestGetByColumn_NonExistingColumn", "column3", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := fields.GetByColumn(tt.column); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}
