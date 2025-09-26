package httplib

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestbuildCookieJar_1(t *testing.T) {
	tests := []struct {
		name string
		b    *BeegoHTTPRequest
		want http.CookieJar
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

			if got := tt.b.buildCookieJar(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BeegoHTTPRequest.buildCookieJar() = %v, want %v", got, tt.want)
			}
		})
	}
}
