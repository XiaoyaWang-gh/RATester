package request

import (
	"fmt"
	"testing"

	httppkg "github.com/fatedier/frp/pkg/util/http"
)

func TestHTTPAuth_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{}
	user := "testUser"
	password := "testPassword"
	r.HTTPAuth(user, password)
	if r.authValue != httppkg.BasicAuth(user, password) {
		t.Errorf("HTTPAuth failed, expected: %s, got: %s", httppkg.BasicAuth(user, password), r.authValue)
	}
}
