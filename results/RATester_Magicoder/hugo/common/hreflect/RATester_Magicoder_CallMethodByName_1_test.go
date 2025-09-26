package hreflect

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestCallMethodByName_1(t *testing.T) {
	type args struct {
		cxt  context.Context
		name string
		v    reflect.Value
	}
	tests := []struct {
		name string
		args args
		want []reflect.Value
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

			if got := CallMethodByName(tt.args.cxt, tt.args.name, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CallMethodByName() = %v, want %v", got, tt.want)
			}
		})
	}
}
