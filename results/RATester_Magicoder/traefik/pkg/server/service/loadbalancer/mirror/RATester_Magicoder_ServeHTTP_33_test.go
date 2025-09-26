package mirror

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServeHTTP_33(t *testing.T) {
	tests := []struct {
		name string
		req  *http.Request
		want http.ResponseWriter
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

			m := &Mirroring{}
			m.ServeHTTP(tt.want, tt.req)
		})
	}
}
