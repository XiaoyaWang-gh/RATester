package models

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetFullName_1(t *testing.T) {
	type args struct {
		typ reflect.Type
	}
	tests := []struct {
		name string
		args args
		want string
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

			if got := GetFullName(tt.args.typ); got != tt.want {
				t.Errorf("GetFullName() = %v, want %v", got, tt.want)
			}
		})
	}
}
