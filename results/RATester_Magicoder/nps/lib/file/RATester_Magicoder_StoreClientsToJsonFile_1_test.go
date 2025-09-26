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
	db.Clients.Store(1, &Client{Id: 1, VerifyKey: "key1"})
	db.Clients.Store(2, &Client{Id: 2, VerifyKey: "key2"})

	// Call the method under test
	db.StoreClientsToJsonFile()

	// Check if the clients were stored correctly in the JSON file
	file, err := os.Open(db.ClientFilePath)
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	var clients []*Client
	err = json.NewDecoder(file).Decode(&clients)
	if err != nil {
		t.Fatalf("Failed to decode JSON: %v", err)
	}

	if len(clients) != 2 {
		t.Errorf("Expected 2 clients, got %d", len(clients))
	}

	for _, client := range clients {
		if client.Id != 1 && client.Id != 2 {
			t.Errorf("Unexpected client ID: %d", client.Id)
		}
		if client.VerifyKey != "key1" && client.VerifyKey != "key2" {
			t.Errorf("Unexpected verify key: %s", client.VerifyKey)
		}
	}
}
