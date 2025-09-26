package pmux

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestAddr_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	listener := &PortListener{
		addr: &net.TCPAddr{
			IP:   net.IPv4(127, 0, 0, 1),
			Port: 8080,
			Zone: "",
		},
	}

	expected := &net.TCPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8080,
		Zone: "",
	}

	actual := listener.Addr()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
