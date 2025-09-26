package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMust_4(t *testing.T) {
	type args struct {
		t   *Template
		err error
	}
	tests := []struct {
		name string
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

			if got := Must(tt.args.t, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Must() = %v, want %v", got, tt.want)
			}
		})
	}
}
