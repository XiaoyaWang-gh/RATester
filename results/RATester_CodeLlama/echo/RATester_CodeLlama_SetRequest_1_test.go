package echo

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSetRequest_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &context{}
	r := &http.Request{}
	c.SetRequest(r)
	if c.request != r {
		t.Errorf("Expected %v, actual %v", r, c.request)
	}
}
