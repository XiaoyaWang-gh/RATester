package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAnd_2(t *testing.T) {
	type args struct {
		arg0 reflect.Value
		args []reflect.Value
	}
	tests := []struct {
		name string
		args args
		want reflect.Value
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

			if got := and(tt.args.arg0, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("and() = %v, want %v", got, tt.want)
			}
		})
	}
}
