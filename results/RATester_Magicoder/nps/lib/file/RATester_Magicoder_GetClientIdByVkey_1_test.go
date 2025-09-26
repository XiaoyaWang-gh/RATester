package file

import (
	"fmt"
	"sync"
	"testing"

	"ehang.io/nps/lib/crypt"
)

func TestGetClientIdByVkey_1(t *testing.T) {
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

	// Test case 1: Client exists
	vkey := "existing_vkey"
	existingClient := &Client{
		Id:        1,
		VerifyKey: crypt.Md5(vkey),
	}
	db.JsonDb.Clients.Store(vkey, existingClient)
	id, err := db.GetClientIdByVkey(vkey)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if id != existingClient.Id {
		t.Errorf("Expected id %d, but got %d", existingClient.Id, id)
	}

	// Test case 2: Client does not exist
	nonExistingVkey := "non_existing_vkey"
	id, err = db.GetClientIdByVkey(nonExistingVkey)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
	if id != 0 {
		t.Errorf("Expected id 0, but got %d", id)
	}
}
