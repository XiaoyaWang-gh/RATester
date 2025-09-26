package httplib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewBeegoRequest_1(t *testing.T) {
	type args struct {
		rawurl string
		method string
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

			if got := NewBeegoRequest(tt.args.rawurl, tt.args.method); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBeegoRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
