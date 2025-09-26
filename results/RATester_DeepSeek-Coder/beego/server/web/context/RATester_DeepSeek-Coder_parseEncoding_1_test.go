package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestParseEncoding_1(t *testing.T) {
	tests := []struct {
		name string
		r    *http.Request
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

			if got := parseEncoding(tt.r); got != tt.want {
				t.Errorf("parseEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}
