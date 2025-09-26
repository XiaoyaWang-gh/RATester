package ssdb

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestSessionReleaseIfPresent_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &SessionStore{}
	c := context.Background()
	w := &httptest.ResponseRecorder{}
	s.SessionReleaseIfPresent(c, w)
}
