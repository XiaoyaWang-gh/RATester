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

	testFunc := func(conn net.Conn, host string) (net.Conn, error) {
		return conn, nil
	}

	muxer.SetRewriteHostFunc(testFunc)

	if muxer.rewriteHost == nil {
		t.Errorf("Expected SetRewriteHostFunc to set rewriteHost function, but it was not set.")
	}
}
