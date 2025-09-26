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
		t.Fatal("NewInternalListener() should not return nil")
	}

	if l.acceptCh == nil {
		t.Fatal("NewInternalListener() should initialize acceptCh")
	}

	if l.closed {
		t.Fatal("NewInternalListener() should not be closed initially")
	}

	if err := l.Close(); err != nil {
		t.Fatalf("l.Close() should not return error, but got %v", err)
	}

	if !l.closed {
		t.Fatal("l.Close() should close the listener")
	}
}
