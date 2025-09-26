package redis_sentinel

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionID_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	rs := &SessionStore{
		sid: "test_sid",
	}

	expected := "test_sid"
	actual := rs.SessionID(ctx)

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
