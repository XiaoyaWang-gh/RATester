package migration

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetOldUnsigned_1(t *testing.T) {
	type fields struct {
		OldName     string
		OldNull     string
		OldDefault  string
		OldUnsign   string
		OldDataType string
		NewName     string
		Column      Column
	}
	type args struct {
		unsign bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *RenameColumn
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

			c := &RenameColumn{
				OldName:     tt.fields.OldName,
				OldNull:     tt.fields.OldNull,
				OldDefault:  tt.fields.OldDefault,
				OldUnsign:   tt.fields.OldUnsign,
				OldDataType: tt.fields.OldDataType,
				NewName:     tt.fields.NewName,
				Column:      tt.fields.Column,
			}
			if got := c.SetOldUnsigned(tt.args.unsign); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetOldUnsigned() = %v, want %v", got, tt.want)
			}
		})
	}
}
