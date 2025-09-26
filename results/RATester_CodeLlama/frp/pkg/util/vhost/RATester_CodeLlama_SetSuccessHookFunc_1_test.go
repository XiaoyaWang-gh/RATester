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

	m := &Muxer{}
	m.SetSuccessHookFunc(func(net.Conn, map[string]string) error {
		return nil
	})
	if m.successHook == nil {
		t.Fatal("successHookFunc is not set")
	}
}
