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
	config := "test_config"
	err := storeMsg.Init(config)
	if err != nil {
		t.Errorf("Init() error = %v, wantErr %v", err, nil)
		return
	}
}
