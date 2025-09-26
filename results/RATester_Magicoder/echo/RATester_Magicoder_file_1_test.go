package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFile_1(t *testing.T) {
	type args struct {
		path string
		file string
		get  func(string, HandlerFunc, ...MiddlewareFunc) *Route
		m    []MiddlewareFunc
	}
	tests := []struct {
		name string
		args args
		want *Route
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

			c := common{}
			if got := c.file(tt.args.path, tt.args.file, tt.args.get, tt.args.m...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("file() = %v, want %v", got, tt.want)
			}
		})
	}
}
