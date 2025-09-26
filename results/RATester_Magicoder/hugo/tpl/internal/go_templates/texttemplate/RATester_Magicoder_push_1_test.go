package template

import (
	"fmt"
	"reflect"
	"testing"
)

func Testpush_1(t *testing.T) {
	type args struct {
		name  string
		value reflect.Value
	}
	tests := []struct {
		name string
		s    *state
		args args
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

			tt.s.push(tt.args.name, tt.args.value)
		})
	}
}
