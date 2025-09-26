package consulcatalog

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestRepeatSend_1(t *testing.T) {
	type args struct {
		ctx      context.Context
		interval time.Duration
		c        chan struct{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			repeatSend(tt.args.ctx, tt.args.interval, tt.args.c)
		})
	}
}
