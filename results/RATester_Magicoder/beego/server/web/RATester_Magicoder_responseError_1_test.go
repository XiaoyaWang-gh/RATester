package web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestresponseError_1(t *testing.T) {
	type args struct {
		rw         http.ResponseWriter
		r          *http.Request
		errCode    int
		errContent string
	}
	tests := []struct {
		name string
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

			responseError(tt.args.rw, tt.args.r, tt.args.errCode, tt.args.errContent)
		})
	}
}
