package common

import (
	"fmt"
	"testing"
)

func TestInit_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	storeMsg := &StoreMsg{}
	err := storeMsg.Init("config")
	if err != nil {
		t.Errorf("Init() failed, expected nil, got %v", err)
	}
}
