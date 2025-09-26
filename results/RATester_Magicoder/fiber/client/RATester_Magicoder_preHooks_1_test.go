package client

import (
	"fmt"
	"testing"
)

func TestpreHooks_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{
		userRequestHooks: []RequestHook{
			func(c *Client, r *Request) error {
				// Add your test logic here
				return nil
			},
		},
		builtinRequestHooks: []RequestHook{
			func(c *Client, r *Request) error {
				// Add your test logic here
				return nil
			},
		},
	}

	core := &core{
		client: client,
		req: &Request{
			client: client,
		},
	}

	err := core.preHooks()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
