package plugin

import (
	"context"
	"fmt"
	"testing"
)

func TestNewReqidContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	reqid := "test_reqid"
	newCtx := NewReqidContext(ctx, reqid)

	if newCtx.Value(reqidKey) != reqid {
		t.Errorf("Expected reqid to be %s, but got %v", reqid, newCtx.Value(reqidKey))
	}
}
