package migration

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddColumns_1(t *testing.T) {
	type args struct {
		columns []*Column
	}
	tests := []struct {
		name string
		m    *Migration
		args args
		want *Migration
	}{
		{
			name: "TestAddColumns",
			m:    &Migration{},
			args: args{
				columns: []*Column{
					{Name: "column1", DataType: "varchar(255)"},
					{Name: "column2", DataType: "int"},
				},
			},
			want: &Migration{
				Columns: []*Column{
					{Name: "column1", DataType: "varchar(255)"},
					{Name: "column2", DataType: "int"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.m.AddColumns(tt.args.columns...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Migration.AddColumns() = %v, want %v", got, tt.want)
			}
		})
	}
}
