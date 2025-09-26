package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func Testconverter_1(t *testing.T) {
	type args struct {
		t reflect.Type
	}
	tests := []struct {
		name string
		c    *cache
		args args
		want Converter
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

			if got := tt.c.converter(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("converter() = %v, want %v", got, tt.want)
			}
		})
	}
}
