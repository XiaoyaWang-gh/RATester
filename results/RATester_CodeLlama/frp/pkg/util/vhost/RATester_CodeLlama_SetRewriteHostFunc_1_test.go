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
	muxer.SetRewriteHostFunc(func(net.Conn, string) (net.Conn, error) {
		return nil, nil
	})
	if muxer.rewriteHost == nil {
		t.Error("rewriteHost should not be nil")
	}
}
