package request

import (
	"fmt"
	"testing"
	"time"

	httppkg "github.com/fatedier/frp/pkg/util/http"
)

func TestHTTPAuth_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{
		protocol: "http",
		addr:     "localhost",
		port:     8080,
		body:     []byte("test"),
		timeout:  10 * time.Second,
	}

	user := "testUser"
	password := "testPassword"

	r.HTTPAuth(user, password)

	if r.authValue != httppkg.BasicAuth(user, password) {
		t.Errorf("HTTPAuth() failed, expected %s, got %s", httppkg.BasicAuth(user, password), r.authValue)
	}
}
