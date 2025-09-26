package api

import (
	"fmt"
	"net/http"
	"testing"
)

func TestgetRouter_1(t *testing.T) {
	type args struct {
		rw      http.ResponseWriter
		request *http.Request
	}
	tests := []struct {
		name string
		h    Handler
		args args
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

			tt.h.getRouter(tt.args.rw, tt.args.request)
		})
	}
}
