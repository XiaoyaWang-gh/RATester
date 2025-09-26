package acme

import (
	"fmt"
	"testing"
)

func TestListenSaveAction_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	s := &LocalStore{
		saveDataChan: make(chan map[string]*StoredData),
		filename:     "filename",
		storedData:   make(map[string]*StoredData),
	}

	// when
	s.listenSaveAction()

	// then
	// TODO
}
