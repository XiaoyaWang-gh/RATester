package framework

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/test/e2e/pkg/request"
)

func TestDo_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := &request.Request{}
	e := &RequestExpect{req: req}

	resp, err := e.Do()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if resp == nil {
		t.Error("Expected non-nil response, but got nil")
	}
}
