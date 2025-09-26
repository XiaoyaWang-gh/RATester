package framework

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/fatedier/frp/test/e2e/pkg/request"
)

func TestExpectResp_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &RequestExpect{
		req: &request.Request{},
	}

	resp := []byte("expected response")
	e.ExpectResp(resp)

	if !bytes.Equal(e.expectResp, resp) {
		t.Errorf("Expected response to be %v, but got %v", resp, e.expectResp)
	}
}
