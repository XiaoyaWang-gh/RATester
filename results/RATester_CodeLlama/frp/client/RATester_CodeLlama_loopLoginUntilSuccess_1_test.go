package client

import (
	"fmt"
	"testing"
	"time"
)

func TestLoopLoginUntilSuccess_1(t *testing.T) {
	type args struct {
		maxInterval    time.Duration
		firstLoginExit bool
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

			svr := &Service{}
			svr.loopLoginUntilSuccess(tt.args.maxInterval, tt.args.firstLoginExit)
		})
	}
}
