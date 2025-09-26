package vhost

import (
	"fmt"
	"net"
	"testing"
)

func TestSetCheckAuthFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	muxer := &Muxer{}

	testFunc := func(conn net.Conn, username, password string, reqInfoMap map[string]string) (bool, error) {
		return true, nil
	}

	muxer.SetCheckAuthFunc(testFunc)

	if muxer.checkAuth == nil {
		t.Error("checkAuth function is not set")
	}
}
