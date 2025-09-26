package vhost

import (
	"fmt"
	"net"
	"testing"
)

func TestSetSuccessHookFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	muxer := &Muxer{}

	testFunc := func(c net.Conn, m map[string]string) error {
		// Test logic here
		return nil
	}

	muxer.SetSuccessHookFunc(testFunc)

	if muxer.successHook == nil {
		t.Error("Expected successHook to be set, but it was not.")
	}
}
