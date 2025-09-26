package httpserver

import (
	"fmt"
	"testing"
)

func TestInitListener_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	server := &Server{
		bindAddr: "localhost",
		bindPort: 8080,
	}

	err := server.initListener()
	if err != nil {
		t.Errorf("initListener() failed: %v", err)
	}

	if server.l == nil {
		t.Error("initListener() did not set the listener")
	}

	server.Close()
}
