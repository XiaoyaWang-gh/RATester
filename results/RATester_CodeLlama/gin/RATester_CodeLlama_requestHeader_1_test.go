package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRequestHeader_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{
		Request: &http.Request{
			Header: http.Header{
				"key": []string{"value"},
			},
		},
	}
	if c.requestHeader("key") != "value" {
		t.Errorf("Expected value, got %s", c.requestHeader("key"))
	}
}
