package auth

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestWriteHeader_2(t *testing.T) {
	type args struct {
		req                *http.Request
		forwardReq         *http.Request
		trustForwardHeader bool
		allowedHeaders     []string
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

			writeHeader(tt.args.req, tt.args.forwardReq, tt.args.trustForwardHeader, tt.args.allowedHeaders)
			if !reflect.DeepEqual(tt.args.forwardReq, tt.want) {
				t.Errorf("writeHeader() = %v, want %v", tt.args.forwardReq, tt.want)
			}
		})
	}
}
