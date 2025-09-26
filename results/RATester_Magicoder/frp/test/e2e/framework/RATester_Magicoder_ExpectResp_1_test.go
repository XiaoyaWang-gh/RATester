package framework

import (
	"fmt"
	"reflect"
	"testing"
)

func TestExpectResp_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	reqExpect := &RequestExpect{}
	resp := []byte("expected response")
	reqExpect.ExpectResp(resp)

	if !reflect.DeepEqual(reqExpect.expectResp, resp) {
		t.Errorf("Expected response does not match the actual response")
	}
}
