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
	if clientId == 0 {
		t.Errorf("Expected clientId to be greater than 0, got %d", clientId)
	}

	client, err := db.GetClient(int(clientId))
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if client.Id != int(clientId) {
		t.Errorf("Expected client.Id to be %d, got %d", clientId, client.Id)
	}
}
