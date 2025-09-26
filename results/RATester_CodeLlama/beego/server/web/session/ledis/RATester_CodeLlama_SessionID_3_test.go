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

	ls := &SessionStore{}
	ls.sid = "test"
	ctx := context.Background()
	if ls.SessionID(ctx) != "test" {
		t.Error("SessionID failed")
	}
}
