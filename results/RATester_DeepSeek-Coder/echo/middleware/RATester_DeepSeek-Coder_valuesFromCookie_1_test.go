package middleware

import (
	"fmt"
	"reflect"
	"testing"
)

func TestValuesFromCookie_1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want ValuesExtractor
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

			if got := valuesFromCookie(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("valuesFromCookie() = %v, want %v", got, tt.want)
			}
		})
	}
}
