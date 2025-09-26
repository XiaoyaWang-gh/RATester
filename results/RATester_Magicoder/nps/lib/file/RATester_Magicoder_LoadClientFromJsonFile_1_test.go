package file

import (
	"fmt"
	"testing"
)

func TestLoadClientFromJsonFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &JsonDb{
		ClientFilePath: "testdata/clients.json",
	}

	db.LoadClientFromJsonFile()

	clientId := db.GetClientId()
	if clientId != 1 {
		t.Errorf("Expected client id to be 1, but got %d", clientId)
	}

	client, err := db.GetClient(1)
	if err != nil {
		t.Errorf("Error getting client: %v", err)
	}

	if client.Id != 1 {
		t.Errorf("Expected client id to be 1, but got %d", client.Id)
	}

	if client.RateLimit != 1024 {
		t.Errorf("Expected client rate limit to be 1024, but got %d", client.RateLimit)
	}
}
