package httplib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParam_1(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name string
		b    *BeegoHTTPRequest
		args args
		want *BeegoHTTPRequest
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

			if got := tt.b.Param(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Param() = %v, want %v", got, tt.want)
			}
		})
	}
}
