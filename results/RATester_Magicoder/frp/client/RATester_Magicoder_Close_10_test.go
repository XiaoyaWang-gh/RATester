package client

import (
	"context"
	"fmt"
	"testing"
)

func TestClose_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	svr := &Service{
		ctl: &Control{
			ctx: context.Background(),
		},
	}

	svr.Close()

	if svr.ctl.ctx.Err() != context.Canceled {
		t.Errorf("Expected context to be canceled, but it's not")
	}
}
