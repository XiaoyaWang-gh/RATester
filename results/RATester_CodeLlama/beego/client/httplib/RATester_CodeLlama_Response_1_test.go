package httplib

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

	b := &BeegoHTTPRequest{}
	_, err := b.Response()
	if err != nil {
		t.Errorf("Response() error = %v", err)
		return
	}
}
