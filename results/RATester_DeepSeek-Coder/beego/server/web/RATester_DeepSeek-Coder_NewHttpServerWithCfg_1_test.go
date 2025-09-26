package web

import (
	"fmt"
	"testing"
)

func TestNewHttpServerWithCfg_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cfg := &Config{
		AppName: "testApp",
		RunMode: "test",
		// other fields...
	}

	server := NewHttpServerWithCfg(cfg)

	if server.Cfg.AppName != "testApp" {
		t.Errorf("Expected AppName to be 'testApp', got %s", server.Cfg.AppName)
	}

	if server.Cfg.RunMode != "test" {
		t.Errorf("Expected RunMode to be 'test', got %s", server.Cfg.RunMode)
	}

	// other assertions...
}
