package client

import (
	"fmt"
	"testing"
)

func TestsetClient_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{}
	response := &Response{}

	response.setClient(client)

	if response.client != client {
		t.Errorf("Expected client to be set, but it was not.")
	}
}
