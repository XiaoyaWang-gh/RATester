package create

import (
	"fmt"
	"net/http"
	"testing"
)

func TestAddDefaultHeaders_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := &http.Request{}
	addDefaultHeaders(req)
	if !hasHeaderKey(req.Header, "User-Agent") {
		t.Errorf("addDefaultHeaders() failed")
	}
}
