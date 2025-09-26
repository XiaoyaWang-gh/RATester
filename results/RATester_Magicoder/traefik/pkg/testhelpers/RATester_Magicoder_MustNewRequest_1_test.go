package testhelpers

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestMustNewRequest_1(t *testing.T) {
	type args struct {
		method string
		urlStr string
		body   io.Reader
	}
	tests := []struct {
		name string
		args args
		want *http.Request
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

			if got := MustNewRequest(tt.args.method, tt.args.urlStr, tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustNewRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
