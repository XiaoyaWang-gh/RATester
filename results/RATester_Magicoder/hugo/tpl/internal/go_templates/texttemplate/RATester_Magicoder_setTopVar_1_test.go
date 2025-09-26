package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestsetTopVar_1(t *testing.T) {
	type args struct {
		n     int
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

			tt.s.setTopVar(tt.args.n, tt.args.value)
		})
	}
}
