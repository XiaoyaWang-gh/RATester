package session

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestIsSecure_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	manager := &Manager{
		config: &ManagerConfig{
			Secure: true,
		},
	}
	req := &http.Request{}
	req.URL = &url.URL{}
	req.URL.Scheme = "https"
	req.TLS = &tls.ConnectionState{}
	if !manager.isSecure(req) {
		t.Error("Test isSecure failed")
	}
}
