package client

import (
	"fmt"
	"testing"
)

func TestR_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{}
	request := client.R()

	if request.client != client {
		t.Errorf("Expected client to be set in request")
	}
}
