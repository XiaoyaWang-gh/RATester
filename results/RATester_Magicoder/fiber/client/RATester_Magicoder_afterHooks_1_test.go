package client

import (
	"fmt"
	"testing"
)

func TestafterHooks_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{}
	req := &Request{}
	resp := &Response{}

	core := &core{
		client: client,
		req:    req,
	}

	err := core.afterHooks(resp)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
