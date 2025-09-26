package client

import (
	"fmt"
	"testing"
)

func TestAddResponseHook_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{}

	hook1 := func(c *Client, resp *Response, req *Request) error {
		// do something
		return nil
	}
	hook2 := func(c *Client, resp *Response, req *Request) error {
		// do something
		return nil
	}

	client.AddResponseHook(hook1, hook2)

	if len(client.userResponseHooks) != 2 {
		t.Errorf("Expected 2 hooks, got %d", len(client.userResponseHooks))
	}
}
