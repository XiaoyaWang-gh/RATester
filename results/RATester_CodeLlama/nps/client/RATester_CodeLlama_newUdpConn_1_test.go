package client

import (
	"fmt"
	"testing"
)

func TestNewUdpConn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: setup your test
	// TODO: call newUdpConn
	// TODO: verify the results
}
