package client

import (
	"fmt"
	"testing"
)

func TestDisableDebug_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{
		debug: true,
	}

	client.DisableDebug()

	if client.debug {
		t.Error("Debug should be disabled")
	}
}
