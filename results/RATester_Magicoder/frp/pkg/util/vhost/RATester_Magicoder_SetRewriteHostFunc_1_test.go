package vhost

import (
	"fmt"
	"net"
	"testing"
)

func TestSetRewriteHostFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	muxer := &Muxer{}

	testFunc := func(c net.Conn, host string) (net.Conn, error) {
		// Test logic here
		return c, nil
	}

	muxer.SetRewriteHostFunc(testFunc)

	if muxer.rewriteHost == nil {
		t.Error("Expected rewriteHost to be set, but it was not.")
	}
}
