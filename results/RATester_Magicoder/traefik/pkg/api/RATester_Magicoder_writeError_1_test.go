package api

import (
	"fmt"
	"net/http"
	"testing"
)

func TestwriteError_1(t *testing.T) {
	type args struct {
		rw   http.ResponseWriter
		msg  string
		code int
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

			writeError(tt.args.rw, tt.args.msg, tt.args.code)
		})
	}
}
