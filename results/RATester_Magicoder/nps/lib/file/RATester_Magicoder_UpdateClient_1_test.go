package file

import (
	"fmt"
	"sync"
	"testing"
)

func TestUpdateClient_1(t *testing.T) {
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

	client := &Client{
		Id: 1,
	}

	err := db.UpdateClient(client)
	if err != nil {
		t.Errorf("UpdateClient failed: %v", err)
	}

	_, ok := db.JsonDb.Clients.Load(client.Id)
	if !ok {
		t.Errorf("Client not found in the map")
	}

	loadedClient, _ := db.JsonDb.Clients.Load(client.Id)
	if loadedClient.(*Client) != client {
		t.Errorf("Loaded client does not match the original client")
	}
}
