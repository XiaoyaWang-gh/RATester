package middleware

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMethodFromHeader_1(t *testing.T) {
	type args struct {
		header string
	}
	tests := []struct {
		name string
		args args
		want MethodOverrideGetter
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

			if got := MethodFromHeader(tt.args.header); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MethodFromHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}
