package file

import (
	"fmt"
	"sync"
	"testing"
)

func TestDelClient_1(t *testing.T) {
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

	// Add a client to the database
	client := &Client{
		Id: 1,
	}
	db.JsonDb.Clients.Store(client.Id, client)

	// Delete the client
	err := db.DelClient(client.Id)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Check if the client is deleted
	_, ok := db.JsonDb.Clients.Load(client.Id)
	if ok {
		t.Errorf("Expected client to be deleted, but it's still in the database")
	}
}
