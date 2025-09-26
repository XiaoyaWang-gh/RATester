package memory

import (
	"fmt"
	"testing"
	"time"
)

func Testgc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	storage := &Storage{
		db:         make(map[string]entry),
		done:       make(chan struct{}),
		gcInterval: time.Second,
	}

	storage.Set("key1", []byte("value1"), time.Second*10)
	storage.Set("key2", []byte("value2"), time.Second*10)

	time.Sleep(time.Second * 11)

	_, err := storage.Get("key1")
	if err == nil {
		t.Error("Expected error, got nil")
	}

	_, err = storage.Get("key2")
	if err != nil {
		t.Errorf("Expected nil, got error: %v", err)
	}
}
