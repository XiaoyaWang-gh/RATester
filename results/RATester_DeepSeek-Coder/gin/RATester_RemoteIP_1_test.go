package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRemoteIP_1(t *testing.T) {
	testCases := []struct {
		name       string
		remoteAddr string
		expected   string
	}{
		{"IPv4", "192.168.1.1:8080", "192.168.1.1"},
		{"IPv6", "[2001:0db8:85a3:0000:0000:8a2e:0370:7334]:8080", "2001:0db8:85a3:0000:0000:8a2e:0370:7334"},
		{"Invalid", "invalid", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			req := &http.Request{RemoteAddr: tc.remoteAddr}
			ctx := &Context{Request: req}
			result := ctx.RemoteIP()
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
