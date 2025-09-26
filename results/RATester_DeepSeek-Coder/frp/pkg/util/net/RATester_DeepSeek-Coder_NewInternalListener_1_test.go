package net

import (
	"fmt"
	"testing"
)

func TestNewInternalListener_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := NewInternalListener()

	if l == nil {
		t.Fatal("NewInternalListener() returned nil")
	}

	if l.acceptCh == nil {
		t.Fatal("NewInternalListener() returned listener with nil acceptCh")
	}

	if l.closed {
		t.Fatal("NewInternalListener() returned listener with closed set")
	}

	if cap(l.acceptCh) != 128 {
		t.Fatalf("NewInternalListener() returned listener with acceptCh capacity %d, expected 128", cap(l.acceptCh))
	}
}
