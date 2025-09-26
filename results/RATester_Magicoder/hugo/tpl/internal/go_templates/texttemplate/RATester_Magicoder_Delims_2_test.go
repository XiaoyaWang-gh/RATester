package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDelims_2(t *testing.T) {
	type args struct {
		left  string
		right string
	}
	tests := []struct {
		name string
		t    *Template
		args args
		want *Template
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

			if got := tt.t.Delims(tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Template.Delims() = %v, want %v", got, tt.want)
			}
		})
	}
}
