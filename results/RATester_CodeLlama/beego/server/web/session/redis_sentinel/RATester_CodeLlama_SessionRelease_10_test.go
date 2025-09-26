package redis_sentinel

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestSessionRelease_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	var rs *SessionStore
	var ctx context.Context
	var w http.ResponseWriter
	rs.SessionRelease(ctx, w)
}
