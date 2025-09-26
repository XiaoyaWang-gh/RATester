package file

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetIdByVerifyKey_1(t *testing.T) {
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

	// Test case 1: VerifyKey exists and client is active
	db.JsonDb.Clients.Store("key1", &Client{
		Id:        1,
		VerifyKey: "key1",
		Status:    true,
	})
	id, err := db.GetIdByVerifyKey("key1", "192.168.1.1")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if id != 1 {
		t.Errorf("Expected id to be 1, but got %v", id)
	}

	// Test case 2: VerifyKey exists but client is not active
	db.JsonDb.Clients.Store("key2", &Client{
		Id:        2,
		VerifyKey: "key2",
		Status:    false,
	})
	id, err = db.GetIdByVerifyKey("key2", "192.168.1.1")
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
	if id != 0 {
		t.Errorf("Expected id to be 0, but got %v", id)
	}

	// Test case 3: VerifyKey does not exist
	id, err = db.GetIdByVerifyKey("key3", "192.168.1.1")
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
	if id != 0 {
		t.Errorf("Expected id to be 0, but got %v", id)
	}
}
