package server

import (
	"fmt"
	"testing"
)

func TestNewControlManager_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cm := NewControlManager()
	if cm == nil {
		t.Fatal("NewControlManager() should not return nil")
	}
	if cm.ctlsByRunID == nil {
		t.Fatal("NewControlManager() should initialize ctlsByRunID")
	}
}
