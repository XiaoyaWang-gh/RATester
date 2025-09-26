package hreflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetMethodIndexByName_1(t *testing.T) {
	type args struct {
		tp   reflect.Type
		name string
	}
	tests := []struct {
		name string
		args args
		want int
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

			if got := GetMethodIndexByName(tt.args.tp, tt.args.name); got != tt.want {
				t.Errorf("GetMethodIndexByName() = %v, want %v", got, tt.want)
			}
		})
	}
}
