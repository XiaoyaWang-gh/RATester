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

	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	ctx, resp := NewMockContext(req)
	if ctx == nil || resp == nil {
		t.Error("NewMockContext returned nil context or response")
	}
	if ctx.Request != req {
		t.Errorf("Expected request to be %v, got %v", req, ctx.Request)
	}
	if resp.Code != 0 {
		t.Errorf("Expected response code to be 0, got %v", resp.Code)
	}
}
