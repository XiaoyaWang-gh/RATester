package tls

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetStore_1(t *testing.T) {
	manager := &Manager{
		lock:         sync.RWMutex{},
		storesConfig: make(map[string]Store),
		stores:       make(map[string]*CertificateStore),
		configs:      make(map[string]Options),
	}

	storeName := "testStore"
	manager.stores[storeName] = &CertificateStore{}

	t.Run("should return the store if it exists", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := manager.GetStore(storeName)
		if result != manager.stores[storeName] {
			t.Errorf("Expected %v, got %v", manager.stores[storeName], result)
		}
	})

	t.Run("should return nil if the store does not exist", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := manager.GetStore("nonExistentStore")
		if result != nil {
			t.Errorf("Expected nil, got %v", result)
		}
	})
}
