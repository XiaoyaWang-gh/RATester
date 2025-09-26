package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/resource"
)

func TestgetParamToLower_2(t *testing.T) {
	type args struct {
		m   resource.ResourceParamsProvider
		key string
	}
	tests := []struct {
		name string
		args args
		want any
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

			if got := getParamToLower(tt.args.m, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getParamToLower() = %v, want %v", got, tt.want)
			}
		})
	}
}
