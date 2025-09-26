package client

import (
	"fmt"
	"testing"
)

func TestSetAuth_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{
		address: "http://localhost:8080",
	}

	user := "testUser"
	pwd := "testPassword"

	client.SetAuth(user, pwd)

	if client.authUser != user {
		t.Errorf("Expected authUser to be %s, got %s", user, client.authUser)
	}

	if client.authPwd != pwd {
		t.Errorf("Expected authPwd to be %s, got %s", pwd, client.authPwd)
	}
}
