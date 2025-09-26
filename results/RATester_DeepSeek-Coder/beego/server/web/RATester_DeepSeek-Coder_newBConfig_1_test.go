package web

import (
	"fmt"
	"testing"
)

func TestNewBConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	config := newBConfig()

	if config.AppName != "beego" {
		t.Errorf("Expected AppName to be 'beego', got %s", config.AppName)
	}
}
