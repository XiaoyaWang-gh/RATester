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

	t.Parallel()

	ctx := context.Background()
	reqid := "test_reqid"
	newCtx := NewReqidContext(ctx, reqid)

	if newCtx.Value(reqidKey) != reqid {
		t.Errorf("Expected reqid to be %s, got %s", reqid, newCtx.Value(reqidKey))
	}
}
