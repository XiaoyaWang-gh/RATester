package web

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestErrorHandler_1(t *testing.T) {
	type args struct {
		code string
		h    http.HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want *HttpServer
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

			if got := ErrorHandler(tt.args.code, tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ErrorHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
