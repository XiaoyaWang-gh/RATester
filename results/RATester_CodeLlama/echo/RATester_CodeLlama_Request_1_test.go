package echo

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRequest_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &context{
		request: &http.Request{},
	}
	if c.Request() != c.request {
		t.Errorf("Request() = %v, want %v", c.Request(), c.request)
	}
}
