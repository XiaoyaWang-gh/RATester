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

	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctx, resp := NewMockContext(req)

	if ctx == nil {
		t.Error("Expected non-nil Context, got nil")
	}

	if resp == nil {
		t.Error("Expected non-nil ResponseRecorder, got nil")
	}

	if ctx.Request != req {
		t.Errorf("Expected Request to be %v, got %v", req, ctx.Request)
	}

	if resp.Code != 0 {
		t.Errorf("Expected ResponseRecorder Code to be 0, got %d", resp.Code)
	}
}
