package validation

import (
	"fmt"
	"reflect"
	"testing"
)

func Testapply_1(t *testing.T) {
	type args struct {
		chk Validator
		obj interface{}
	}
	tests := []struct {
		name string
		v    *Validation
		args args
		want *Result
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

			if got := tt.v.apply(tt.args.chk, tt.args.obj); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Validation.apply() = %v, want %v", got, tt.want)
			}
		})
	}
}
