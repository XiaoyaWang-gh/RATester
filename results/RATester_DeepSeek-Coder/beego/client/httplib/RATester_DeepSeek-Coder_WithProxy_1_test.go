package httplib

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestWithProxy_1(t *testing.T) {
	type args struct {
		proxy func(*http.Request) (*url.URL, error)
	}
	tests := []struct {
		name string
		args args
		want ClientOption
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

			if got := WithProxy(tt.args.proxy); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithProxy() = %v, want %v", got, tt.want)
			}
		})
	}
}
