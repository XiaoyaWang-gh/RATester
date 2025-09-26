package httplib

import (
	"fmt"
	"testing"
)

func TestHandleFiles_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: setup your test
	b := &BeegoHTTPRequest{}
	b.files = map[string]string{}
	b.params = map[string][]string{}
	// TODO: test and assert stuff

}
