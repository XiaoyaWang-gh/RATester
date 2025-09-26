package mysql

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionID_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	st := &SessionStore{}
	st.sid = "123456"
	ctx := context.Background()
	if st.SessionID(ctx) != "123456" {
		t.Error("SessionID failed")
	}
}
