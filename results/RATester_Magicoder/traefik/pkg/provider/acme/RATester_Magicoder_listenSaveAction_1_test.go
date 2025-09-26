package acme

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestListenSaveAction_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	store := &LocalStore{
		saveDataChan: make(chan map[string]*StoredData),
		filename:     "test.json",
	}

	data := map[string]*StoredData{
		"test": {
			Account: &Account{
				Email: "test@example.com",
			},
		},
	}

	store.saveDataChan <- data

	time.Sleep(1 * time.Second)

	_, err := os.Stat(store.filename)
	if err != nil {
		t.Errorf("Expected file to exist, but it does not")
	}

	os.Remove(store.filename)
}
