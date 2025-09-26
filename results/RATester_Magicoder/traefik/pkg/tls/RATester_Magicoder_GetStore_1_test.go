package tls

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetStore_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	manager := &Manager{
		lock:         sync.RWMutex{},
		storesConfig: map[string]Store{},
		stores:       map[string]*CertificateStore{},
		configs:      map[string]Options{},
		certs:        []*CertAndStores{},
	}

	storeName := "testStore"
	expectedStore := &CertificateStore{}
	manager.stores[storeName] = expectedStore

	result := manager.GetStore(storeName)

	if result != expectedStore {
		t.Errorf("Expected %v, got %v", expectedStore, result)
	}
}
