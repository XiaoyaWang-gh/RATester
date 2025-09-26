package session

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestSessionReleaseIfPresent_7(t *testing.T) {
	type args struct {
		ctx context.Context
		w   http.ResponseWriter
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

			fs := &FileSessionStore{}
			fs.SessionReleaseIfPresent(tt.args.ctx, tt.args.w)
		})
	}
}
