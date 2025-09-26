package session

import (
	"context"
	"fmt"
	"testing"
)

func TestFlush_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	st := &CookieSessionStore{}
	st.Flush(context.Background())
	if len(st.values) != 0 {
		t.Errorf("Flush() failed")
	}
}
