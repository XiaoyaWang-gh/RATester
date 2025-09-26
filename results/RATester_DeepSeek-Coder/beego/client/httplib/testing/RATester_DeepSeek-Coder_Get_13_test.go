package testing

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGet_13(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want *TestHTTPRequest
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

			if got := Get(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
