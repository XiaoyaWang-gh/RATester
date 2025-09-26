package session

import (
	"fmt"
	"testing"
)

func TestacquireSession_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := acquireSession()
	if s == nil {
		t.Error("Expected a session, but got nil")
	}
	if s.data == nil {
		t.Error("Expected session data, but got nil")
	}
	if s.byteBuffer == nil {
		t.Error("Expected byte buffer, but got nil")
	}
	if !s.fresh {
		t.Error("Expected fresh session, but got not fresh")
	}
}
