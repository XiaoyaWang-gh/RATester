package ledis

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestSessionReleaseIfPresent_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ls := &SessionStore{}
	c := context.Background()
	w := &httptest.ResponseRecorder{}
	ls.SessionReleaseIfPresent(c, w)
}
