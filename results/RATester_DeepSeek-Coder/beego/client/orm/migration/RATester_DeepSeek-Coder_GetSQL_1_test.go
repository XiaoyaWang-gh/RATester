package migration

import (
	"fmt"
	"testing"
)

func TestGetSQL_1(t *testing.T) {
	type fields struct {
		ModifyType string
		TableName  string
		Engine     string
		Charset    string
		Columns    []*Column
		Primary    []*Column
		Uniques    []*Unique
		Foreigns   []*Foreign
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			m := &Migration{
				ModifyType: tt.fields.ModifyType,
				TableName:  tt.fields.TableName,
				Engine:     tt.fields.Engine,
				Charset:    tt.fields.Charset,
				Columns:    tt.fields.Columns,
				Primary:    tt.fields.Primary,
				Uniques:    tt.fields.Uniques,
				Foreigns:   tt.fields.Foreigns,
			}
			if got := m.GetSQL(); got != tt.want {
				t.Errorf("Migration.GetSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}
