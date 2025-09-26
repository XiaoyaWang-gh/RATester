package file

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStoreClientsToJsonFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &JsonDb{
		ClientFilePath: "testdata/clients.json",
	}

	// Add some clients to the sync.Map
	db.Clients.Store(1, &Client{Id: 1, Addr: "192.0.2.1"})
	db.Clients.Store(2, &Client{Id: 2, Addr: "192.0.2.2"})

	// Call the method under test
	db.StoreClientsToJsonFile()

	// Check that the JSON file was written correctly
	data, err := os.ReadFile(db.ClientFilePath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	var clients []*Client
	if err := json.Unmarshal(data, &clients); err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	if len(clients) != 2 {
		t.Errorf("Expected 2 clients, got %d", len(clients))
	}

	for _, client := range clients {
		if client.Id != 1 && client.Id != 2 {
			t.Errorf("Unexpected client ID: %d", client.Id)
		}
		if client.Addr != "192.0.2.1" && client.Addr != "192.0.2.2" {
			t.Errorf("Unexpected client address: %s", client.Addr)
		}
	}
}
