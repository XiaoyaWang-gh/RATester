package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNot_1(t *testing.T) {
	type args struct {
		arg reflect.Value
	}
	tests := []struct {
		name string
		args args
		want bool
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

			if got := not(tt.args.arg); got != tt.want {
				t.Errorf("not() = %v, want %v", got, tt.want)
			}
		})
	}
}
