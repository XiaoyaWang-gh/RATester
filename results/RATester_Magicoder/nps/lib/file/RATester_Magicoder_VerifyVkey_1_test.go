package file

import (
	"fmt"
	"sync"
	"testing"
)

func TestVerifyVkey_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	dbUtils := &DbUtils{
		JsonDb: &JsonDb{
			Clients: sync.Map{},
		},
	}

	// Test case 1: VerifyVkey returns true when the vkey is not found in the clients
	dbUtils.JsonDb.Clients.Store("client1", &Client{
		Id:        1,
		VerifyKey: "vkey1",
	})
	res := dbUtils.VerifyVkey("vkey2", 1)
	if res != true {
		t.Errorf("Expected true, got %v", res)
	}

	// Test case 2: VerifyVkey returns false when the vkey is found in the clients
	res = dbUtils.VerifyVkey("vkey1", 1)
	if res != false {
		t.Errorf("Expected false, got %v", res)
	}

	// Test case 3: VerifyVkey returns true when the vkey is not found in the clients
	res = dbUtils.VerifyVkey("vkey2", 2)
	if res != true {
		t.Errorf("Expected true, got %v", res)
	}
}
