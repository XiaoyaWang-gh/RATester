package migration

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetOldNullable_1(t *testing.T) {
	type args struct {
		null bool
	}
	tests := []struct {
		name string
		c    *RenameColumn
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

			if got := tt.c.SetOldNullable(tt.args.null); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RenameColumn.SetOldNullable() = %v, want %v", got, tt.want)
			}
		})
	}
}
