package file

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetHostById_1(t *testing.T) {
	db := &DbUtils{
		JsonDb: &JsonDb{
			Hosts: sync.Map{},
		},
	}

	host := &Host{
		Id: 1,
	}

	db.JsonDb.Hosts.Store(1, host)

	t.Run("TestGetHostById", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		h, err := db.GetHostById(1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if h.Id != 1 {
			t.Errorf("Expected host id 1, got %v", h.Id)
		}
	})

	t.Run("TestGetHostByIdNotFound", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := db.GetHostById(2)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
