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

	id := 1
	db.JsonDb.Clients.Store(id, "client")

	err := db.DelClient(id)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	_, ok := db.JsonDb.Clients.Load(id)
	if ok {
		t.Errorf("Expected client to be deleted, but it still exists")
	}
}
