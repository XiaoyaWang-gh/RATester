package ledis

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionID_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	store := &SessionStore{
		sid: "test_session_id",
	}

	expected := "test_session_id"
	actual := store.SessionID(ctx)

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
