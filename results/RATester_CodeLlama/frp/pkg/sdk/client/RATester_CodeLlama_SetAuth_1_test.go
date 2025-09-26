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

	c := &Client{}
	c.SetAuth("user", "pwd")
	if c.authUser != "user" {
		t.Errorf("authUser should be user, but got %s", c.authUser)
	}
	if c.authPwd != "pwd" {
		t.Errorf("authPwd should be pwd, but got %s", c.authPwd)
	}
}
