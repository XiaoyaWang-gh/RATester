package file

import (
	"fmt"
	"sync"
	"testing"
)

func TestDelHost_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &DbUtils{
		JsonDb: &JsonDb{
			Hosts: sync.Map{},
		},
	}

	id := 1
	db.JsonDb.Hosts.Store(id, "test")

	err := db.DelHost(id)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	_, ok := db.JsonDb.Hosts.Load(id)
	if ok {
		t.Errorf("Expected host to be deleted, but it still exists")
	}
}
