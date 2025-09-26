package data

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAddDefaultHeaders_2(t *testing.T) {
	type args struct {
		req     *http.Request
		accepts []string
	}
	tests := []struct {
		name string
		args args
		want *http.Request
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

			addDefaultHeaders(tt.args.req, tt.args.accepts...)
			if !reflect.DeepEqual(tt.args.req, tt.want) {
				t.Errorf("addDefaultHeaders() = %v, want %v", tt.args.req, tt.want)
			}
		})
	}
}
