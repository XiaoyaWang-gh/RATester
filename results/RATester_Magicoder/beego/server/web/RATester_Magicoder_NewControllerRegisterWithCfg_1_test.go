package web

import (
	"fmt"
	"testing"
)

func TestNewControllerRegisterWithCfg_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cfg := &Config{
		AppName: "testApp",
		RunMode: "testMode",
		// other fields
	}
	cr := NewControllerRegisterWithCfg(cfg)

	// Testing the fields of the returned ControllerRegister
	if cr.cfg.AppName != cfg.AppName {
		t.Errorf("Expected AppName to be %s, but got %s", cfg.AppName, cr.cfg.AppName)
	}
	if cr.cfg.RunMode != cfg.RunMode {
		t.Errorf("Expected RunMode to be %s, but got %s", cfg.RunMode, cr.cfg.RunMode)
	}
	// Add more tests for other fields as needed
}
