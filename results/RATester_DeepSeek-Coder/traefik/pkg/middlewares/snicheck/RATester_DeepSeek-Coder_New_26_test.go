package snicheck

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNew_26(t *testing.T) {
	testCases := []struct {
		desc              string
		tlsOptionsForHost map[string]string
		next              http.Handler
		want              *SNICheck
	}{
		{
			desc:              "Test case 1",
			tlsOptionsForHost: map[string]string{"example.com": "tlsOptions1"},
			next:              http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}),
			want: &SNICheck{
				next:              http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}),
				tlsOptionsForHost: map[string]string{"example.com": "tlsOptions1"},
			},
		},
		// Add more test cases as needed
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := New(tC.tlsOptionsForHost, tC.next)
			if !reflect.DeepEqual(got, tC.want) {
				t.Errorf("New() = %v, want %v", got, tC.want)
			}
		})
	}
}
