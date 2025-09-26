package acme

import (
	"fmt"
	"testing"
	"time"
)

func TestSave_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	store := &LocalStore{
		saveDataChan: make(chan map[string]*StoredData),
		filename:     "test.json",
		storedData:   make(map[string]*StoredData),
	}

	resolverName := "testResolver"
	storedData := &StoredData{
		Account: &Account{
			Email: "test@example.com",
		},
	}

	store.save(resolverName, storedData)

	select {
	case data := <-store.saveDataChan:
		if data[resolverName].Account.Email != storedData.Account.Email {
			t.Errorf("Expected email %s, but got %s", storedData.Account.Email, data[resolverName].Account.Email)
		}
	case <-time.After(time.Second):
		t.Error("Timeout waiting for saveDataChan")
	}
}
