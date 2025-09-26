package mock

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestSessionRelease_5(t *testing.T) {
	type args struct {
		ctx context.Context
		rw  http.ResponseWriter
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

			s := &SessionStore{}
			s.SessionRelease(tt.args.ctx, tt.args.rw)
		})
	}
}
