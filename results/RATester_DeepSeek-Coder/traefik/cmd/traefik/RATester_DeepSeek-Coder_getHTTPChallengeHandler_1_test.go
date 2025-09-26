package main

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/provider/acme"
)

func TestGetHTTPChallengeHandler_1(t *testing.T) {
	type args struct {
		acmeProviders         []*acme.Provider
		httpChallengeProvider http.Handler
	}
	tests := []struct {
		name string
		args args
		want http.Handler
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

			if got := getHTTPChallengeHandler(tt.args.acmeProviders, tt.args.httpChallengeProvider); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getHTTPChallengeHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
