package client

import (
	"fmt"
	"testing"
)

func TestBasicAuth_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	username := "test"
	password := "test"
	want := "dGVzdDp0ZXN0"
	got := basicAuth(username, password)
	if got != want {
		t.Errorf("basicAuth() = %v, want %v", got, want)
	}
}
