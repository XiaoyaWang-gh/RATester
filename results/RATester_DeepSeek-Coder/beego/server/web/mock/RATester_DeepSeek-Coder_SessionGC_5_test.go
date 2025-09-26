package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionGC_5(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		s    *SessionProvider
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

			tt.s.SessionGC(tt.args.ctx)
		})
	}
}
