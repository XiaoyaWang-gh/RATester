package main

import (
	"fmt"
	"testing"
)

func TestLoadConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	host, vkey, connType := loadConfig()
	if host == "" || vkey == "" || connType == "" {
		t.Error("loadConfig failed")
	}
}
