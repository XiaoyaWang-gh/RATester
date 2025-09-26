package dashboard

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServeHTTP_23(t *testing.T) {
	tests := []struct {
		name    string
		handler Handler
		request *http.Request
		want    http.ResponseWriter
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

			tt.handler.ServeHTTP(tt.want, tt.request)
		})
	}
}
