package file

import (
	"fmt"
	"sync"
	"testing"
)

func TestVerifyUserName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &DbUtils{
		JsonDb: &JsonDb{
			Clients: sync.Map{},
		},
	}

	// Test case 1: User name exists but is not the same as the provided id
	db.JsonDb.Clients.Store(1, &Client{WebUserName: "testUser"})
	res := db.VerifyUserName("testUser", 2)
	if res != false {
		t.Errorf("Expected false, got %v", res)
	}

	// Test case 2: User name does not exist
	res = db.VerifyUserName("nonExistentUser", 1)
	if res != true {
		t.Errorf("Expected true, got %v", res)
	}

	// Test case 3: User name exists and is the same as the provided id
	res = db.VerifyUserName("testUser", 1)
	if res != true {
		t.Errorf("Expected true, got %v", res)
	}
}
