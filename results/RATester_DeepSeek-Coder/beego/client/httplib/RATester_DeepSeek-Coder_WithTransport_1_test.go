package httplib

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestWithTransport_1(t *testing.T) {
	type args struct {
		transport http.RoundTripper
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

			if got := WithTransport(tt.args.transport); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithTransport() = %v, want %v", got, tt.want)
			}
		})
	}
}
