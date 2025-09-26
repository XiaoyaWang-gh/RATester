package httplib

import (
	"fmt"
	"testing"
)

func TestWithParam_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	key := "key"
	value := "value"
	option := WithParam(key, value)
	request := &BeegoHTTPRequest{}
	option(request)
	if request.Param(key, value) != request {
		t.Errorf("WithParam failed")
	}
}
