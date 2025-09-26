package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCanCompare_1(t *testing.T) {
	type args struct {
		v1 reflect.Value
		v2 reflect.Value
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

			if got := canCompare(tt.args.v1, tt.args.v2); got != tt.want {
				t.Errorf("canCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
