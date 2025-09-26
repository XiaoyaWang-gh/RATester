package file

import (
	"fmt"
	"testing"
)

func TestHasHost_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{Id: 1}
	host := &Host{Client: client, Host: "example.com", Location: "/"}

	GetDb().JsonDb.Hosts.Store(1, host)

	if !client.HasHost(host) {
		t.Error("Expected true, got false")
	}

	GetDb().JsonDb.Hosts.Delete(1)

	if client.HasHost(host) {
		t.Error("Expected false, got true")
	}
}
