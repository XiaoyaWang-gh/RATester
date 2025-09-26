package mock

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNewMockContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := &http.Request{}
	ctx, resp := NewMockContext(req)
	if ctx == nil {
		t.Error("ctx is nil")
	}
	if resp == nil {
		t.Error("resp is nil")
	}
}
