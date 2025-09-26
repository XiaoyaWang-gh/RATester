package redis_cluster

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestSessionReleaseIfPresent_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	var rs *SessionStore
	var ctx context.Context
	var w http.ResponseWriter
	rs.SessionReleaseIfPresent(ctx, w)
}
