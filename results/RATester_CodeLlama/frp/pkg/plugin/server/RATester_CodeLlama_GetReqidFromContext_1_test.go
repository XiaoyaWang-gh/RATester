package plugin

import (
	"context"
	"fmt"
	"testing"
)

func TestGetReqidFromContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	reqid := "reqid"
	ctx = context.WithValue(ctx, reqidKey, reqid)
	if GetReqidFromContext(ctx) != reqid {
		t.Errorf("GetReqidFromContext() = %v, want %v", GetReqidFromContext(ctx), reqid)
	}
}
