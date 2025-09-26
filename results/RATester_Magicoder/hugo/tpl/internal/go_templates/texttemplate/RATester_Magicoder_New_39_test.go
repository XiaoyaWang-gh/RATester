package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew_39(t *testing.T) {
	type args struct {
		name string
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

			if got := tt.t.New(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Template.New() = %v, want %v", got, tt.want)
			}
		})
	}
}
