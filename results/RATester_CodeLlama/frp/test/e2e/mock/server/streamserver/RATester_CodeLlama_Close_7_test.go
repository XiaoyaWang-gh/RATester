package streamserver

import (
	"fmt"
	"testing"
)

func TestClose_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Server{
		netType:     TCP,
		bindAddr:    "127.0.0.1",
		bindPort:    8080,
		respContent: []byte("hello world"),
	}
	err := s.Run()
	if err != nil {
		t.Fatal(err)
	}
	err = s.Close()
	if err != nil {
		t.Fatal(err)
	}
}
