package net

import (
	"fmt"
	"testing"
)

func TestRead_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &FakeUDPConn{
		packets: make(chan []byte, 1),
	}
	c.packets <- []byte("hello")
	b := make([]byte, 10)
	n, err := c.Read(b)
	if err != nil {
		t.Fatal(err)
	}
	if n != 5 {
		t.Fatal("read wrong length")
	}
	if string(b[:n]) != "hello" {
		t.Fatal("read wrong content")
	}
}
