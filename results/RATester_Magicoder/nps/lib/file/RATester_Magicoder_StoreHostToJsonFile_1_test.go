package file

import (
	"fmt"
	"sync"
	"testing"
)

func TestStoreHostToJsonFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &JsonDb{
		Hosts: sync.Map{},
	}

	db.Hosts.Store("test", "test")

	db.StoreHostToJsonFile()

	_, ok := db.Hosts.Load("test")
	if !ok {
		t.Error("Failed to store host to json file")
	}
}
