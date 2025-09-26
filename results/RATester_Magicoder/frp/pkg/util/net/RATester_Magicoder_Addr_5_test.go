package net

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestAddr_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ln := &WebsocketListener{
		ln: &net.TCPListener{},
	}
	expected := ln.ln.Addr()
	actual := ln.Addr()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
