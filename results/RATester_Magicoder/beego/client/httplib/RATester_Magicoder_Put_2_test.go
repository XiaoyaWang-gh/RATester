package httplib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPut_2(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
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

			if got := Put(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Put() = %v, want %v", got, tt.want)
			}
		})
	}
}
