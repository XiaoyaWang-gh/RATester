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
		AppName: "TestApp",
		// other fields
	}

	server := NewHttpServerWithCfg(cfg)

	if server.Cfg.AppName != cfg.AppName {
		t.Errorf("Expected AppName to be %s, but got %s", cfg.AppName, server.Cfg.AppName)
	}

	// Add more assertions as needed for other fields in the Config struct
}
