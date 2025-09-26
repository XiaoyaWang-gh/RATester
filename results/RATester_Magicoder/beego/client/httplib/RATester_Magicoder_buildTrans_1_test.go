package httplib

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestbuildTrans_1(t *testing.T) {
	tests := []struct {
		name string
		b    *BeegoHTTPRequest
		want http.RoundTripper
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

			if got := tt.b.buildTrans(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BeegoHTTPRequest.buildTrans() = %v, want %v", got, tt.want)
			}
		})
	}
}
