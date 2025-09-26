package vhost

import (
	"fmt"
	"testing"
)

func TestAddr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &Listener{}
	if l.Addr() != nil {
		t.Error("Addr() should return nil")
	}
}
