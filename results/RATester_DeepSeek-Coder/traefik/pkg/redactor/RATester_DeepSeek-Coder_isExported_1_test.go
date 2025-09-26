package redactor

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsExported_1(t *testing.T) {
	type args struct {
		f reflect.StructField
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

			if got := isExported(tt.args.f); got != tt.want {
				t.Errorf("isExported() = %v, want %v", got, tt.want)
			}
		})
	}
}
