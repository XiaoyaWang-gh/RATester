package redis

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestSessionReleaseIfPresent_4(t *testing.T) {
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
