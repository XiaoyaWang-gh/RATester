package ping

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestWithContext_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc string
		ctx  context.Context
	}{
		{
			desc: "case 1: context with cancel",
			ctx: func() context.Context {
				ctx, cancel := context.WithCancel(context.Background())
				go func() {
					time.Sleep(1 * time.Second)
					cancel()
				}()
				return ctx
			}(),
		},
		{
			desc: "case 2: context with timeout",
			ctx: func() context.Context {
				ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
				return ctx
			}(),
		},
		{
			desc: "case 3: context with deadline",
			ctx: func() context.Context {
				ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
				return ctx
			}(),
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			h := &Handler{}
			h.WithContext(test.ctx)

			time.Sleep(2 * time.Second)

			if h.terminating {
				t.Errorf("expected h.terminating to be false, got true")
			}
		})
	}
}
