package redirect

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRawURL_1(t *testing.T) {
	tests := []struct {
		name string
		req  *http.Request
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

			if got := rawURL(tt.req); got != tt.want {
				t.Errorf("rawURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
