package fiber

import (
	"fmt"
	"testing"
)

func TestJSONP_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var c *DefaultCtx
	var data any
	var callback []string
	var err error
	var result string

	// TEST
	err = c.JSONP(data, callback...)
	if err != nil {
		t.Fatalf("JSONP failed: %s", err.Error())
	}

	// TEST
	result = c.app.getString(c.fasthttp.Response.Body())
	if result != "" {
		t.Fatalf("JSONP failed: %s", result)
	}
}
