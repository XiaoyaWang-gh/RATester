package session

import (
	"fmt"
	"testing"
)

func TestAcquireSession_1(t *testing.T) {
	t.Run("Testing acquireSession", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		session := acquireSession()
		if session == nil {
			t.Errorf("Expected a session, got nil")
		}
		if session.data == nil {
			t.Errorf("Expected session data, got nil")
		}
		if session.byteBuffer == nil {
			t.Errorf("Expected byte buffer, got nil")
		}
		if !session.fresh {
			t.Errorf("Expected fresh session, got not fresh")
		}
	})
}
