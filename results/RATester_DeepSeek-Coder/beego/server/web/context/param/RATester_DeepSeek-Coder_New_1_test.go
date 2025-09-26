package param

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew_1(t *testing.T) {
	type args struct {
		name string
		opts []MethodParamOption
	}
	tests := []struct {
		name string
		args args
		want *MethodParam
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

			if got := New(tt.args.name, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
