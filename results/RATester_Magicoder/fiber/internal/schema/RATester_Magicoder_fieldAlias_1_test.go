package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestfieldAlias_1(t *testing.T) {
	type args struct {
		field   reflect.StructField
		tagName string
	}
	tests := []struct {
		name      string
		args      args
		wantAlias string
		wantOpts  tagOptions
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

			gotAlias, gotOpts := fieldAlias(tt.args.field, tt.args.tagName)
			if gotAlias != tt.wantAlias {
				t.Errorf("fieldAlias() gotAlias = %v, want %v", gotAlias, tt.wantAlias)
			}
			if !reflect.DeepEqual(gotOpts, tt.wantOpts) {
				t.Errorf("fieldAlias() gotOpts = %v, want %v", gotOpts, tt.wantOpts)
			}
		})
	}
}
