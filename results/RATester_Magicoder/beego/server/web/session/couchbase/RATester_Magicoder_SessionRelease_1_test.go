package couchbase

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestSessionRelease_1(t *testing.T) {
	type args struct {
		ctx  context.Context
		w    http.ResponseWriter
		cs   *SessionStore
		data map[interface{}]interface{}
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

			tt.args.cs.SessionRelease(tt.args.ctx, tt.args.w)
		})
	}
}
