package migration

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetOldDefault_1(t *testing.T) {
	type args struct {
		def string
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

			if got := tt.c.SetOldDefault(tt.args.def); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RenameColumn.SetOldDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
