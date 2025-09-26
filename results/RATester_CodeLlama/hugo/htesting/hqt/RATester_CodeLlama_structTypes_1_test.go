package hqt

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStructTypes_1(t *testing.T) {
	type args struct {
		v reflect.Value
		m map[reflect.Type]struct{}
	}
	tests := []struct {
		name string
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

			structTypes(tt.args.v, tt.args.m)
		})
	}
}
