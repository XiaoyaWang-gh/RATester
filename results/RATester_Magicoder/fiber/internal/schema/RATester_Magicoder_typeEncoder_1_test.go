package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TesttypeEncoder_1(t *testing.T) {
	type args struct {
		t   reflect.Type
		reg map[reflect.Type]encoderFunc
	}
	tests := []struct {
		name string
		args args
		want encoderFunc
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

			if got := typeEncoder(tt.args.t, tt.args.reg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("typeEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}
