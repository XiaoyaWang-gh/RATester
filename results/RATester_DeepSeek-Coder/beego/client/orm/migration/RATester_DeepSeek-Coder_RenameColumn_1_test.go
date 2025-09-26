package migration

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRenameColumn_1(t *testing.T) {
	type args struct {
		from string
		to   string
	}
	tests := []struct {
		name string
		m    *Migration
		args args
		want *RenameColumn
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

			if got := tt.m.RenameColumn(tt.args.from, tt.args.to); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Migration.RenameColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}
