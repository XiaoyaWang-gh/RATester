package client

import (
	"fmt"
	"testing"
)

func TestClose_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &TRPClient{
		svrAddr: "localhost:8080",
	}

	client.Close()

	// Add assertions here to verify the expected behavior
}
