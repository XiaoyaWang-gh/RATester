package file

import (
	"fmt"
	"sync"
	"testing"
)

func TestNewTask_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &DbUtils{
		JsonDb: &JsonDb{
			Tasks: sync.Map{},
		},
	}

	tunnel := &Tunnel{
		Id:       1,
		Port:     8080,
		Mode:     "secret",
		Password: "test_password",
	}

	err := db.NewTask(tunnel)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	_, ok := db.JsonDb.Tasks.Load(tunnel.Id)
	if !ok {
		t.Errorf("Expected task to be stored in the map, but it was not")
	}

	tunnel2 := &Tunnel{
		Id:       2,
		Port:     8080,
		Mode:     "secret",
		Password: "test_password",
	}

	err = db.NewTask(tunnel2)
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	_, ok = db.JsonDb.Tasks.Load(tunnel2.Id)
	if ok {
		t.Errorf("Expected task to not be stored in the map, but it was")
	}
}
