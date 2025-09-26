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
	pwd := "testPwd"

	client.SetAuth(user, pwd)

	if client.authUser != user || client.authPwd != pwd {
		t.Errorf("Expected authUser to be %s and authPwd to be %s, but got %s and %s", user, pwd, client.authUser, client.authPwd)
	}
}
