package hydratation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMakeKeyName_1(t *testing.T) {
	tests := []struct {
		typ  reflect.Type
		want string
	}{
		{typ: reflect.TypeOf(1), want: "int"},
		{typ: reflect.TypeOf(1.0), want: "float64"},
		{typ: reflect.TypeOf(""), want: "string"},
		{typ: reflect.TypeOf([]int{}), want: "[]int"},
		{typ: reflect.TypeOf(map[string]int{}), want: "map[string]int"},
		{typ: reflect.TypeOf(struct{}{}), want: "struct{}"},
		{typ: reflect.TypeOf(func() {}), want: "func()"},
		{typ: reflect.TypeOf(new(int)), want: "int"},
		{typ: reflect.TypeOf(new(string)), want: "string"},
		{typ: reflect.TypeOf(new([]int)), want: "[]int"},
		{typ: reflect.TypeOf(new(map[string]int)), want: "map[string]int"},
		{typ: reflect.TypeOf(new(struct{})), want: "struct{}"},
		{typ: reflect.TypeOf(new(func())), want: "func()"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := makeKeyName(tt.typ); got != tt.want {
				t.Errorf("makeKeyName() = %v, want %v", got, tt.want)
			}
		})
	}
}
