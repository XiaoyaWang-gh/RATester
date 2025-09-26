package echo

import (
	"fmt"
	"testing"
)

func TestResponse_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &context{
		response: &Response{},
	}
	if c.Response() != c.response {
		t.Errorf("Response() = %v, want %v", c.Response(), c.response)
	}
}
