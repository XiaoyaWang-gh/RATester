package framework

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/test/e2e/pkg/request"
)

func TestExpectResponseCode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	resp := &request.Response{Code: 200}
	ensureFunc := ExpectResponseCode(200)

	if !ensureFunc(resp) {
		t.Errorf("ExpectResponseCode failed")
	}
}
