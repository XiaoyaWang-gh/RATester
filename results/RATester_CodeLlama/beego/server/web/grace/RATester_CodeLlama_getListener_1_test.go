package grace

import (
	"fmt"
	"testing"
)

func TestGetListener_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// CONTEXT_0
	srv := &Server{}
	// CONTEXT_1
	// CONTEXT_2
	laddr := "127.0.0.1:8080"
	// CONTEXT_3
	l, err := srv.getListener(laddr)
	if err != nil {
		t.Errorf("getListener error: %v", err)
		return
	}
	// CONTEXT_4
	if l == nil {
		t.Errorf("getListener error: l is nil")
		return
	}
	// CONTEXT_5
	if l.Addr() == nil {
		t.Errorf("getListener error: l.Addr() is nil")
		return
	}
	// CONTEXT_6
	if err != nil {
		t.Errorf("getListener error: %v", err)
		return
	}
}
