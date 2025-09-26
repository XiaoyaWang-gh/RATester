package file

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetHostById_1(t *testing.T) {
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

	testHost := &Host{
		Id: 1,
	}

	db.JsonDb.Hosts.Store(testHost.Id, testHost)

	host, err := db.GetHostById(testHost.Id)
	if err != nil {
		t.Errorf("GetHostById returned an error: %v", err)
	}

	if host.Id != testHost.Id {
		t.Errorf("GetHostById returned the wrong host. Expected: %v, Got: %v", testHost, host)
	}

	_, err = db.GetHostById(2)
	if err == nil {
		t.Error("GetHostById did not return an error for a non-existent host")
	}
}
