package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsMissing_1(t *testing.T) {
	type args struct {
		v reflect.Value
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

			if got := isMissing(tt.args.v); got != tt.want {
				t.Errorf("isMissing() = %v, want %v", got, tt.want)
			}
		})
	}
}
