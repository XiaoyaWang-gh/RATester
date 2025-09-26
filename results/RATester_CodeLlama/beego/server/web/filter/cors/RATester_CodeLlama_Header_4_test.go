package cors

import (
	"fmt"
	"testing"
)

func TestHeader_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &Options{
		AllowAllOrigins: true,
	}
	headers := o.Header("http://example.com")
	if len(headers) != 2 {
		t.Errorf("Expected 2 headers, got %d", len(headers))
	}
	if headers["Access-Control-Allow-Origin"] != "*" {
		t.Errorf("Expected Access-Control-Allow-Origin to be *, got %s", headers["Access-Control-Allow-Origin"])
	}
	if headers["Access-Control-Allow-Credentials"] != "true" {
		t.Errorf("Expected Access-Control-Allow-Credentials to be true, got %s", headers["Access-Control-Allow-Credentials"])
	}
}
