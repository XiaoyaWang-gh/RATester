package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestOption_2(t *testing.T) {
	type args struct {
		opt []string
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

			if got := tt.t.Option(tt.args.opt...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Template.Option() = %v, want %v", got, tt.want)
			}
		})
	}
}
