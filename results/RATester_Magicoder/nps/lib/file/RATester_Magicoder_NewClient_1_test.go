package file

import (
	"fmt"
	"testing"
)

func TestNewClient_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	vKey := "testKey"
	noStore := false
	noDisplay := false
	client := NewClient(vKey, noStore, noDisplay)

	if client.VerifyKey != vKey {
		t.Errorf("Expected VerifyKey to be %s, but got %s", vKey, client.VerifyKey)
	}

	if client.NoStore != noStore {
		t.Errorf("Expected NoStore to be %t, but got %t", noStore, client.NoStore)
	}

	if client.NoDisplay != noDisplay {
		t.Errorf("Expected NoDisplay to be %t, but got %t", noDisplay, client.NoDisplay)
	}

	if client.Id != 0 {
		t.Errorf("Expected Id to be 0, but got %d", client.Id)
	}

	if client.Addr != "" {
		t.Errorf("Expected Addr to be empty, but got %s", client.Addr)
	}

	if client.Status != true {
		t.Errorf("Expected Status to be true, but got %t", client.Status)
	}

	// Add more test cases as needed
}
