package httplib

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetRequest_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BeegoHTTPRequest{}
	b.req = &http.Request{}
	if b.GetRequest() != b.req {
		t.Errorf("TestGetRequest failed")
	}
}
