package framework

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/fatedier/frp/test/e2e/pkg/request"
)

func TestEnsure_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &RequestExpect{
		req:         &request.Request{},
		f:           &Framework{},
		expectResp:  []byte("expected response"),
		expectError: false,
		explain:     []interface{}{"test ensure"},
	}

	e.Ensure(func(resp *request.Response) bool {
		return bytes.Equal(resp.Content, e.expectResp)
	})
}
