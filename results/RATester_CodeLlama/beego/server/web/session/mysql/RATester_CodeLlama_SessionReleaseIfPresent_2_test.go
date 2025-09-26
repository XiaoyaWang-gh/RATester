package mysql

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestSessionReleaseIfPresent_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	st := &SessionStore{}
	ctx := context.Background()
	w := &httptest.ResponseRecorder{}
	st.SessionReleaseIfPresent(ctx, w)
}
