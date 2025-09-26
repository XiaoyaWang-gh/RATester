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
	reqid := "123456"
	ctx = NewReqidContext(ctx, reqid)
	if ctx.Value(reqidKey) != reqid {
		t.Errorf("NewReqidContext failed")
	}
}
