package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestBasicAuth_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	accounts := Accounts{
		"user1": "password1",
		"user2": "password2",
	}

	handler := BasicAuth(accounts)

	// Test case 1: Valid credentials
	req, _ := http.NewRequest("GET", "/", nil)
	req.SetBasicAuth("user1", "password1")
	ctx := &Context{Request: req}
	handler(ctx)
	if ctx.IsAborted() {
		t.Error("Expected not to abort the request with valid credentials")
	}

	// Test case 2: Invalid credentials
	req, _ = http.NewRequest("GET", "/", nil)
	req.SetBasicAuth("user1", "wrongpassword")
	ctx = &Context{Request: req}
	handler(ctx)
	if !ctx.IsAborted() {
		t.Error("Expected to abort the request with invalid credentials")
	}
}
