package couchbase

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestSessionReleaseIfPresent_1(t *testing.T) {
	type args struct {
		c context.Context
		w http.ResponseWriter
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

			cs := &SessionStore{}
			cs.SessionReleaseIfPresent(tt.args.c, tt.args.w)
		})
	}
}
