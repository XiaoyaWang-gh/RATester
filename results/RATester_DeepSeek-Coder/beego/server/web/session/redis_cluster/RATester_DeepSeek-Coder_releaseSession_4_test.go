package redis_cluster

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestReleaseSession_4(t *testing.T) {
	type args struct {
		ctx            context.Context
		w              http.ResponseWriter
		requirePresent bool
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

			rs := &SessionStore{}
			rs.releaseSession(tt.args.ctx, tt.args.w, tt.args.requirePresent)
		})
	}
}
