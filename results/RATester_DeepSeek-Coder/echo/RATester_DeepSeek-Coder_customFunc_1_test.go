package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCustomFunc_1(t *testing.T) {
	type args struct {
		sourceParam    string
		customFunc     func(values []string) []error
		valueMustExist bool
	}
	tests := []struct {
		name string
		b    *ValueBinder
		args args
		want *ValueBinder
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

			if got := tt.b.customFunc(tt.args.sourceParam, tt.args.customFunc, tt.args.valueMustExist); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValueBinder.customFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
