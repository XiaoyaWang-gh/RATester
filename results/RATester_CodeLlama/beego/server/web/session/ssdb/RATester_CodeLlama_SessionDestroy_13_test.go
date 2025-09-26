package ssdb

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionDestroy_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Provider{}
	ctx := context.Background()
	sid := "123"
	err := p.SessionDestroy(ctx, sid)
	if err != nil {
		t.Errorf("SessionDestroy error %v", err)
	}
}
