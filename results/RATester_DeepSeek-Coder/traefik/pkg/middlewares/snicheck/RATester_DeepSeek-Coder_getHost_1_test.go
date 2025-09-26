package snicheck

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetHost_1(t *testing.T) {
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name string
		args args
		want string
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

			if got := getHost(tt.args.req); got != tt.want {
				t.Errorf("getHost() = %v, want %v", got, tt.want)
			}
		})
	}
}
