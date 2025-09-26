package create

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRoundTrip_1(t *testing.T) {
	tests := []struct {
		name    string
		req     *http.Request
		wantErr bool
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

			tr := &transport{}
			_, err := tr.RoundTrip(tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("transport.RoundTrip() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
